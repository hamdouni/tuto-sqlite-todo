package task

// Repository décrit la façon dont le métier intéragit avec le dépot de données
type Repository interface {
	Create(Item) (id int64, err error)
	Update(Item) error
	GetAll() []Item
	GetByID(id int64) (Item, error)
	GetByState(status Status) []Item
	Close() error
}
