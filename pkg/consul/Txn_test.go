package consul

import (
	"github.com/hashicorp/consul/api"
	"testing"
)

func TestAPI_ClientTxn(t *testing.T) {
	NewConsul()
	//session := ConsulClient.Session()
	txn := ConsulClient.Txn()
	//// Make a session.
	//id, _, err := session.CreateNoChecks(nil, nil)
	//if err != nil {
	//	t.Fatalf("err: %v", err)
	//}
	//defer session.Destroy(id, nil)

	// Acquire and get the key via a transaction, but don't supply a valid
	// session.
	txnTestKey := "txnTestKey"
	value := []byte("FangFang")
	ops := api.TxnOps{
		&api.TxnOp{
			KV: &api.KVTxnOp{
				Verb:  api.KVSet,
				Key:   txnTestKey,
				Value: value,
			},
		},
		&api.TxnOp{
			KV: &api.KVTxnOp{
				Verb: api.KVGet,
				Key:  txnTestKey,
			},
		},
	}
	ok, ret, _, err := txn.Txn(ops, nil)
	if err != nil {
		t.Fatalf("err: %v", err)
	} else if !ok {
		t.Fatalf("transaction should have failed")
	}
	for _, value := range ret.Results {
		t.Log(value.KV.Key, string(value.KV.Value))
	}
}
