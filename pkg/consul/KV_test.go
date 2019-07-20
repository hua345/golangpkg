package consul

import (
	"github.com/hashicorp/consul/api"
	"testing"
)

func TestKv(t *testing.T) {
	NewConsul()
	testConsulKey := "testConsulKey"
	testConsulValue := "fangfang"
	// Get a handle to the KV API
	kv := ConsulClient.KV()

	// PUT a new KV pair
	p := &api.KVPair{Key: testConsulKey, Value: []byte(testConsulValue)}
	_, err := kv.Put(p, nil)
	if err != nil {
		panic(err)
	}

	// Lookup the pair
	pair, _, err := kv.Get(testConsulKey, nil)
	if err != nil {
		panic(err)
	}
	t.Log(pair.Key, string(pair.Value))
}
