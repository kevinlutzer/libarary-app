package cmd

import (
	"klutzer/library-app/shared"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

func NewCmdCreate() Cmd {
	return &cmdCreate{}
}

type cmdCreate struct {
}

func (c *cmdCreate) Command() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = "create [resource]"
	cmd.Short = "Creates a resource"
	cmd.Long = "Can create either a book or a collection of books. Additional options can be specified to provide more details about the resource."
	cmd.RunE = c.Run

	// Hostname
	cmd.Flags().String("host", "localhost:8080", "The hostname of the server to connect to, this must include the port")

	cmd.Args = cobra.ExactArgs(2)

	return cmd
}

func (c *cmdCreate) Run(cmd *cobra.Command, args []string) error {
	return nil
}

//
// Create Book
//

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
	id          *string

	host *string
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
	c.genre = cmd.Flags().String("genre", "", "The genre of the book, valid genres are: "+shared.ValidGenreStr+"")
	c.edition = cmd.Flags().Uint8("edition", 1, "The edition of the book")
	c.id = cmd.Flags().String("id", "", "The id of the book, if not provided a new id will be generated")

	// Book specified
	cmd.Args = cobra.ExactArgs(1)

	// Hostname
	c.host = cmd.Flags().String("host", "localhost:8080", "The hostname of the server to connect to, this must include the port")

	return cmd
}

func (c *cmdCreateBook) Run(cmd *cobra.Command, args []string) error {
	url := "http://" + *c.host + "/v1/book"

	tt, err := time.Parse(time.DateOnly, *c.publishedAt)
	if err != nil {
		cmd.Println("The published date provided is not valid, the date must be in the format of YYYY-MM-DD")
		return err
	}

	data := shared.BookPutRequest{
		Title: args[0],
		ID:    *c.id,
		Data: &shared.BookData{
			Author:      *c.author,
			Description: *c.description,
			PublishedAt: tt,
			Genre:       *c.genre,
			Edition:     *c.edition,
		},
	}

	res := shared.ApiResponse[shared.BookPutResponse]{}
	if err := makeRequest[shared.ApiResponse[shared.BookPutResponse]](&data, &res, url, http.MethodPut, c.httpClient); err != nil {
		return err
	}

	cmd.Println("Successfully created book with id: " + res.Data.ID)

	return nil
}

//
// Create Collection
//

func NewCmdCreateCollection(httpClient *http.Client) Cmd {
	return &cmdCreateCollection{
		httpClient: httpClient,
	}
}

type cmdCreateCollection struct {
	httpClient *http.Client

	bookIDs *[]string

	host *string
}

func (c *cmdCreateCollection) Command() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = "collection [title]"
	cmd.Short = "Creates a book collection with the specified title"
	cmd.Long = "Creates a collection book with the specified title. Addtionall you can specify with book ids are apart of the collection"
	cmd.RunE = c.Run
	cmd.Example = "libraryapp create collection \"The Lord of the Rings Trilogy\" --bookid=d95647a8-0c0e-43df-9104-86452accbe8a --bookid=1b2b9c1a-2e4d-4f6f-8b1e-9e9b9c1d2e4d --bookid=3c4c5c6c-5c6c-6c7c-7c8c-8c9c9c1d2e4d"

	// Book specific args
	c.bookIDs = cmd.Flags().StringArray("bookid", []string{}, "The id of a book to add to the collection, this can be specified multiple times")

	// Hostname
	c.host = cmd.Flags().String("host", "localhost:8080", "The hostname of the server to connect to, this must include the port")

	// Book specified
	cmd.Args = cobra.ExactArgs(1)

	return cmd
}

func (c *cmdCreateCollection) Run(cmd *cobra.Command, args []string) error {
	data := shared.CollectionPutRequest{
		Name:    args[0],
		BookIDs: *c.bookIDs,
	}

	if err := data.Validate(); err != nil {
		return err
	}

	url := "http://" + *c.host + "/v1/collection"

	res := shared.ApiResponse[shared.CollectionPutResponse]{}
	if err := makeRequest[shared.ApiResponse[shared.CollectionPutResponse]](&data, &res, url, http.MethodPut, c.httpClient); err != nil {
		return err
	}

	cmd.Println("Successfully created collection of books with id: " + res.Data.ID)

	return nil
}
