package task

import "errors"

var (
	ErrEmptyTask            = errors.New("empty task not allowed")
	ErrRepositoryNotDefined = errors.New("task repository not defined")
)
