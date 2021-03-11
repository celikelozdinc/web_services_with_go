package main

import "net/http"

type barHandler struct {
	Message string
}

func (b *barHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	writer.Write([]byte(b.Message))
}

func bazHandler(writer http.ResponseWriter, req *http.Request) {
	writer.Write([]byte("baz is called"))
}

func main() {

	http.Handle("/bar", &barHandler{Message: "bar is called"})
	http.HandleFunc("/baz", bazHandler)
	http.ListenAndServe(":1717", nil)
}
