package task_test

import (
	"errors"
	"testing"

	"github.com/hamdouni/tuto-sqlite-todo/task"
)

// ensure the repo is initialized before writing
func TestWritingNoRepo(t *testing.T) {
	_, err := task.Create("create without a repo")
	if !errors.Is(err, task.ErrRepositoryNotConfigured) {
		t.Fatalf("expected error %s got %s", task.ErrRepositoryNotConfigured, err)
	}
}

func TestReadAllFromNoRepo(t *testing.T) {
	_, err := task.GetAll()
	if err != task.ErrRepositoryNotConfigured {
		t.Fatalf("expected error %s got %s", task.ErrRepositoryNotConfigured, err)
	}
}
func TestReadAllOpenedFromNoRepo(t *testing.T) {
	_, err := task.GetAllOpened()
	if err != task.ErrRepositoryNotConfigured {
		t.Fatalf("expected error %s got %s", task.ErrRepositoryNotConfigured, err)
	}
}
func TestReadAllClosedFromNoRepo(t *testing.T) {
	_, err := task.GetAllClosed()
	if err != task.ErrRepositoryNotConfigured {
		t.Fatalf("expected error %s got %s", task.ErrRepositoryNotConfigured, err)
	}
}
