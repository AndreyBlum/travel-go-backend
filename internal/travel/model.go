package travel

type TravelType string

type Travel struct {
	Id         int        `json:"id"`
	Name       string     `json:"name"`
	Country    string     `json:"country"`
	TravelType TravelType `json:"travel_type"`
}

const (
	PlaneTravel TravelType = "plane"
	BusTravel   TravelType = "bus"
	ShipTravel  TravelType = "ship"
)
