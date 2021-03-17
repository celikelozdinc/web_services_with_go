package main

import (
	"net/http"

	"github.com/celikelozdinc/web_services_with_go/database"
	product "github.com/celikelozdinc/web_services_with_go/product"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	database.SetupDatabase()
	product.SetupRoutes()
	http.ListenAndServe(":1717", nil)
}
