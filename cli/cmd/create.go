package cmd

import (
	"bytes"
	"encoding/json"
	"klutzer/conanical-library-app/shared"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

func NewCmdCreate() Cmd {
	return &cmdCreate{}
}

type cmdCreate struct{}

func (c *cmdCreate) Command() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = "create [resource]"
	cmd.Short = "Creates a resource"
	cmd.Long = "Can create either a book or a collection of books. Additional options can be specified to provide more details about the resource."
	cmd.RunE = c.Run

	return cmd
}

func (c *cmdCreate) Run(cmd *cobra.Command, args []string) error {
	return nil
}

func NewCmdCreateBook(httpClient *http.Client) Cmd {
	return &cmdCreateBook{
		httpClient: httpClient,
	}
}

type cmdCreateBook struct {
	httpClient *http.Client

	author      *string
	description *string
	publishedAt *string
	genre       *string
	edition     *uint8

	hostname *string
	protocol *string
}

func (c *cmdCreateBook) Command() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = "book [title]"
	cmd.Short = "Creates a book with the specified title"
	cmd.Long = "Creates a book with the specified title. Additional options can be specified to provide more details about the book."
	cmd.RunE = c.Run
	cmd.Example = "libraryapp create book \"The Lord of the Rings\" --author \"J.R.R. Tolkien\" --description \"A hobbit goes on an adventure\" --published 1954-07-29 --genre fantasy --edition 1"

	// Book specific args
	c.author = cmd.Flags().String("author", "", "The author of the book")
	c.description = cmd.Flags().String("description", "", "A brief description of the book")
	c.publishedAt = cmd.Flags().String("published", "1970-01-01", "The data the book was published")
	c.genre = cmd.Flags().String("genre", "", "A brief description of the book")
	c.edition = cmd.Flags().Uint8("edition", 1, "The edition of the book")

	// Book specified
	cmd.Args = cobra.ExactArgs(1)

	// Request specific args
	c.hostname = cmd.Flags().String("host", "localhost:8080", "The host the server is on, including the port")
	c.protocol = cmd.Flags().String("protocol", "http", "The protocol to use when making the request")

	return cmd
}

func (c *cmdCreateBook) Run(cmd *cobra.Command, args []string) error {
	url := *c.protocol + "://" + *c.hostname + "/v1/book"

	tt, err := time.Parse(time.DateOnly, *c.publishedAt)
	if err != nil {
		cmd.Println("The published date provided is not valid, the date must be in the format of YYYY-MM-DD")
	}

	data := shared.BookPutRequest{
		Title: args[0],
		Data: &shared.BookData{
			Author:      *c.author,
			Description: *c.description,
			PublishedAt: tt,
			Genre:       *c.genre,
			Edition:     *c.edition,
		},
		FieldMask: []string{"author", "description", "publishedAt", "genre", "edition"},
	}

	b, err := json.Marshal(data)
	if err != nil {
		cmd.Println("Failed to created book, please try again")
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(b))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		cmd.Println("Failed to created book, please try again" + err.Error())
		return nil
	}
	defer resp.Body.Close()

	return nil
}
