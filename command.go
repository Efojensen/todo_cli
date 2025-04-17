package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Toggle int
	List   bool
	Edit   string
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "Add", "", "Add a new todo; specify title ")
	flag.StringVar(&cf.Edit, "Edit", "", "Edit a todo by index; specify a new title ")
	flag.IntVar(&cf.Del, "Del", -1, "Delete a todo by index ")
	flag.IntVar(&cf.Toggle, "Toggle", -1, "Specify a todo by index to toggle")
	flag.BoolVar(&cf.List, "List", false, "List all todo's")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute (todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Edit != "":
		parts := strings.Split(cf.Edit, ":")

		if len(parts) != 2 {
			fmt.Println("error, invalid format for edit, please use id:nex_title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("error: invalid index for exit")
			os.Exit(1)
		}

		todos.editTodo(index, parts[1])
	case cf.Toggle != -1:
		todos.toggleCompleted(cf.Toggle)
	case cf.Del != -1:
		todos.delete(cf.Del)
	default:
		fmt.Println("No command match")
	}
}