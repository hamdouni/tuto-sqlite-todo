package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/hamdouni/tuto-sqlite-todo/repo/db"
	"github.com/hamdouni/tuto-sqlite-todo/task"

	_ "modernc.org/sqlite"
)

const pragma = "_pragma=foreign_keys(1)&_pragma=journal_mode(WAL)&_pragma=synchronous(NORMAL)&_pragma=busy_timeout(8000)&_pragma=journal_size_limit(100000000)"

func main() {

	log.Println("starting...")

	tmpDir, err := os.MkdirTemp("", "temp")
	if err != nil {
		log.Fatalf("could not create a temporary folder for database: %s", err)
	}
	defer os.RemoveAll(tmpDir)

	dbpath := filepath.Join(tmpDir, "testbase.db")
	store, err := db.Open(dbpath, pragma)
	if err != nil {
		log.Fatalf("could not open database %s: %s", dbpath, err)
	}
	defer store.Close()

	task.Init(&store)

	for i := 1; i < 10; i++ {
		t := fmt.Sprintf("faire %d café", i)
		if i > 1 {
			t = t + "s"
		}
		id, err := task.Create(t)
		if err != nil {
			log.Fatalf("could not create task %s", err)
		}
		log.Printf("last id %d", id)
	}

	items := store.GetAll()
	for _, item := range items {
		fmt.Printf("uid: %d description: %s state: %d\n", item.ID, item.Description, item.State)
	}
}
