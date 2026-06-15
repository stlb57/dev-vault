package store

import (
	"dev-vault-cli/models"
)

type Store interface {
	Save(snippet models.Snippet) error
	GetAll() ([]models.Snippet, error)
	GetByID(id string) (*models.Snippet, error)
	Delete(id string) error
	Search(id string) error
}
