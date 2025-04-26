package repository

import (
	"database/sql"
	"travel-go/backend/internal/travel"
)

type PostgresTravelRepository struct {
	DB *sql.DB
}

func NewPostgresTravelRepository(db *sql.DB) *PostgresTravelRepository {
	return &PostgresTravelRepository{DB: db}
}

func (r *PostgresTravelRepository) Create(travel *travel.Travel) error {
	_, err := r.DB.Exec("INSERT INTO travel (name, country, travel_type) VALUES ($1, $2, $3)", travel.Name, travel.Country, travel.TravelType)
	return err
}

func (r *PostgresTravelRepository) Get() ([]travel.Travel, error) {
	rows, err := r.DB.Query("SELECT * FROM travel")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var travels []travel.Travel

	for rows.Next() {
		var t travel.Travel
		if err := rows.Scan(&t.Id, &t.Name, &t.Country, &t.TravelType); err != nil {
			return nil, err
		}
		travels = append(travels, t)
	}
	return travels, nil
}
