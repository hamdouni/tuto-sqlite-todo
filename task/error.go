package task

import "errors"

var (
	ErrEmptyTask               = errors.New("empty task not allowed")
	ErrRepositoryNotConfigured = errors.New("task repository not defined")
)
