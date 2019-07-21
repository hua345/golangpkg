package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
	"testing"
)

func TestRegister(t *testing.T) {
	NewConsul()
	serviceAddress := "192.168.137.128"
	servicePort := 8080
	const serviceName = "gin-service"
	const serviceId = "gin-service03"
	//创建一个新服务。
	registration := &api.AgentServiceRegistration{
		ID:      serviceId,
		Name:    serviceName,
		Address: serviceAddress,
		Port:    servicePort,
		Tags:    []string{"gin"},
		Check: &api.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%d%s", serviceAddress, servicePort, "/ping"),
			Timeout:  "5s",
			Interval: "2s",
		},
	}

	err := ConsulClient.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal("register server error : ", err)
	}
}

func TestDiscover(t *testing.T) {
	NewConsul()
	services, _, err := ConsulClient.Catalog().Services(&api.QueryOptions{})
	if err != nil {
		t.Error(err)
	}
	for name := range services {
		t.Log(name)
		servicesData, _, err := ConsulClient.Health().Service(name, "", true,
			&api.QueryOptions{})
		if err != nil {
			t.Error(err)
		}
		for _, entry := range servicesData {
			t.Log(entry.Checks.AggregatedStatus())
			t.Log(entry.Service.ID, entry.Service.Service)
			t.Log(entry.Service.Address, entry.Service.Port)
		}
	}
}

func TestDeregister(t *testing.T) {
	const serviceId = "gin-service"
	NewConsul()
	err := ConsulClient.Agent().ServiceDeregister(serviceId)
	if err != nil {
		log.Fatal("register server error : ", err)
	}
}
