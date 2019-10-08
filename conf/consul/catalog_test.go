package consul

import (
	"github.com/hashicorp/consul/api"
	"testing"
)

func init() {

}

func TestCatalog(t *testing.T) {
	catalog.Register(&api.CatalogRegistration{
		ID:      "catalog_service",
		Node:    "local",
		Address: "dfew",
	}, nil)
}
