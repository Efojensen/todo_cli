package main

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (ptrTodos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	*ptrTodos = append(*ptrTodos, todo)
}

func (ptrTodos *Todos) validateIndex(index int) error {
	if index < 0 || index > len(*ptrTodos) {
		err := errors.New("invalid index")
		return err
	}

	return nil
}

func (ptrTodos *Todos) delete(index int) error {
	t := *ptrTodos
	if err := ptrTodos.validateIndex(index); err != nil {
		return err
	}

	*ptrTodos = append(t[:index], t[index+1:]...)
	return nil
}

func (ptrTodos *Todos) toggleCompleted(index int) error {
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

func (ptrTodos *Todos) editTodo(index int, task string) error {
	if err := ptrTodos.validateIndex(index); err != nil {
		return err
	}

	todos := *ptrTodos
	todos[index].Title = task

	return nil
}

func (ptrTodos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Task", "Completed", "Created At", "Completed At")
	for index, todo := range *ptrTodos {
		completed := "❌"
		completedAt := ""

		if todo.Completed {
			completed = "✅"
			if todo.CompletedAt != nil {
				completedAt = todo.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(index), todo.Title, completed, todo.CreatedAt.Format(time.RFC1123), completedAt)
	}

	table.Render()
}
