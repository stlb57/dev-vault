package cmd

import (
	"dev-vault-cli/models"
	"dev-vault-cli/store"
	"flag"
	"fmt"
	"strings"
)

func AddCommand(s store.Store, args []string) error {
	fs := flag.NewFlagSet("add", flag.ContinueOnError)

	title := fs.String("title", "", "snippet title")
	content := fs.String("content", "", "snippet content")
	tags := fs.String("tags", "", "comma separated tags")
	priority := fs.Int("priority", 0, "0=low,1=medium,2=high")

	if err := fs.Parse(args); err != nil {
		return err
	}

	if *title == "" {
		return fmt.Errorf("title is required")
	}

	if *content == "" {
		return fmt.Errorf("content is required")
	}

	var tagList []string
	if *tags != "" {
		tagList = strings.Split(*tags, ",")
	}

	snippet := models.Snippet{
		Title:    *title,
		Content:  *content,
		Tags:     tagList,
		Priority: models.Priority(*priority),
	}

	if err := s.Save(snippet); err != nil {
		return err
	}

	fmt.Println("snippet added successfully")
	return nil
}
