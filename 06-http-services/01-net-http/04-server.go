package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type UUID struct {
}

func NewUUID() UUID {
	return UUID{}
}

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
	time.Sleep(10 * time.Second)
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

func TraceMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func main() {

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: http.DefaultServeMux,
	}
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/products", logMiddleware(ProductsHandler))
	http.HandleFunc("/customers", CustomersHandler)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	fmt.Println("Received kill signal")

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(15*time.Second))
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	server.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)

}
