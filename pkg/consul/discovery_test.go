package consul

import "github.com/hashicorp/consul/api"

var ConsulClient *api.Client

func NewConsul() {
	var err error
	// Get a new client
	ConsulClient, err = api.NewClient(&api.Config{
		Address: "192.168.137.128:8500",
		Scheme:  "http",
	})

	if err != nil {
		panic(err)
	}
}
