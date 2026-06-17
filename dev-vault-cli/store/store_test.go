package store

import (
	"dev-vault-cli/models"
	"errors"
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
	err = s.Delete("1")
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.GetByID("1")
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestSearch(t *testing.T) {
	f, err := os.CreateTemp("", "snippets-*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())

	s := NewFileStore(f.Name())

	snippet1 := models.Snippet{
		Id:       "1",
		Title:    "Go Concurrency",
		Content:  "Worker pools using goroutines",
		Tags:     []string{"go", "backend", "concurrency"},
		Priority: models.High,
	}

	snippet2 := models.Snippet{
		Id:       "2",
		Title:    "Redis Cache",
		Content:  "Caching with Redis",
		Tags:     []string{"redis", "database"},
		Priority: models.Medium,
	}

	err = s.Save(snippet1)
	if err != nil {
		t.Fatal(err)
	}

	err = s.Save(snippet2)
	if err != nil {
		t.Fatal(err)
	}

	results, err := s.Search("GO")
	if err != nil {
		t.Fatal(err)
	}

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	if results[0].Title != "Go Concurrency" {
		t.Fatalf("expected Go Concurrency, got %s", results[0].Title)
	}
}
