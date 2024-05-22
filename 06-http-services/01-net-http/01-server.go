package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Cost     float64 `json:"cost"`
	Category string  `json:"category"`
}

var products []Product = []Product{
	{Id: 1, Name: "Pen", Cost: 10, Category: "Stationary"},
	{Id: 2, Name: "Pencil", Cost: 5, Category: "Stationary"},
	{Id: 3, Name: "Marker", Cost: 50, Category: "Stationary"},
}

type Server struct {
}

// implementation of the http.Handler interface
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s  %s\n", r.Method, r.URL.Path)
	switch r.URL.Path {
	case "/":
		fmt.Fprintln(w, "Hello, World!")
	case "/products":
		// fmt.Fprintln(w, "All products data will be served!")
		switch r.Method {
		case http.MethodGet:
			if err := json.NewEncoder(w).Encode(products); err != nil {
				http.Error(w, "unable to process the request", http.StatusInternalServerError)
			}
		case http.MethodPost:
			var newProduct Product
			if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
				http.Error(w, "payload error", http.StatusBadRequest)
				return
			}
			newProduct.Id = len(products) + 1
			products = append(products, newProduct)
			if err := json.NewEncoder(w).Encode(newProduct); err != nil {
				http.Error(w, "unable to process the request", http.StatusInternalServerError)
			}
		}

	case "/customers":
		fmt.Fprintln(w, "All customers data will be served!")
	default:
		http.Error(w, "resource not found", http.StatusNotFound)
	}

}

func main() {
	server := &Server{}
	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Server listening on 8080.....")
}
