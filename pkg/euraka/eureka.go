package euraka

import "github.com/ArthurHlt/go-eureka-client/eureka"

var eurekaClient *eureka.Client

func NewEurekaClient() {
	eurekaClient = eureka.NewClient([]string{
		"http://127.0.0.1:8080/eureka", //From a spring boot based eureka server
		// add others servers here
	})
}
func RegisterEureka() {
	NewEurekaClient()
	instance := eureka.NewInstanceInfo("localhost", "test", "127.0.0.1", 80, 30, false) //Create a new instance to register
	instance.Metadata = &eureka.MetaData{
		Map: make(map[string]string),
	}
	instance.Metadata.Map["foo"] = "bar"             //add metadata for example
	eurekaClient.RegisterInstance("myapp", instance) // Register new instance in your eureka(s)
}
