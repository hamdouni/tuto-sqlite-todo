package task

// Repository décrit la façon dont le métier intéragit avec le dépot de données
type Repository interface {
	Create(Item) (id int64, err error)
	Update(Item) error
	GetAll() ([]Item, error)
	GetByID(id int64) (Item, error)
	GetByState(status Status) ([]Item, error)
	Close() error
}

// ** Defensive Programming **

// DefaultRepository is assigned to task if no repository is configured
type DefaultRepository struct{}

func init() {
	dr := DefaultRepository{}
	config.repo = dr
}

func (r DefaultRepository) Create(i Item) (int64, error) {
	return 0, ErrRepositoryNotConfigured
}
func (r DefaultRepository) Update(i Item) error {
	return ErrRepositoryNotConfigured
}
func (r DefaultRepository) GetAll() ([]Item, error) {
	return []Item{}, ErrRepositoryNotConfigured
}
func (r DefaultRepository) GetByID(id int64) (Item, error) {
	return Item{}, ErrRepositoryNotConfigured
}
func (r DefaultRepository) Close() error {
	return ErrRepositoryNotConfigured
}
func (r DefaultRepository) GetByState(status Status) ([]Item, error) {
	return []Item{}, ErrRepositoryNotConfigured
}
