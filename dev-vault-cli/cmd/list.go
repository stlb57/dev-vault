package cmd

import (
	"dev-vault-cli/store"
	"fmt"
)

func ListCommand(s store.Store) error {
	snippets, err := s.GetAll()
	if err != nil {
		return err
	}

	for _, snippet := range snippets {
		fmt.Printf("%+v\n", snippet)
	}

	return nil
}
