package consul

import (
	"github.com/hashicorp/consul/api"
	"testing"
)

func TestA(t *testing.T) {
	//agent.Join("192.168.0.140", false)
	//time.Sleep(50 * time.Second)

	agent.ServiceRegister(&api.AgentServiceRegistration{
		Kind:    "kk",
		ID:      "ee",
		Name:    "ee2",
		Tags:    []string{"tag1", "tag2"},
		Port:    9800,
		Address: "192.168.0.189:9999",
	})
}
