package main

import (
	"net/http"

	product "github.com/celikelozdinc/web_services_with_go/product"
)

var (
	productList = []product.Product{}
)

func main() {
	product.SetupRoutes()
	http.ListenAndServe(":1717", nil)
}
