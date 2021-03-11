package main

import (
	"encoding/json"
	"net/http"
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

func main() {

	http.Handle("/bar", &barHandler{Message: "bar is called"})
	http.HandleFunc("/baz", bazHandler)
	http.HandleFunc("/products", productsHandler)
	http.ListenAndServe(":1717", nil)
}

func init() {
	p1 := Product{
		ProductID:      100,
		Manufacturer:   "ABB",
		Sku:            "100-100-XXYY",
		Upc:            "WTF",
		PricePerUnit:   "Y",
		QuantityOnHand: 700,
		ProductName:    "Product1",
	}

	p2 := Product{
		ProductID:      200,
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
