package task_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/hamdouni/tuto-sqlite-todo/repo/ram"
	"github.com/hamdouni/tuto-sqlite-todo/repo/sqlite"
	"github.com/hamdouni/tuto-sqlite-todo/task"
)

// Fake repo with specified opened and closed tasks
func initRepo(opened, closed int) (task.Repository, error) {
	// repo, err := fakeSQLiteRepo()
	repo, err := fakeRamRepo()
	if err != nil {
		return &repo, err
	}
	task.Init(&repo)
	for i := 0; i < opened; i++ {
		task.Create(fmt.Sprintf("write a book %v", i+1))
	}
	for i := 0; i < closed; i++ {
		id, err := task.Create(fmt.Sprintf("finish a book %v", i+1))
		if err != nil {
			return nil, fmt.Errorf("creating task: %s", err)
		}
		err = task.Close(id)
		if err != nil {
			return nil, fmt.Errorf("closing task %d: %s", id, err)
		}
	}
	all := repo.GetAll()
	for _, item := range all {
		fmt.Printf("%#v\n", item)
	}
	return &repo, nil
}

func fakeRamRepo() (ram.List, error) {
	return ram.List{}, nil
}

func fakeSQLiteRepo() (store sqlite.Store, err error) {
	dir, err := os.MkdirTemp("", "testdb")
	if err != nil {
		return store, err
	}
	dbpath := filepath.Join(dir, "testfile.db")
	db, err := sqlite.Open(dbpath, "")
	if err != nil {
		return store, err
	}
	return db, nil
}

// ensure the repo is initialized before writing
func TestWritingNoRepo(t *testing.T) {
	_, err := task.Create("create without a repo")
	if err != task.ErrRepositoryNotDefined {
		t.Fatalf("expected error %s got %s", task.ErrRepositoryNotDefined, err)
	}
}

func TestReadAllFromNoRepo(t *testing.T) {
	_, err := task.GetAll()
	if err != task.ErrRepositoryNotDefined {
		t.Fatalf("expected error %s got %s", task.ErrRepositoryNotDefined, err)
	}
}
func TestReadAllOpenedFromNoRepo(t *testing.T) {
	_, err := task.GetAllOpened()
	if err != task.ErrRepositoryNotDefined {
		t.Fatalf("expected error %s got %s", task.ErrRepositoryNotDefined, err)
	}
}
func TestReadAllClosedFromNoRepo(t *testing.T) {
	_, err := task.GetAllClosed()
	if err != task.ErrRepositoryNotDefined {
		t.Fatalf("expected error %s got %s", task.ErrRepositoryNotDefined, err)
	}
}

func TestGetAll(t *testing.T) {
	db, err := initRepo(1, 1)
	if err != nil {
		t.Fatalf("could not init repo: %s", err)
	}
	defer db.Close()

	items, err := task.GetAll()
	if err != nil {
		t.Fatalf("could not get all tasks: %s", err)
	}
	size := len(items)
	if size != 2 {
		t.Fatalf("expected size 2 got %v", size)
	}
	// item 0
	got := items[0].Description
	want := "write a book 1"
	if want != got {
		t.Fatalf("expected item %s got %s", want, got)
	}
	// item 1
	got = items[1].Description
	want = "finish a book 1"
	if want != got {
		t.Fatalf("expected item %s got %s", want, got)
	}
}

func TestGetOpened(t *testing.T) {
	db, err := initRepo(3, 3)
	if err != nil {
		t.Fatalf("could not init repo %s", err)
	}
	defer db.Close()
	openedItems, err := task.GetAllOpened()
	if err != nil {
		t.Fatalf("could not get all opened tasks: %s", err)
	}
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
	db, err := initRepo(3, 3)
	if err != nil {
		t.Fatalf("could not init repo %s", err)
	}
	defer db.Close()
	closedItems, err := task.GetAllClosed()
	if err != nil {
		t.Fatalf("could not get all closed tasks: %s", err)
	}
	size := len(closedItems)
	if size != 3 {
		t.Fatalf("expected size 3 got %v", size)
	}
	for i := 1; i <= 3; i++ {
		want := fmt.Sprintf("finish a book %v", i)
		got := closedItems[i-1].Description
		if want != got {
			t.Fatalf("expected item %v to be %v got %v", i, want, got)
		}
	}
}

func TestCreateTask(t *testing.T) {
	db, err := initRepo(3, 3)
	if err != nil {
		t.Fatalf("could not init repo %s", err)
	}
	defer db.Close()

	want := "Only test the parts of the application that you want to work"

	id, err := task.Create(want)
	if err != nil {
		t.Fatalf("expected create item not fail got %s", err)
	}
	got, err := task.Get(id)
	if err != nil {
		t.Fatalf("expected item id %d exists got %s", id, err)
	}
	if got.Description != want {
		t.Fatalf("expected %s got %s", want, got)
	}
}

func TestEmptyTask(t *testing.T) {
	db, err := initRepo(0, 0)
	if err != nil {
		t.Fatalf("could not init repo %s", err)
	}
	defer db.Close()
	_, err = task.Create("")
	if err != task.ErrEmptyTask {
		t.Fatalf("expecting error %s got %s", task.ErrEmptyTask, err)
	}
}

func TestCloseTask(t *testing.T) {

	// empty repo
	db, err := initRepo(0, 0)
	if err != nil {
		t.Fatalf("could not init repo %s", err)
	}
	defer db.Close()

	taskDescription := "The only way to get more done is to have less to do"

	id, err := task.Create(taskDescription)
	if err != nil {
		t.Fatalf("expected create item not fail got %s", err)
	}

	closed, err := task.GetAllClosed()
	if err != nil {
		t.Fatalf("expected get all closed not fail: %s", err)
	}
	numberClosed := len(closed)
	if 0 != numberClosed {
		t.Fatalf("expected no closed item in repo got %d", numberClosed)
	}

	task.Close(id)
	closed, err = task.GetAllClosed()
	if err != nil {
		t.Fatalf("expected get all closed not fail: %s", err)
	}
	numberClosed = len(closed)
	if 1 != numberClosed {
		t.Fatalf("expected 1 closed item in repo got %d", numberClosed)
	}
	got := closed[0].Description
	if got != taskDescription {
		t.Fatalf("expected closed item to be '%s' got '%s'", taskDescription, got)
	}
}
