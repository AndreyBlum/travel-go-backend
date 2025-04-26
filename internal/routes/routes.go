package routes

import (
	"travel-go/backend/db"
	"travel-go/backend/internal/travel/handler"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	gorila := mux.NewRouter()
	conn := db.InitDB()
	travel_handler := handler.NewTravelHandler(conn)
	gorila.HandleFunc("/travel", travel_handler.CreateTravel).Methods("POST")
	gorila.HandleFunc("/travels", travel_handler.GetAllTravels).Methods("GET")
	return gorila
}
