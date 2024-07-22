package main

import (
	"fmt"
	"go_microservice/sdk/client"
	"go_microservice/sdk/client/products"
	"testing"
)

func TestOurClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:7000")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	params := products.NewListProductsParams()
	prod, err := c.Products.ListProducts(params)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v", prod.GetPayload()[0])
	t.Fail()
}
