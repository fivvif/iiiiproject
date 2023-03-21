package repository

import (
	"github.com/jmoiron/sqlx"
	"iiiproject"
)

type Search interface {
}

type Authorization interface {
	CreateUser(user iiiproject.User) (int, error)
	GetUser(username, password string) (iiiproject.User, error)
}

type TodoList interface {
	Create(userId int, list iiiproject.TodoList) (int, error)
	GetAll(userId int) ([]iiiproject.TodoList, error)
	GetById(userId, listId int) (iiiproject.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input iiiproject.UpdateListInput) error
}

/*
type TodoItem interface {
	Create(listId int, item iiiproject.TodoItem) (int, error)
	GetAll(userId, listId int) ([]iiiproject.TodoItem, error)
	GetById(userId, itemId int) (iiiproject.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input iiiproject.UpdateItemInput) error
}
*/

type Repository struct {
	Authorization
	TodoList
	//TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
