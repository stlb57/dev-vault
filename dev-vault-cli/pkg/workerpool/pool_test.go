package workerpool

import (
	"dev-vault-cli/models"
	"testing"
)

func TestBasicIndexing(t *testing.T) {
	snippets := []models.Snippet{
		{Id: "1", Tags: []string{"go", "backend"}},
		{Id: "2", Tags: []string{"redis", "backend"}},
		{Id: "3", Tags: []string{"docker"}},
	}
	index := IndexSnippets(snippets, 3)
	if len(index["backend"]) != 2 {
		t.Fatalf("expected 2 backend snippets, got %d", len(index["backend"]))
	}

	if len(index["go"]) != 1 {
		t.Fatalf("expected 1 go snippet, got %d", len(index["go"]))
	}

}
