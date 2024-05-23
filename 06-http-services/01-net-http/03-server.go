package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
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

type Middleware func(handler http.HandlerFunc) http.HandlerFunc

type Server struct {
	routes      map[string]http.HandlerFunc
	middlewares []Middleware
}

func (s *Server) addRoute(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	for _, middleware := range s.middlewares {
		handler = middleware(handler)
	}
	s.routes[pattern] = handler
}

func (s *Server) useMiddleware(middleware Middleware) {
	s.middlewares = append(s.middlewares, middleware)
}

// implementation of the http.Handler interface
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler := s.routes[r.URL.Path]; handler == nil {
		http.Error(w, "resource not found", http.StatusNotFound)
	} else {
		handler(w, r)
	}
}

func NewServer() *Server {
	return &Server{
		routes: make(map[string]http.HandlerFunc),
	}
}

// application handlers
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
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
}

func CustomersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "All customers data will be served!")
}

// middlewares
func logMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s	%s\n", r.Method, r.URL.Path)
		handler(w, r)
	}
}

func profileMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler(w, r)
		elapsed := time.Since(start)
		fmt.Println("elapsed :", elapsed)
	}
}

func traceMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		traceId := rand.Intn(10000)
		traceCtx := context.WithValue(r.Context(), "trace-id", traceId)
		reqClone := r.WithContext(traceCtx)
		handler(w, reqClone)
	}
}

func main() {
	// server := &Server{}
	server := NewServer()
	server.useMiddleware(traceMiddleware)
	server.useMiddleware(logMiddleware)
	server.useMiddleware(profileMiddleware)
	server.addRoute("/", IndexHandler)
	server.addRoute("/products", ProductsHandler)
	server.addRoute("/customers", CustomersHandler)
	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Server listening on 8080.....")
}
