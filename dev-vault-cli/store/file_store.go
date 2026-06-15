package store

import (
	"dev-vault-cli/models"
	"encoding/json"
	"os"
)

type FileStore struct {
	filepath string
}

func (f FileStore) load() ([]models.Snippet, error) {
	file, err := os.Open(f.filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var snippets []models.Snippet

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&snippets)
	if err != nil {
		return nil, err
	}

	return snippets, nil
}

func (f FileStore) save(snippets []models.Snippet) error {
	file, err := os.Create(f.filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)

	err = encoder.Encode(snippets)
	if err != nil {
		return err
	}

	return nil
}
