package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	getCommand := flag.NewFlagSet("get", flag.ExitOnError)

	createCommand := flag.NewFlagSet("create", flag.ExitOnError)

	deleteCommand := flag.NewFlagSet("delete", flag.ExitOnError)

	updateCommand := flag.NewFlagSet("update", flag.ExitOnError)
	if len(os.Args) < 2 {
		fmt.Println("expected 'create' or 'get' or 'delete' or 'update' subcommands")
		os.Exit(1)
	}

	// Check which subcommand is invoked.
	switch os.Args[1] {

	// For every subcommand, we parse its own flags and
	// have access to trailing positional arguments.
	case "get":
		getCommand.Parse(os.Args[2:])
		fmt.Println("Tasks List")
		getTask()
	case "create":
		createCommand.Parse(os.Args[2:])
		fmt.Println("Creating Task")
		if len(createCommand.Args()) < 2 {
			fmt.Println("expected 'Task Name' and 'Status'")
			os.Exit(1)
		}
		createTask(createCommand.Args())
		getTask()
	case "delete":
		deleteCommand.Parse(os.Args[2:])
		fmt.Println("Deleting Task'")
		if len(deleteCommand.Args()) < 1 {
			fmt.Println("expected 'Task Name'")
			os.Exit(1)
		}
		deleteTask(deleteCommand.Args())
		getTask()
	case "update":
		updateCommand.Parse(os.Args[2:])
		fmt.Println("subcommand 'update'")
		if len(updateCommand.Args()) < 2 {
			fmt.Println("expected 'Task Name' and 'Status'")
			os.Exit(1)
		}
		updateTask(updateCommand.Args())
		getTask()
	default:
		fmt.Println("expected 'create' or 'get' or 'delete' or 'update' subcommands")
		os.Exit(1)
	}
}
