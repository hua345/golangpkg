package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"testing"
)

const Id = "1234567890"

func TestRegister(t *testing.T) {
	config := api.DefaultConfig()
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	//创建一个新服务。
	registration := new(api.AgentServiceRegistration)
	registration.ID = Id
	registration.Name = "user-tomcat"
	registration.Port = 8080
	registration.Tags = []string{"user-tomcat"}
	registration.Address = "127.0.0.1"

	//增加check。
	check := new(api.AgentServiceCheck)
	check.HTTP = fmt.Sprintf("http://%s:%d%s", registration.Address, registration.Port, "/check")
	//设置超时 5s。
	check.Timeout = "5s"
	//设置间隔 5s。
	check.Interval = "5s"
	//注册check服务。
	registration.Check = check
	t.Log("get check.HTTP:", check)

	err = client.Agent().ServiceRegister(registration)

	if err != nil {
		t.Error("register server error : ", err)
	}

}
