package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	redisAddr := flag.String("redis", "localhost:6379", "Redis server address")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		printUsage()
		os.Exit(1)
	}

	store, err := NewTodoStore(*redisAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	command := args[0]

	switch command {
	case "add":
		if len(args) < 2 {
			fmt.Fprintf(os.Stderr, "Error: 'add' requires todo text\n")
			os.Exit(1)
		}
		text := strings.Join(args[1:], " ")
		id, err := store.CreateTodo(text)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Created todo: %s\n", id)

	case "list":
		todos, err := store.ListTodos()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		if len(todos) == 0 {
			fmt.Println("No todos found")
			return
		}
		for _, todo := range todos {
			fmt.Printf("[%s] %s (created: %s)\n", todo.ID, todo.Text, todo.CreatedAt.Format("2006-01-02 15:04:05"))
		}

	case "get":
		if len(args) < 2 {
			fmt.Fprintf(os.Stderr, "Error: 'get' requires todo ID\n")
			os.Exit(1)
		}
		todo, err := store.GetTodo(args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("[%s] %s (created: %s)\n", todo.ID, todo.Text, todo.CreatedAt.Format("2006-01-02 15:04:05"))

	case "update":
		if len(args) < 3 {
			fmt.Fprintf(os.Stderr, "Error: 'update' requires ID and text\n")
			os.Exit(1)
		}
		text := strings.Join(args[2:], " ")
		err := store.UpdateTodo(args[1], text)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Updated todo: %s\n", args[1])

	case "delete":
		if len(args) < 2 {
			fmt.Fprintf(os.Stderr, "Error: 'delete' requires todo ID\n")
			os.Exit(1)
		}
		err := store.DeleteTodo(args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Deleted todo: %s\n", args[1])

	default:
		fmt.Fprintf(os.Stderr, "Error: unknown command %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: todo [--redis=ADDR] <command> [arguments]")
	fmt.Println("\nCommands:")
	fmt.Println("	add <text>           Create a new todo")
	fmt.Println("	list                 List all todos")
	fmt.Println("	get <id>             Get a todo by ID")
	fmt.Println("	updated <id> <text>  Update a todo")
	fmt.Println("	delete <id>          Delete a todo")
}
