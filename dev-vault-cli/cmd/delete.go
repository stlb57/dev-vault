package cmd

import (
	"dev-vault-cli/store"
	"errors"
	"flag"
	"fmt"
)

func DeleteCommand(s store.Store, args []string) error {
	fs := flag.NewFlagSet("delete", flag.ContinueOnError)

	id := fs.String("id", "", "snippet id")

	if err := fs.Parse(args); err != nil {
		return err
	}

	if *id == "" {
		return fmt.Errorf("id is required")
	}

	err := s.Delete(*id)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			return fmt.Errorf("snippet not found")
		}
		return err
	}

	fmt.Println("snippet deleted successfully")
	return nil
}
