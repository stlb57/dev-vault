package cmd

import (
	"dev-vault-cli/store"
	"flag"
	"fmt"
)

func SearchCommand(s store.Store, args []string) error {
	fs := flag.NewFlagSet("search", flag.ContinueOnError)

	query := fs.String("query", "", "search query")

	if err := fs.Parse(args); err != nil {
		return err
	}

	if *query == "" {
		return fmt.Errorf("query is required")
	}

	results, err := s.Search(*query)
	if err != nil {
		return err
	}

	for _, snippet := range results {
		fmt.Printf("%+v\n", snippet)
	}

	return nil
}
