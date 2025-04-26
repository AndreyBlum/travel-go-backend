package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"travel-go/backend/internal/travel"
	"travel-go/backend/internal/travel/repository"
	"travel-go/backend/internal/travel/service"
)

type TravelHandler struct {
	Service *service.TravelService
}

func NewTravelHandler(db *sql.DB) *TravelHandler {
	repo := repository.NewPostgresTravelRepository(db)
	svc := service.NewTravelService(repo)
	return &TravelHandler{Service: svc}
}

func (TravelHandler *TravelHandler) GetAllTravels(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(writer, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	defer request.Body.Close()

	allTavels, err := TravelHandler.Service.GetTravel()

	if err != nil {
		http.Error(writer, "Erro ao buscar a viagem", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(allTavels)
}

func (travelHandler *TravelHandler) CreateTravel(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var payload travel.Travel
	if err := json.NewDecoder(request.Body).Decode(&payload); err != nil {
		http.Error(writer, "JSON inválido", http.StatusBadRequest)
		return
	}
	defer request.Body.Close()

	if err := travelHandler.Service.CreateTravel(&payload); err != nil {
		http.Error(writer, "Erro ao criar viagem", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte("Viagem criada com sucesso"))
}
