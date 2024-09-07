package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: todo <command>")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "add":
		addTask()
	case "list":
		listTasks()
	default:
		fmt.Println("Undefined command.")
	}
}

func addTask() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: todo add <task>")
		os.Exit(1)
	}
	fmt.Println("Add task:", os.Args[2:])
}

func listTasks() {
	fmt.Println("All tasks:")
}
