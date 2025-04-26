package travel

type Repository interface {
	Create(travel *Travel) error
	Get() ([]Travel, error)
}
