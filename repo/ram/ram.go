package ram

import (
	"fmt"
	"goExploration/sqlite/task"
)

// List is a simple in memory slice for storing tasks.
// It implements task.repository interface.
type List struct {
	repo []task.Item
}

func (l List) Close() error {
	return nil
}

func (l *List) Save(t task.Item) (ID int, err error) {
	t.ID = len(l.repo) + 1
	l.repo = append(l.repo, t)
	return t.ID, nil
}
func (l List) GetAll() []task.Item {
	return l.repo
}
func (l List) GetByID(ID int) (t task.Item, err error) {
	for _, it := range l.repo {
		if it.ID == ID {
			return it, nil
		}
	}
	return t, fmt.Errorf("Could not found ID %d", ID)
}
func (l List) GetByState(st task.State) []task.Item {
	var items []task.Item
	for _, it := range l.repo {
		if it.State == st {
			items = append(items, it)
		}
	}
	return items
}
func (l *List) Update(item task.Item) error {
	l.repo[item.ID-1] = item
	return nil
}
