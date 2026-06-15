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

func (f FileStore) GetAll() ([]models.Snippet, error) {
	snippets, err := f.load()
	if err != nil {
		return nil, err
	}
	return snippets, nil
}

func (f FileStore) GetByID(id int) (models.Snippet, error) {
	snippets, err := f.load()
	if err != nil {
		return models.Snippet{}, err
	}
	for i := 0; i < len(snippets); i++ {
		if snippets[i].Id == id {
			return snippets[i], nil
		}
	}
	return models.Snippet{}, ErrNotFound
}

func (f FileStore) Delete(id int) error {
	snippets, err := f.load()
	if err != nil {
		return err
	}
	found := false
	var updated []models.Snippet

	for _, snippet := range snippets {
		if snippet.Id == id {
			found = true
			continue
		}
		updated = append(updated, snippet)
	}

	if !found {
		return ErrNotFound
	}

	return f.save(updated)
}

func (f FileStore) Search()
