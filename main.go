package main

func main() {
	todos := Todos{}
	todos.add("Go to school")
	todos.add("Get a job")
	todos.add("Drink whisky once")
	todos.toggleCompleted(1)
	todos.print()
	todos.delete(2)
}
