package cmd

import (
	"errors"
	shared "klutzer/library-app/shared"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"
)

type cmdDelete struct{}

func NewCmdDelete() Cmd {
	return &cmdDelete{}
}

func (c *cmdDelete) Command() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = "delete"
	cmd.Short = "Deletes a resource"
	cmd.Long = "Can delete either a singular book or a book collection."
	cmd.RunE = c.Run

	// Hostname
	cmd.Flags().String("host", "localhost:8080", "The hostname of the server to connect to, this must include the port")

	cmd.Args = cobra.ExactArgs(2)

	return cmd
}

func (c *cmdDelete) Run(cmd *cobra.Command, args []string) error {
	return nil
}

type cmdDeleteBook struct {
	httpClient *http.Client

	protocal *string
	host     *string
}

func NewCmdDeleteBook(httpClient *http.Client) Cmd {
	return &cmdDeleteBook{
		httpClient: httpClient,
	}
}

func (c *cmdDeleteBook) Command() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = "book id"
	cmd.Short = "Deletes a book with the specified id"
	cmd.Long = "Deletes a book with the specified id"
	cmd.RunE = c.Run
	cmd.Example = "libraryapp delete book f33bc515-84fb-4bad-9c58-89a5a19cd329"

	// The id arg is required
	cmd.Args = cobra.ExactArgs(1)

	// Hostname and Protocal
	c.host = cmd.Flags().String("host", "localhost:8080", "The hostname of the server to connect to, this must include the port")
	c.protocal = cmd.Flags().String("protocal", "http://", "The protocal to use when connecting to the server")
	return cmd
}

func (c *cmdDeleteBook) Run(cmd *cobra.Command, args []string) error {
	id := args[0]
	if ok := shared.IsValidID(id); !ok {
		return errors.New(shared.InvalidIdMsg)
	}

	url := *c.protocal + *c.host + "/v1/book?id=" + url.QueryEscape(id)
	err := makeRequest[any](nil, nil, url, http.MethodDelete, c.httpClient)
	if err != nil {
		return err
	}

	cmd.Println("Successfully deleted the book")

	return nil
}

//
// Delete Collection
//

type cmdDeleteCollection struct {
	httpClient *http.Client

	protocal *string
	host     *string
}

func NewCmdDeleteCollection(httpClient *http.Client) Cmd {
	return &cmdDeleteCollection{
		httpClient: httpClient,
	}
}

func (c *cmdDeleteCollection) Command() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = "collection id"
	cmd.Short = "Deletes a book collection with the specified id"
	cmd.Long = "Deletes a book collection with the specified id"
	cmd.RunE = c.Run
	cmd.Example = "libraryapp delete collection f33bc515-84fb-4bad-9c58-89a5a19cd329"

	// The id arg is required
	cmd.Args = cobra.ExactArgs(1)

	// Hostname and Protocal
	c.host = cmd.Flags().String("host", "localhost:8080", "The hostname of the server to connect to, this must include the port")
	c.protocal = cmd.Flags().String("protocal", "http://", "The protocal to use when connecting to the server")
	return cmd
}

func (c *cmdDeleteCollection) Run(cmd *cobra.Command, args []string) error {

	id := args[0]
	if ok := shared.IsValidID(id); !ok {
		return errors.New(shared.InvalidIdMsg)
	}

	url := *c.protocal + *c.host + "/v1/collection?id=" + url.QueryEscape(id)
	if err := makeRequest[any](nil, nil, url, http.MethodDelete, c.httpClient); err != nil {
		return err
	}

	cmd.Println("Successfully deleted the book collection")

	return nil
}
