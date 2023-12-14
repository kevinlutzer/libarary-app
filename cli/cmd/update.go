package cmd

import (
	"errors"
	"klutzer/library-app/shared"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

func NewCmdUpdate() Cmd {
	return &cmdUpdate{}
}

type cmdUpdate struct {
}

func (c *cmdUpdate) Command() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = "update [resource]"
	cmd.Short = "Update a resource"
	cmd.Long = "Can update either a book or a collection of books. Additional options can be specified to provide what fields to update with what values."
	cmd.RunE = c.Run

	// Hostname
	cmd.Flags().String("host", "localhost:8080", "The hostname of the server to connect to, this must include the port")

	cmd.Args = cobra.ExactArgs(1)

	return cmd
}

func (c *cmdUpdate) Run(cmd *cobra.Command, args []string) error {
	return nil
}

//
// Update Book
//

func NewCmdUpdateBook(httpClient *http.Client) Cmd {
	return &cmdUpdateBook{
		httpClient: httpClient,
	}
}

type cmdUpdateBook struct {
	httpClient *http.Client

	author      *string
	description *string
	publishedAt *string
	genre       *string
	edition     *uint8

	protocal *string
	host     *string
}

func (c *cmdUpdateBook) Command() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = "book [title]"
	cmd.Short = "Creates a book with the specified title"
	cmd.Long = "Creates a book with the specified title. Additional options can be specified to provide more details about the book."
	cmd.RunE = c.Run
	cmd.Example = "libraryapp update book bdca5e73-fc0c-4ed0-afbf-a6f45a044909 --title=\"The Lord of the Rings\" --author=\"J.R.R. Tolkien\" --description=\"A hobbit goes on an adventure\" --published=\"1954-07-29\" --genre=\"fantasy\" --edition=1"

	// Book specific args
	c.author = cmd.Flags().String("author", "", "The author of the book")
	c.description = cmd.Flags().String("description", "", "A brief description of the book")
	c.publishedAt = cmd.Flags().String("published", "", "The data the book was published")
	c.genre = cmd.Flags().String("genre", "", "The genre of the book, valid genres are: "+shared.ValidGenreStr+"")
	c.edition = cmd.Flags().Uint8("edition", 0, "The edition of the book")

	// Hostname and Protocal
	c.host = cmd.Flags().String("host", "localhost:8080", "The hostname of the server to connect to, this must include the port")
	c.protocal = cmd.Flags().String("protocal", "http://", "The protocal to use when connecting to the server")
	// Book specified [id]
	cmd.Args = cobra.ExactArgs(1)

	return cmd
}

func (c *cmdUpdateBook) Run(cmd *cobra.Command, args []string) error {
	url := *c.protocal + *c.host + "/v1/book"

	data := &shared.BookData{}
	fieldMask := []string{}

	// Author
	if *c.author != "" {
		fieldMask = append(fieldMask, "author")
		data.Author = *c.author
	}

	// Description
	if *c.description != "" {
		fieldMask = append(fieldMask, "description")
		data.Description = *c.description
	}

	// PublishedAt
	if *c.publishedAt != "" {
		tt, err := time.Parse(time.DateOnly, *c.publishedAt)
		if err != nil {
			cmd.Println("The published date provided is not valid, the date must be in the format of YYYY-MM-DD")
			return err
		}
		fieldMask = append(fieldMask, "publishedAt")
		data.PublishedAt = tt
	} else {
		data.PublishedAt = time.Time{}
	}

	// Genre
	if *c.genre != "" {
		fieldMask = append(fieldMask, "genre")
		data.Genre = *c.genre
	}

	// Edition
	if *c.edition > 0 {
		fieldMask = append(fieldMask, "edition")
		data.Edition = *c.edition
	}

	id := args[0]
	if ok := shared.IsValidID(id); !ok {
		return errors.New(shared.InvalidIdMsg)
	}

	req := shared.BookUpdateRequest{
		ID:        id,
		Data:      data,
		FieldMask: fieldMask,
	}

	if err := makeRequest[any](&req, nil, url, http.MethodPost, c.httpClient); err != nil {
		return err
	}

	cmd.Println("Successfully updated book.")

	return nil
}

//
// Update Collection
//

func NewCmdUpdateCollection(httpClient *http.Client) Cmd {
	return &cmdUpdateCollection{
		httpClient: httpClient,
	}
}

type cmdUpdateCollection struct {
	httpClient *http.Client

	name    *string
	bookIDs *[]string

	protocal *string
	host     *string
}

func (c *cmdUpdateCollection) Command() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = "collection id"
	cmd.Short = "Updates a book collection with the specified id"
	cmd.Long = "Updates a book collection with the specified id. Additional options can be specified to provide more details about the book collection."
	cmd.RunE = c.Run
	cmd.Example = "libraryapp update collection bdca5e73-fc0c-4ed0-afbf-a6f45a044909 --name=\"ASD\" --bookid=\"7a46415a-ceda-4631-91ca-000ecf011045\" "

	// Book specific args
	c.name = cmd.Flags().String("name", "", "The name of the collection")
	c.bookIDs = cmd.Flags().StringArray("bookid", []string{}, "The id of a book to add to the collection, this can be specified multiple times")

	// Hostname and Protocal
	c.host = cmd.Flags().String("host", "localhost:8080", "The hostname of the server to connect to, this must include the port")
	c.protocal = cmd.Flags().String("protocal", "http://", "The protocal to use when connecting to the server")
	// Book specified [id]
	cmd.Args = cobra.ExactArgs(1)

	return cmd
}

func (c *cmdUpdateCollection) Run(cmd *cobra.Command, args []string) error {
	url := *c.protocal + *c.host + "/v1/collection"

	data := &shared.CollectionData{}
	fieldMask := []string{}

	// Name
	if *c.name != "" {
		fieldMask = append(fieldMask, "name")
		data.Name = *c.name
	}

	// Book IDs
	if len(*c.bookIDs) > 0 {
		fieldMask = append(fieldMask, "bookIDs")
		data.BookIDs = *c.bookIDs
	}

	id := args[0]
	if ok := shared.IsValidID(id); !ok {
		return errors.New(shared.InvalidIdMsg)
	}

	req := shared.CollectionUpdateRequest{
		ID:        id,
		Data:      data,
		FieldMask: fieldMask,
	}

	if err := makeRequest[any](&req, nil, url, http.MethodPost, c.httpClient); err != nil {
		return err
	}

	cmd.Println("Successfully updated collection.")

	return nil
}
