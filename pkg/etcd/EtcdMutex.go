package etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"sync"
	"time"
)

const (
	defaultTTL   = 60
	defaultTry   = 3
	deleteAction = "delete"
	expireAction = "expire"
)

type EtcdMutex struct {
	TTL        int64  //租约时间(单位s)
	Key        string //etcd的key
	mutex      *sync.Mutex
	cancelFunc context.CancelFunc //关闭续租的func
	lease      clientv3.Lease
	leaseID    clientv3.LeaseID
	txn        clientv3.Txn
}

func main() {
	InitEtcd()
	NewLock("")
	for i := 0; i < 20; i++ {

		//groutine1
		go func() {
			err := etcdMutex.Lock()
			if err != nil {
				fmt.Println("groutine" + string(i) + "抢锁失败")
				fmt.Println(err)
				return
			}
			fmt.Println("groutine" + string(i) + "抢锁成功")
			time.Sleep(10 * time.Second)
			defer etcdMutex.UnLock()
		}()
	}
	time.Sleep(5 * time.Second)
}
func NewLock(key string, ttl int64) (*EtcdMutex, error) {
	var err error
	var ctx context.Context
	etcdMutex := &EtcdMutex{
		TTL: ttl,
		Key: key,
	}

	//创建事务
	etcdMutex.txn = clientv3.NewKV(EtcdClient).Txn(context.TODO())
	//上锁（创建租约，自动续租）
	etcdMutex.lease = clientv3.NewLease(EtcdClient)
	//设置一个ctx取消自动续租
	ctx, etcdMutex.cancelFunc = context.WithCancel(context.TODO())
	//设置租约时间（过期时间）
	leaseResp, err := etcdMutex.lease.Grant(context.TODO(), etcdMutex.TTL)
	if err != nil {
		return nil, err
	}
	// 拿到租约id
	etcdMutex.leaseID = leaseResp.ID

	//自动续租（不停地往管道中扔租约信息）
	leaseRespChan, err := etcdMutex.lease.KeepAlive(ctx, etcdMutex.leaseID)
	//启动一个协程去监听
	go listenLeaseChan(leaseRespChan)
	return etcdMutex, nil
}

func (em *EtcdMutex) Lock() error {
	em.mutex.Lock()
	for try := 1; try <= defaultTry; try++ {
		em.txn.If(clientv3.Compare(clientv3.CreateRevision(em.Key), "=", 0)).
			Then(clientv3.OpPut(em.Key, "", clientv3.WithLease(em.leaseID))).
			Else(clientv3.OpGet(em.Key))
		//提交事务
		txnResp, err := em.txn.Commit()
		if err != nil {
			return err
		}
		if txnResp.Succeeded { //判断txn.if条件是否成立
			return nil
		}
		if !txnResp.Succeeded && try < defaultTry {
			return fmt.Errorf("抢锁失败: %v", txnResp)
		}
	}
	return nil
}
func (em *EtcdMutex) UnLock() {
	em.cancelFunc()
	_, err := em.lease.Revoke(context.TODO(), em.leaseID)
	if err != nil {
		panic(err)
	}
	fmt.Println("释放了锁")
}
