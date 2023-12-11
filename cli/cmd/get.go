package cmd

import (
	"errors"
	"klutzer/conanical-library-app/shared"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

func NewCmdGet() Cmd {
	return &cmdGet{}
}

type cmdGet struct {
}

func (c *cmdGet) Command() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = "get resource"
	cmd.Short = "Gets a list of the resource"
	cmd.Long = "Can either get a list of books or a collection of books. Additional options can be specified to filter the resource"
	cmd.RunE = c.Run

	// Hostname
	cmd.Flags().String("host", "localhost:8080", "The hostname of the server to connect to, this must include the port")

	cmd.Args = cobra.ExactArgs(1)

	return cmd
}

func (c *cmdGet) Run(cmd *cobra.Command, args []string) error {
	return nil
}

//
// Book
//

type cmdGetBook struct {
	httpClient *http.Client

	host *string

	title          *string
	ids            *[]string
	author         *string
	genre          *string
	publishedStart *string
	publishedEnd   *string
}

func NewCmdGetBook(httpClient *http.Client) Cmd {
	return &cmdGetBook{
		httpClient: httpClient,
	}
}

func (c *cmdGetBook) Command() *cobra.Command {
	cmd := &cobra.Command{}

	cmd.Use = "book"
	cmd.Short = "Gets a list of books"
	cmd.Long = "Gets a list of books"
	cmd.RunE = c.Run
	cmd.Example = "libraryapp get book"

	// Book specific args
	c.ids = cmd.Flags().StringArray("id", []string{}, "The id(s) of the books to get can be specified multiple times")
	c.author = cmd.Flags().String("author", "", "The author of the book, this will be a fuzzy match")
	c.genre = cmd.Flags().String("genre", "", "The genre of the book, valid genres are: "+shared.ValidGenreStr+"")
	c.publishedStart = cmd.Flags().String("publishedstart", "", "The start of the published date range of books to get, this must be in the format of 2006-01-02")
	c.publishedEnd = cmd.Flags().String("publishedend", "", "The end of the published date range of books to get, this must in the format of 2006-01-02")

	// Hostname
	c.host = cmd.Flags().String("host", "localhost:8080", "The hostname of the server to connect to, this must include the port")

	return cmd
}

func (c *cmdGetBook) Run(cmd *cobra.Command, args []string) error {
	data := shared.BookGetRequest{
		IDs:    *c.ids,
		Author: *c.author,
		Genre:  *c.genre,
	}

	// Parse/Check Range Start
	if *c.publishedStart != "" {
		t, err := time.Parse(time.DateOnly, *c.publishedStart)
		if err != nil {
			return errors.New("published start must be in the format of 2006-01-02")
		}
		data.RangeStart = t
	}

	// Parse/Check Range End
	if *c.publishedEnd != "" {
		t, err := time.Parse(time.DateOnly, *c.publishedEnd)
		if err != nil {
			return errors.New("published end must be in the format of 2006-01-02")
		}
		data.RangeEnd = t
	}

	resp := shared.ApiResponse[shared.BookLoadResponse]{}

	u := "http://" + *c.host + "/v1/book?" + data.ToQueryStr()

	cmd.Println(u)
	err := makeRequest[shared.ApiResponse[shared.BookLoadResponse]](nil, &resp, u, http.MethodGet, c.httpClient)
	if err != nil {
		return err
	}

	if len(resp.Data.Books) > 0 {
		cmd.Printf("Successfully loaded books:\n\n")
		for _, book := range resp.Data.Books {
			cmd.Printf("- Title: %s, ID: %s, Author: %s, Description: %s, Edition: %o, Genre: %s", book.Title, book.ID, book.Author, book.Description, book.Edition, book.Genre)

			// Make sure the published at is not zero
			if !book.PublishedAt.IsZero() {
				cmd.Printf(", Published: %s", book.PublishedAt.Format(time.RFC3339))
			}

			cmd.Println("\n")
		}
	} else {
		cmd.Printf("No books found matching the query\n")
	}

	return nil
}

//
// Collection
//

type cmdGetCollection struct {
	httpClient *http.Client

	host *string

	includeBooks *bool
}

func NewCmdGetCollection(httpClient *http.Client) Cmd {
	return &cmdGetCollection{
		httpClient: httpClient,
	}
}

func (c *cmdGetCollection) Command() *cobra.Command {
	cmd := &cobra.Command{}

	cmd.Use = "collection"
	cmd.Short = "Gets a list of book collections"
	cmd.Long = "Gets a list of book collections"
	cmd.RunE = c.Run
	cmd.Example = "libraryapp get collection"

	// Additional Argss
	c.includeBooks = cmd.Flags().Bool("includebooks", false, "Include the books in the collection")

	// Hostname
	c.host = cmd.Flags().String("host", "localhost:8080", "The hostname of the server to connect to, this must include the port")

	return cmd
}

func (c *cmdGetCollection) Run(cmd *cobra.Command, args []string) error {

	resp := shared.ApiResponse[shared.CollectionLoadResponse]{}

	u := "http://" + *c.host + "/v1/collection"
	if *c.includeBooks {
		u += "?includebooks=true"
	}

	if err := makeRequest[shared.ApiResponse[shared.CollectionLoadResponse]](nil, &resp, u, http.MethodGet, c.httpClient); err != nil {
		return err
	}

	if len(resp.Data.Collections) > 0 {
		cmd.Printf("Successfully loaded collections:\n\n")
		for _, collection := range resp.Data.Collections {
			cmd.Printf("- Name: %s, ID: %s \n", collection.Name, collection.ID)

			if len(collection.Books) > 0 {
				for j, book := range collection.Books {
					cmd.Printf(" - Title: %s, ID: %s, Author: %s, Description: %s, Edition: %o, Genre: %s", book.Title, book.ID, book.Author, book.Description, book.Edition, book.Genre)

					// Make sure the published at is not zero
					if !book.PublishedAt.IsZero() {
						cmd.Printf(", Published: %s", book.PublishedAt.Format(time.RFC3339))
					}

					if j < len(collection.Books)-1 {
						cmd.Println()
					}
				}
			}
			cmd.Println("\n")
		}
	} else {
		cmd.Printf("No collections found\n")
	}

	return nil
}
