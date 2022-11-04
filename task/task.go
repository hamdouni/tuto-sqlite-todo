// Le package "task" définit ma couche métier.
// C'est la partie qui doit définir les fonctionnalités propre à mon métier et
// les interactions éventuelles avec l'extérieur.
package task

// Item est notre structure principale. C'est notre model de donnée pour
// stocker une tâche.
type Item struct {
	ID          int
	Description string
	State       State
}

// State définit les différents états possibles d'une tâche.
// La liste des états possibles sont : Opened, Closed
type State int

const (
	Opened State = iota
	Closed
)

// Repository décrit la façon dont le métier intéragit avec le dépot de données
type Repository interface {
	Save(Item) (ID int, err error)
	Update(Item) error
	GetAll() []Item
	GetByID(ID int) (Item, error)
	GetByState(status State) []Item
	Close() error
}

// config stock la configuration de l'application.
// C'est le bon endroit pour "pluger" des comportements externes au
// métier comme le dépot de données.
// Pas besoin de l'exposer à l'extérieur, donc un c minuscule !
var config = struct {
	repo Repository
}{}

// Init permet de pluger le dépot de données.
func Init(r Repository) {
	config.repo = r
}

// Create crée une tâche à partir d'une description. On s'appuie principalement
// sur la façon de stocker une tâche par le dépot de données (bref, on fait pas
// grand chose au niveau "métier")
func Create(desc string) (int, error) {
	return config.repo.Save(Item{
		Description: desc,
		State:       Opened,
	})
}

// Close passe le statut d'une tâche à "fermé".
// D'un point de vue "métier" :
// 1. récupérer la tâche
// 2. modifier son statut
// 3. sauvegarder la tâche
func Close(ID int) error {
	it, err := config.repo.GetByID(ID)
	if err != nil {
		return err
	}
	it.State = Closed
	return config.repo.Update(it)
}

// String permet à notre structure d'être facilement affichage par la lib
// standard Go (fmt)
func (it Item) String() string {
	return it.Description
}
