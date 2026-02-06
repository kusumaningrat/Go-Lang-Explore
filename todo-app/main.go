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
			Title:       "Golang Gin Framework",
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

	fmt.Println()

	// Get todo by status
	fmt.Println("====== Get Todo By Status =====")
	getOneTodo, err := todoHandler.GetTodoByStatus(true)
	if err != nil {
		fmt.Println("Error while listing all todos", err)
	}

	getOneTodoResult := utils.FormattedResponse("Todos loaded successfully", getOneTodo)
	fmt.Println(utils.PrettyJson(getOneTodoResult))

	// Get todo by id
	fmt.Println("====== Get Todo By ID =====")
	getOneTodoById, err := todoHandler.GetTodoByID(1)
	if err != nil {
		fmt.Println("Error while listing all todos", err)
	}

	getOneTodoByIdResult := utils.FormattedResponse("Todos loaded successfully", getOneTodoById)
	fmt.Println(utils.PrettyJson(getOneTodoByIdResult))

	// Update todo
	fmt.Println("====== Update Todo =====")

	title := "Update Todo"
	desc := "Update description"
	completed := false

	inputUpdate := handlers.TodoUpdateInput{
		Title:       &title,
		Description: &desc,
		Completed:   &completed,
	}

	updateTodo, err := todoHandler.UpdateTodo(1, inputUpdate)
	if err != nil {
		fmt.Println("Error while updating todo", err)
	}

	updateTodoResult := utils.FormattedResponse("Todo updated successfully", updateTodo)
	fmt.Println(utils.PrettyJson(updateTodoResult))

	newTodos := utils.FormattedResponse("Todos loaded successfully", allTodos)
	fmt.Println(utils.PrettyJson(newTodos))

	// Delete todo
	fmt.Println("====== Deleted Todo =====")
	err = todoHandler.DeleteTodo(1)
	if err != nil {
		fmt.Println("Error while deleting todos", err)
		return
	}

	fmt.Println(utils.FormattedResponse("Todo deleted successfully", true))
	newTodosRecord, _ := todoHandler.GetAllTodos()

	newTodosAfterDelete := utils.FormattedResponse("Todos loaded successfully", newTodosRecord)
	fmt.Println(utils.PrettyJson(newTodosAfterDelete))
}
