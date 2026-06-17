package store

import (
	"dev-vault-cli/models"
	"os"
	"testing"
)

func TestSaveGetAll(t *testing.T) {
	f, err := os.CreateTemp("", "snippets-*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())
	s := NewFileStore(f.Name())
	snippet := models.Snippet{
		Id:       "1",
		Title:    "Go Concurrency",
		Content:  "Worker pools using goroutines",
		Tags:     []string{"go", "backend", "concurrency"},
		Priority: models.High,
	}
	err = s.Save(snippet)
	if err != nil {
		t.Fatal(err)
	}
	all, err := s.GetAll()
	if err != nil {
		t.Fatal(err)
	}
	if len(all) != 1 {
		t.Fatalf("expected 1 snippet, got %d", len(all))
	} else {
		t.Log("Test 1 succesful")
	}
}

func TestDelete(t *testing.T) {
	f, err := os.CreateTemp("", "snippets-*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())
	s := NewFileStore(f.Name())
	snippet := models.Snippet{
		Id:       "1",
		Title:    "Go Concurrency",
		Content:  "Worker pools using goroutines",
		Tags:     []string{"go", "backend", "concurrency"},
		Priority: models.High,
	}
	err = s.Save(snippet)
	if err != nil {
		t.Fatal(err)
	}

}
