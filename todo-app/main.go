package main

import (
	"fmt"
	"todo-app/handlers"
	"todo-app/utils"
)

func main() {

	todoHandler := &handlers.TodoHandler{}

	// Add todo
	fmt.Println("====== Add Todo =====")
	todo1, err := todoHandler.CreateTodo(
		handlers.Todo{
			Title:       "Basic Golang",
			Description: "Learn basic programming on golang",
			Completed:   false,
		},
	)
	// Check if inser operation failed
	if err != nil {
		fmt.Println("Error while creating a todo", err)
	}

	resultTodo1 := utils.FormattedResponse("Todo created successfully", todo1)

	fmt.Println(utils.PrettyJson(resultTodo1))

	todo2, err := todoHandler.CreateTodo(
		handlers.Todo{
			Title:       "Backend Development on Golang",
			Description: "Basic backend development on Golang with Gin Framework",
			Completed:   true,
		},
	)
	// Check if inser operation failed
	if err != nil {
		fmt.Println("Error while creating a todo", err)
	}

	resultTodo2 := utils.FormattedResponse("Todo created successfully", todo2)

	fmt.Println(utils.PrettyJson(resultTodo2))

	todo3, err := todoHandler.CreateTodo(
		handlers.Todo{
			Description: "Basic backend development on Golang with Gin Framework",
			Completed:   true,
		},
	)
	// Check if inser operation failed
	if err != nil {
		fmt.Println("Error while creating a todo", err)
	}

	resultTodo3 := utils.FormattedResponse("Todo created successfully", todo3)

	fmt.Println(utils.PrettyJson(resultTodo3))

	fmt.Println()

	// Get All todos
	fmt.Println("====== Get All Todos =====")
	allTodos, err := todoHandler.GetAllTodos()
	if err != nil {
		fmt.Println("Error while listing all todos", err)
	}

	todosResult := utils.FormattedResponse("Todos loaded successfully", allTodos)
	fmt.Println(utils.PrettyJson(todosResult))
}
