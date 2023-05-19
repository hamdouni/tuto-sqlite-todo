package ram

import (
	"fmt"

	"github.com/hamdouni/tuto-sqlite-todo/task"
)

// List is a simple in memory slice for storing tasks.
// It implements task.repository interface.
type List struct {
	repo []task.Item
}

func (l List) Close() error {
	return nil
}

func (l *List) Create(t task.Item) (ID int64, err error) {
	t.ID = int64(len(l.repo) + 1)
	l.repo = append(l.repo, t)
	return t.ID, nil
}
func (l List) GetAll() ([]task.Item, error) {
	return l.repo, nil
}
func (l List) GetByID(ID int64) (t task.Item, err error) {
	for _, it := range l.repo {
		if it.ID == ID {
			return it, nil
		}
	}
	return t, fmt.Errorf("Could not found ID %d", ID)
}
func (l List) GetByState(st task.Status) ([]task.Item, error) {
	var items []task.Item
	for _, it := range l.repo {
		if it.State == st {
			items = append(items, it)
		}
	}
	return items, nil
}
func (l *List) Update(item task.Item) error {
	l.repo[item.ID-1] = item
	return nil
}
