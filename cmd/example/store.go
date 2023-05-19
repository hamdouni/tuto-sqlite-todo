package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hamdouni/tuto-sqlite-todo/repo/sqlite"
)

// a nice sqlite pragma inspired by project pocketbase
const pragma = "_pragma=foreign_keys(1)&_pragma=journal_mode(WAL)&_pragma=synchronous(NORMAL)&_pragma=busy_timeout(8000)&_pragma=journal_size_limit(100000000)"

// getStore returns a store and a teardown function to clean every thing before quitting
func getStore() (dbs sqlite.Store, teardown func() error, err error) {
	tmpDir, err := os.MkdirTemp("", "temp")
	if err != nil {
		return dbs, teardown, fmt.Errorf("could not create a temporary folder for database: %w", err)
	}

	dbpath := filepath.Join(tmpDir, "testbase.db")
	dbs, err = sqlite.Open(dbpath, pragma)
	if err != nil {
		return dbs, teardown, fmt.Errorf("could not open database: %w", err)
	}
	// the teardown function
	return dbs, func() error {
		err := dbs.Close()
		if err != nil {
			return fmt.Errorf("closing db: %w", err)
		}
		err = os.RemoveAll(tmpDir)
		if err != nil {
			return fmt.Errorf("removing temporary dir %s: %w", tmpDir, err)
		}
		return nil
	}, nil
}
