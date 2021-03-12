package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type (
	barHandler struct {
		Message string
	}

	// Product struct
	Product struct {
		ProductID      int    `json:"productId"`
		Manufacturer   string `json:"manufacturer"`
		Sku            string `json:"sku"`
		Upc            string `json:"upc"`
		PricePerUnit   string `json:"pricePerUnit"`
		QuantityOnHand int    `json:"quantityOnHand"`
		ProductName    string `json:"productName"`
	}
)

var (
	productList = []Product{}
)

func (b *barHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	writer.Write([]byte(b.Message))
}

func bazHandler(writer http.ResponseWriter, req *http.Request) {
	writer.Write([]byte("baz is called")) //=> Actual Response
}

func productsHandler(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		log.Printf("productsHandler() :: Will convert => %#v", productList)
		productsInJSON, marshallErr := json.Marshal(productList)
		if marshallErr != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		// If no error is returned
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(productsInJSON) // => Actual Response
	}
}

func productHandler(writer http.ResponseWriter, req *http.Request) {
	urlPathParameters := strings.Split(req.URL.Path, "products/")
	log.Printf("productHandler() :: URL path parameters => %#v", urlPathParameters)
	productID, convErr := strconv.Atoi(urlPathParameters[len(urlPathParameters)-1])
	if convErr != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	// TODO : need a method for finding product by its unique ID
	product, marshallErr := json.Marshal(productList[productID-1])
	if marshallErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// If no error is returned
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(product) // => Actual Response
}

func main() {

	http.Handle("/bar", &barHandler{Message: "bar is called"})
	http.HandleFunc("/baz", bazHandler)
	http.HandleFunc("/products", productsHandler)
	http.HandleFunc("/products/", productHandler) //=> For parsing URL path parameters
	http.ListenAndServe(":1717", nil)
}

func init() {
	p1 := Product{
		ProductID:      1,
		Manufacturer:   "ABB",
		Sku:            "100-100-XXYY",
		Upc:            "WTF",
		PricePerUnit:   "Y",
		QuantityOnHand: 700,
		ProductName:    "Product1",
	}

	p2 := Product{
		ProductID:      2,
		Manufacturer:   "Hitachi",
		Sku:            "200-200-ZZTT",
		Upc:            "WTF",
		PricePerUnit:   "Y",
		QuantityOnHand: 8800,
		ProductName:    "Product2",
	}

	productList = append(productList, p1)
	productList = append(productList, p2)

}
