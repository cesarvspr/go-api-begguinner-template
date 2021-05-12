package main_test

import (
	"fmt"
	"testing"

	"github.com/cesarvspr/golang-modules/sdk/client"
	"github.com/cesarvspr/golang-modules/sdk/client/products"
)

func TestOurClient(t *testing.T) {
	// cfg := client.DefaultTransportConfig().WithHost("localhost:9001")
	// c := client.NewHTTPClientWithConfig(nil, cfg)

	// params := products.NewListProductsParams()

	// prod, err := c.Products.ListProducts(params)

	// if err != nil {
	// 	t.Fatal(err)
	// }
	// fmt.Printf("%#v", prod.GetPayload())
	// t.Fail()
	
	cfg := client.DefaultTransportConfig().WithHost("localhost:9001")
	
	c := client.NewHTTPClientWithConfig(nil, cfg)
	
	params := products.NewListProductsParams()
	
	prods, err := c.Products.ListProducts(params)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(prods) 
}
