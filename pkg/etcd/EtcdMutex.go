package etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"sync"
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

func NewLock(key string, ttl int64) (*EtcdMutex, error) {
	var err error
	ctx := context.TODO()
	etcdMutex := &EtcdMutex{
		TTL: ttl,
		Key: key,
	}
	etcdMutex.mutex = new(sync.Mutex)
	//创建事务
	etcdMutex.txn = clientv3.NewKV(EtcdClient).Txn(ctx)
	//上锁（创建租约，自动续租）
	etcdMutex.lease = clientv3.NewLease(EtcdClient)
	//设置一个ctx取消自动续租
	ctx, etcdMutex.cancelFunc = context.WithCancel(ctx)
	//设置租约时间（过期时间）
	leaseResp, err := etcdMutex.lease.Grant(ctx, etcdMutex.TTL)
	if err != nil {
		return nil, err
	}
	// 拿到租约id
	etcdMutex.leaseID = leaseResp.ID

	//自动续租（不停地往管道中扔租约信息）
	leaseRespChan, err := etcdMutex.lease.KeepAlive(ctx, etcdMutex.leaseID)
	if err != nil {
		panic(err)
	}
	//启动一个协程去监听续租情况 (不然一会儿就会写满keepalive的队列，导致大量异常日是志)
	go listenLeaseChan(leaseRespChan)
	return etcdMutex, nil
}
func listenLeaseChan(leaseRespChan <-chan *clientv3.LeaseKeepAliveResponse) {
	var (
		leaseKeepResp *clientv3.LeaseKeepAliveResponse
	)
	for {
		select {
		case leaseKeepResp = <-leaseRespChan:
			if leaseKeepResp == nil {
				fmt.Println("租约失效了")
				return
			}
		}
	}
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
			fmt.Println("抢锁失败: ", txnResp)
		}
	}
	return nil
}
func (em *EtcdMutex) UnLock() {
	em.mutex.Unlock()
	em.cancelFunc()
	_, err := em.lease.Revoke(context.TODO(), em.leaseID)
	if err != nil {
		panic(err)
	}
	fmt.Println("释放了锁")
}
