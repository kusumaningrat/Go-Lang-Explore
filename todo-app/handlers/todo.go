package handlers

import (
	"errors"
)

type Todo struct {
	ID          int
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type TodoHandler struct {
	idCounter int
	todos     []Todo
}

func (ts *TodoHandler) CreateTodo(t Todo) ([]Todo, error) {
	// Check if title or description is empty
	if t.Title == "" || t.Description == "" {
		return nil, errors.New("title or description must not be empty")
	}

	// Add id (simulation)
	ts.idCounter++
	t.ID = ts.idCounter

	// Append new todo to todos by TodoHandler
	ts.todos = append(ts.todos, t)

	return ts.todos, nil
}

func (ts *TodoHandler) GetAllTodos() ([]Todo, error) {
	return ts.todos, nil
}
