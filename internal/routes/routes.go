package routes

import (
	"fmt"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/api/data", apiDataHandler)

	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the webpage!")
}

func apiDataHandler(w http.ResponseWriter, r *http.Request) {
	data := "Data from API"
	fmt.Fprint(w, data)
}
