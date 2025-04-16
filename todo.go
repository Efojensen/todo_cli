package main

import (
	"errors"
	"time"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (ptrTodoSl *Todos) add (title string) {
	todo := Todo{
		Title: title,
		Completed: false,
		CreatedAt: time.Now(),
		CompletedAt: nil,
	}

	*ptrTodoSl = append(*ptrTodoSl, todo)
}

func (ptrTodoSl *Todos) validateIndex (index int) error {
	if index < 0 || index > len(*ptrTodoSl) {
		err := errors.New("invalid index")
		return err
	}

	return nil
}

func (ptrTodoSl *Todos) delete (index int) error {
	t := *ptrTodoSl
	if err := ptrTodoSl.validateIndex(index); err != nil {
		return err
	}

	*ptrTodoSl = append(t[:index], t[index+1:]...)
	return nil
}

func (ptrTodos *Todos) toggleCompleted (index int) error {
	if err := ptrTodos.validateIndex(index); err != nil {
		return err
	}

	todos := *ptrTodos

	if !todos[index].Completed {
		completedTime := time.Now()
		todos[index].CompletedAt = &completedTime
	}

	todos[index].Completed = !todos[index].Completed

	return nil
}

func (ptrTodos *Todos) editTodo (index int, task string) error {
	if err := ptrTodos.validateIndex(index); err != nil {
		return err
	}

	todos := *ptrTodos
	todos[index].Title = task

	return nil
}