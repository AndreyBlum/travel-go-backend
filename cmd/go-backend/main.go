package main

import (
	"fmt"
	"net/http"
	"travel-go/backend/internal/routes"

	_ "github.com/lib/pq"
)

func main() {
	router := routes.NewRouter()
	port := 8080
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server listening on http://localhost%s\n", addr)

	httpErr := http.ListenAndServe(addr, router)
	if httpErr != nil {
		panic(httpErr)
	}
}
