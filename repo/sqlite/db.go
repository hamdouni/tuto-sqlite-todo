package sqlite

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"

	"github.com/hamdouni/tuto-sqlite-todo/task"
)

// Store is a SQLite storage.
// It implements task.repository interface.
type Store struct {
	dbpath   string
	pragma   string
	database *sql.DB
}

func Open(path, params string) (s Store, err error) {
	dsn := fmt.Sprintf("%s?%s", path, params)
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return s, fmt.Errorf("error opening database %s got %s", path, err)
	}
	s.database = db
	s.dbpath = path
	s.pragma = params
	err = s.init_schema()
	return s, err
}

func (s Store) Close() error {
	return s.database.Close()
}

func (s Store) Create(t task.Item) (id int64, err error) {
	// insert row and return last id
	q := "INSERT INTO task(description, state) values(?,?)"
	st, err := s.database.Prepare(q)
	if err != nil {
		return id, err
	}
	res, err := st.Exec(t.Description, t.State)
	if err != nil {
		return id, err
	}
	lastid, err := res.LastInsertId()
	if err != nil {
		return id, err
	}
	return lastid, nil
}

func (s Store) Update(t task.Item) error {
	q := "UPDATE task SET description=?, state=? WHERE id=?"
	st, err := s.database.Prepare(q)
	if err != nil {
		return err
	}
	_, err = st.Exec(t.Description, t.State, t.ID)
	return err
}
func (s Store) GetAll() []task.Item {
	var t []task.Item
	q := "SELECT id,description,state FROM task"
	rows, err := s.database.Query(q)
	if err != nil {
		return t
	}
	var item task.Item
	for rows.Next() {
		err = rows.Scan(&item.ID, &item.Description, &item.State)
		if err != nil {
			continue
		}
		t = append(t, item)
	}
	return t
}
func (s Store) GetByID(id int64) (task.Item, error) {
	var t task.Item
	q := "SELECT id,description,state FROM task WHERE id=?"
	err := s.database.QueryRow(q, id).Scan(&t.ID, &t.Description, &t.State)
	if err != nil {
		return t, err
	}
	return t, nil
}
func (s Store) GetByState(status task.Status) []task.Item {
	var t []task.Item
	q := "SELECT id,description,state FROM task WHERE state=?"
	rows, err := s.database.Query(q, status)
	if err != nil {
		return t
	}
	var item task.Item
	for rows.Next() {
		err = rows.Scan(&item.ID, &item.Description, &item.State)
		if err != nil {
			continue
		}
		t = append(t, item)
	}
	return t
}
