package repository

type Authorization struct {
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository() *Repository {
	return &Repository{}
}
