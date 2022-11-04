package task

type State int

const (
	Opened State = iota
	Closed
)

type Item struct {
	ID          int
	Description string
	State       State
}

type Repository interface {
	Save(Item) (ID int, err error)
	Update(Item) error
	GetAll() []Item
	GetByID(ID int) (Item, error)
	GetByState(status State) []Item
	Close() error
}

var config = struct {
	repo Repository
}{}

func Init(r Repository) {
	config.repo = r
}

func Create(desc string) (int, error) {
	return config.repo.Save(Item{
		Description: desc,
		State:       Opened,
	})
}

func Close(ID int) error {
	it, err := config.repo.GetByID(ID)
	if err != nil {
		return err
	}
	it.State = Closed
	return config.repo.Update(it)
}

func (it Item) String() string {
	return it.Description
}
