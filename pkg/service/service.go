package service

import (
	"iiiproject"
	"iiiproject/pkg/repository"
)

type Search interface {
}

type Authorization interface {
	CreateUser(user iiiproject.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
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
	Create(userId, listId int, item iiiproject.TodoItem) (int, error)
	GetAll(userId, listId int) ([]iiiproject.TodoItem, error)
	GetById(userId, itemId int) (iiiproject.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input iiiproject.UpdateItemInput) error
}

*/

type Service struct {
	Authorization
	TodoList
	///TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
