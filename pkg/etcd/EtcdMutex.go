package etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
)

type EtcdMutex struct {
	TTL        int64              //租约时间(单位s)
	Key        string             //etcd的key
	cancelFunc context.CancelFunc //关闭续租的func
	lease      clientv3.Lease
	leaseID    clientv3.LeaseID
	txn        clientv3.Txn
}

func (em *EtcdMutex) init() error {
	var err error
	var ctx context.Context
	//创建事务
	em.txn = clientv3.NewKV(EtcdClient).Txn(context.TODO())
	//上锁（创建租约，自动续租）
	em.lease = clientv3.NewLease(EtcdClient)
	//设置一个ctx取消自动续租
	ctx, em.cancelFunc = context.WithCancel(context.TODO())
	//设置租约时间（过期时间）
	leaseResp, err := em.lease.Grant(context.TODO(), em.TTL)
	if err != nil {
		return err
	}
	// 拿到租约id
	em.leaseID = leaseResp.ID

	//自动续租（不停地往管道中扔租约信息）
	leaseRespChan, err := em.lease.KeepAlive(ctx, em.leaseID)
	//启动一个协程去监听
	go listenLeaseChan(leaseRespChan)
	return err
}

func (em *EtcdMutex) Lock() error {
	err := em.init()
	if err != nil {
		return err
	}
	//LOCK:
	em.txn.If(clientv3.Compare(clientv3.CreateRevision(em.Key), "=", 0)).
		Then(clientv3.OpPut(em.Key, "", clientv3.WithLease(em.leaseID))).
		Else(clientv3.OpGet(em.Key))

	//提交事务
	txnResp, err := em.txn.Commit()
	if err != nil {
		return err
	}
	if !txnResp.Succeeded { //判断txn.if条件是否成立
		return fmt.Errorf("抢锁失败: %v", txnResp)
	}
	return nil
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
				goto END
			} else {
				fmt.Println(leaseKeepResp)
			}
		}
	}
END:
}
func (em *EtcdMutex) UnLock() {
	em.cancelFunc()
	_, err := em.lease.Revoke(context.TODO(), em.leaseID)
	if err != nil {
		panic(err)
	}
	fmt.Println("释放了锁")
}
