package main

import (
	"dev-vault-cli/cmd"
	"dev-vault-cli/store"
	"fmt"
	"os"
)

func main() {
	s := store.NewFileStore("snippets.json")
	if len(os.Args) < 2 {
		fmt.Println("Command length too short")
		return
	}
	command := os.Args[1]
	switch command {
	case "add":
		if err := cmd.AddCommand(s, os.Args[2:]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "list":
		if err := cmd.ListCommand(s); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "delete":
		if err := cmd.DeleteCommand(s, os.Args[2:]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "search":
		if err := cmd.SearchCommand(s, os.Args[2:]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Println("Use add/list/delete/search only")

	}
}
