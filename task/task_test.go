package task_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/hamdouni/tuto-sqlite-todo/repo/db"
	"github.com/hamdouni/tuto-sqlite-todo/repo/ram"
	"github.com/hamdouni/tuto-sqlite-todo/task"
)

func initRepo(withdata bool) (task.Repository, error) {
	repo, err := initFakeSQLiteRepo()
	// repo, err := initFakeRamRepo()
	if err != nil {
		return &repo, err
	}
	task.Init(&repo)
	if withdata {
		for i := 0; i < 3; i++ {
			task.Create(fmt.Sprintf("write a book %v", i+1))
		}
		for i := 3; i < 6; i++ {
			task.Create(fmt.Sprintf("finish a book %v", i+1))
			task.Close(i + 1)
		}
	}
	all := repo.GetAll()
	for _, item := range all {
		fmt.Printf("%#v\n", item)
	}
	return &repo, nil
}

// Fake repo of 6 tasks : 3 opened and 3 closed
func initFakeRamRepo() (ram.List, error) {
	return ram.List{}, nil
}

func initFakeSQLiteRepo() (store db.Store, err error) {
	dir, err := os.MkdirTemp("", "testdb")
	if err != nil {
		return store, err
	}
	dbpath := filepath.Join(dir, "testfile.db")
	taskRepo, err := db.Open(dbpath, "")
	if err != nil {
		return store, err
	}
	return taskRepo, nil
}

func TestSave(t *testing.T) {
	taskRepo, err := initRepo(true)
	if err != nil {
		t.Fatalf("could not init repo %s", err)
	}
	defer taskRepo.Close()
	id := 2
	item, err := taskRepo.GetByID(id)
	if err != nil {
		t.Fatalf("could not get task id %v: %v", id, err)
	}
	want := "write a book 2"
	got := item.Description
	if got != want {
		t.Fatalf("expected %v got %v", want, got)
	}
}

func TestGetAll(t *testing.T) {
	taskRepo, err := initRepo(true)
	if err != nil {
		t.Fatalf("could not init repo %s", err)
	}
	defer taskRepo.Close()
	allItems := taskRepo.GetAll()
	size := len(allItems)
	if size != 6 {
		t.Fatalf("expected size 6 got %v", size)
	}
	for i := 1; i <= 3; i++ {
		want := fmt.Sprintf("write a book %v", i)
		got := allItems[i-1].Description
		if want != got {
			t.Fatalf("expected item %v to be %v got %v", i, want, got)
		}
	}
	for i := 4; i <= 6; i++ {
		want := fmt.Sprintf("finish a book %v", i)
		got := allItems[i-1].Description
		if want != got {
			t.Fatalf("expected item %v to be %v got %v", i, want, got)
		}
	}
}

func TestGetOpened(t *testing.T) {
	taskRepo, err := initRepo(true)
	if err != nil {
		t.Fatalf("could not init repo %s", err)
	}
	defer taskRepo.Close()
	openedItems := taskRepo.GetByState(task.Opened)
	size := len(openedItems)
	if size != 3 {
		t.Fatalf("expected size 3 got %v", size)
	}
	for i := 1; i <= 3; i++ {
		want := fmt.Sprintf("write a book %v", i)
		got := openedItems[i-1].Description
		if want != got {
			t.Fatalf("expected item %v to be %v got %v", i, want, got)
		}
	}
}

func TestGetClosed(t *testing.T) {
	taskRepo, err := initRepo(true)
	if err != nil {
		t.Fatalf("could not init repo %s", err)
	}
	defer taskRepo.Close()
	closedItems := taskRepo.GetByState(task.Closed)
	size := len(closedItems)
	if size != 3 {
		t.Fatalf("expected size 3 got %v", size)
	}
	for i := 1; i <= 3; i++ {
		want := fmt.Sprintf("finish a book %v", i+3)
		got := closedItems[i-1].Description
		if want != got {
			t.Fatalf("expected item %v to be %v got %v", i+3, want, got)
		}
	}
}

func TestCreateTask(t *testing.T) {
	taskRepo, err := initRepo(false)
	if err != nil {
		t.Fatalf("could not init repo %s", err)
	}
	defer taskRepo.Close()
	want := "Only test the parts of the application that you want to work"
	id, err := task.Create(want)
	if err != nil {
		t.Fatalf("expected create item not fail got %s", err)
	}
	got, err := taskRepo.GetByID(id)
	if err != nil {
		t.Fatalf("expected item id %d exists got %s", id, err)
	}
	if got.Description != want {
		t.Fatalf("expected %s got %s", want, got)
	}
}

func TestCloseTask(t *testing.T) {
	taskRepo, err := initRepo(false)
	if err != nil {
		t.Fatalf("could not init repo %s", err)
	}
	defer taskRepo.Close()
	empty := len(taskRepo.GetAll())
	if 0 != empty {
		t.Fatalf("expected empty repo but got %d", empty)
	}
	want := "The only way to get more done is to have less to do"
	id, err := task.Create(want)
	if err != nil {
		t.Fatalf("expected create item not fail got %s", err)
	}
	closed := len(taskRepo.GetByState(task.Closed))
	if 0 != closed {
		t.Fatalf("expected no closed item in repo got %d", closed)
	}
	task.Close(id)
	closed = len(taskRepo.GetByState(task.Closed))
	if 1 != closed {
		t.Fatalf("expected 1 closed item in repo got %d", closed)
	}
	got := taskRepo.GetByState(task.Closed)[0].Description
	if got != want {
		t.Fatalf("expected closed item to be '%s' got '%s'", want, got)
	}
}
