package etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
)

func waitDelete(ctx context.Context, client *clientv3.Client, key string) error {
	cctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var wr clientv3.WatchResponse
	wch := client.Watch(cctx, key, clientv3.WithRev(int64(3939)))
	for wr = range wch {
		for _, ev := range wr.Events {
			if ev.Type == mvccpb.DELETE {
				return nil
			}
		}
	}
	if err := wr.Err(); err != nil {
		return err
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	return fmt.Errorf("lost watcher waiting for delete")
}
