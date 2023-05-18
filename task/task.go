// Le package "task" définit la couche métier.
// Il définit les fonctionnalités propre au métier et les interactions
// éventuelles avec l'extérieur.
package task

import "fmt"

// Status définit les différents états possibles d'une tâche.
type Status int

const (
	StatusOpened Status = iota
	StatusClosed
)

// Item est notre structure principale. C'est notre modèle de donnée pour
// stocker une tâche.
type Item struct {
	ID          int64
	Description string
	State       Status
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

// Create crée une tâche à partir d'une description
func Create(desc string) (int64, error) {
	if desc == "" {
		return 0, ErrEmptyTask
	}
	if config.repo == nil {
		return 0, ErrRepositoryNotDefined
	}
	id, err := config.repo.Create(Item{
		Description: desc,
		State:       StatusOpened,
	})
	if err != nil {
		return 0, fmt.Errorf("creating task '%s': %s", desc, err)
	}
	return id, nil
}

// Close passe le statut d'une tâche à "fermé"
func Close(id int64) error {
	it, err := config.repo.GetByID(id)
	if err != nil {
		return fmt.Errorf("closing task id %d: %s", id, err)
	}
	it.State = StatusClosed
	return config.repo.Update(it)
}

// Get retourne la tâche avec l'id spécifié
func Get(id int64) (it Item, err error) {
	it, err = config.repo.GetByID(id)
	if err != nil {
		return it, fmt.Errorf("getting task id %d: %s", id, err)
	}
	return it, nil
}

// GetAll retourne toutes les tâches
func GetAll() (its []Item, err error) {
	if config.repo == nil {
		return its, ErrRepositoryNotDefined
	}
	return config.repo.GetAll(), nil
}

// GetAllOpened retourne toutes les tâches ouvertes
func GetAllOpened() (its []Item, err error) {
	if config.repo == nil {
		return its, ErrRepositoryNotDefined
	}
	return config.repo.GetByState(StatusOpened), nil
}

// GetAllClosed retourne toutes les tâches fermées
func GetAllClosed() (its []Item, err error) {
	if config.repo == nil {
		return its, ErrRepositoryNotDefined
	}
	return config.repo.GetByState(StatusClosed), nil
}

// String permet à notre structure d'être facilement affichable
func (it Item) String() string {
	return it.Description
}
