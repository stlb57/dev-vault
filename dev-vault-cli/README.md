# DevVault CLI

A lightweight command-line application written in Go for storing, searching, listing, and deleting code snippets. Snippets are persisted locally in a JSON file, making DevVault a simple personal knowledge base for developers.

## Features

* Add snippets with title, content, tags, and priority
* List all saved snippets
* Search snippets by title, content, or tags
* Delete snippets by ID
* Persistent JSON-based storage
* Custom error handling with sentinel errors
* Unit tests for core functionality

## Project Structure

```text
dev-vault-cli/
├── main.go
├── cmd/
│   ├── add.go
│   ├── list.go
│   ├── search.go
│   └── delete.go
├── store/
│   ├── store.go
│   ├── file_store.go
│   ├── errors.go
│   └── store_test.go
└── models/
    └── snippet.go
```

## Concepts Practiced

This project was built as a learning exercise to practice:

* Structs
* Methods
* Interfaces
* Custom error types
* Error wrapping and handling
* File I/O
* JSON encoding and decoding
* Command-line argument parsing
* Unit testing
* Table-driven testing basics
* Package organization

## Installation

Clone the repository:

```bash
git clone <repository-url>
cd dev-vault-cli
```

Build the application:

```bash
go build
```

## Usage

### Add a snippet

```bash
devvault add --title "Go Concurrency" --content "Worker pools using goroutines" --tags "go,backend,concurrency" --priority 2
```

### List all snippets

```bash
devvault list
```

### Search snippets

```bash
devvault search --query "go"
```

### Delete a snippet

```bash
devvault delete --id "<snippet-id>"
```

## Testing

Run all tests:

```bash
go test ./...
```

Run tests with the race detector:

```bash
go test -race ./...
```

> Note: The race detector requires a supported Go architecture (e.g., amd64).

## Future Improvements

* Update existing snippets
* Export/import functionality
* SQLite storage backend
* Advanced search and filtering
* Sorting and pagination
* Interactive terminal UI

## License

This project was created for learning and educational purposes.
