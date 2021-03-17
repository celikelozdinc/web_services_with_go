package product

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type (
	barHandler struct {
		Message string
	}
)

// SetupRoutes :
func SetupRoutes() {
	http.Handle("/bar", &barHandler{Message: "bar is called"})
	http.HandleFunc("/baz", bazHandler)
	http.HandleFunc("/products", productsHandler)

	// By using middleware
	productItemHandler := http.HandlerFunc(productHandler)
	http.Handle("/products/", middlewareHandler(productItemHandler)) //=> For parsing URL path parameters
}

func (b *barHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	writer.Write([]byte(b.Message))
}

func bazHandler(writer http.ResponseWriter, req *http.Request) {
	writer.Write([]byte("baz is called")) //=> Actual Response
}

func productsHandler(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		productList := getProductList()
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
	product := getProduct(productID)
	p, marshallErr := json.Marshal(product)
	if marshallErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// If no error is returned
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(p) // => Actual Response
}

func middlewareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		log.Printf("Entering middleware at %s", time.Now())
		handler.ServeHTTP(writer, req)
		log.Printf("Exiting middleware at %s", time.Now())
	})
}
