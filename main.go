package main

import "fmt"

func main() {
	todos := Todos{}
	todos.add("Go to school")
	todos.add("Get a job")
	todos.add("Marry someone beautiful and wise")
	fmt.Printf("%+v\n", todos)
	todos.delete(2)
	fmt.Printf("%+v", todos)
}
