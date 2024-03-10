package service

import "github.com/XATAB1CH/todo-app/pkg/repository"

type Authorization struct {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
