package redis

import (
	"github.com/go-redis/redis"
	"runtime"
)

var RedisClient *redis.Client
var ClusterClient *redis.ClusterClient

func NewRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "192.168.137.129:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 20 * runtime.NumCPU(),
	})
}
func NewRedisCluster() {
	ClusterClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"192.168.137.128:6379", "192.168.137.128:6380",
			"192.168.137.128:6381", "192.168.137.128:6382",
			"192.168.137.128:6383", "192.168.137.128:6384"},
	})
	ClusterClient.Ping()
}

// creates a cluster without using cluster mode or Redis Sentinel.
func NewRedisClusterNoClusterMode() {
	// clusterSlots returns cluster slots information.
	// It can use service like ZooKeeper to maintain configuration information
	// and Cluster.ReloadState to manually trigger state reloading.
	clusterSlots := func() ([]redis.ClusterSlot, error) {
		slots := []redis.ClusterSlot{
			// First node with 1 master and 1 slave.
			{
				Start: 0,
				End:   5460,
				Nodes: []redis.ClusterNode{{
					Addr: "192.168.137.128:6384", // master
				}, {
					Addr: "192.168.137.128:6379", // 1st slave
				}},
			},
			// Second node with 1 master and 1 slave.
			{
				Start: 5461,
				End:   10922,
				Nodes: []redis.ClusterNode{{
					Addr: "192.168.137.128:6380", // master
				}, {
					Addr: "192.168.137.128:6382", // 1st slave
				}},
			},
			// Second node with 1 master and 1 slave.
			{
				Start: 10923,
				End:   16383,
				Nodes: []redis.ClusterNode{{
					Addr: "192.168.137.128:6381", // master
				}, {
					Addr: "192.168.137.128:6383", // 1st slave
				}},
			},
		}
		return slots, nil
	}

	ClusterClient = redis.NewClusterClient(&redis.ClusterOptions{
		ClusterSlots:  clusterSlots,
		RouteRandomly: true,
	})
	ClusterClient.Ping()

	// ReloadState reloads cluster state. It calls ClusterSlots func
	// to get cluster slots information.
	err := ClusterClient.ReloadState()
	if err != nil {
		panic(err)
	}
}
