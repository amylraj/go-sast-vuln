// YOLO comment added for testing purposes
package main

import (
	"format"
	"net/http"
	"log"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(resp, "Hello World!")
	})
	log.Fatal(http.ListenAndServe(":8080", r))
}
