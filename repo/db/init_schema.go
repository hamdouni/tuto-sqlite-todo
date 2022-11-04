package db

func (s *Store) init_schema() error {
	// schema for task table
	q := `CREATE TABLE IF NOT EXISTS "task" ( 
		"id"          INTEGER PRIMARY KEY AUTOINCREMENT,
		"description" VARCHAR(64) DEFAULT "",
		"state"       INTEGER DEFAULT 0
	)`
	_, err := s.database.Exec(q)
	return err
}
