package main

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	todos.add("Go to school")
	todos.add("Get a job")
	todos.add("Drink whisky once")
	todos.toggleCompleted(1)
	todos.print()
	storage.Save(todos)
}
