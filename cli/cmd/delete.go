package cmd

import (
	"errors"
	shared "klutzer/conanical-library-app/shared"
	"net/http"

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

	return cmd
}

func (c *cmdDelete) Run(cmd *cobra.Command, args []string) error {
	return nil
}

type cmdDeleteBook struct {
	httpClient *http.Client

	host     *string
	protocol *string
}

func NewCmdDeleteBook(httpClient *http.Client, host *string, protocol *string) Cmd {
	return &cmdDeleteBook{
		httpClient: httpClient,
		host:       host,
		protocol:   protocol,
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

	return cmd
}

func (c *cmdDeleteBook) Run(cmd *cobra.Command, args []string) error {
	id := args[0]
	if ok := shared.IsValidID(id); !ok {
		return errors.New("the id specified is not a valid id")
	}

	data := shared.CollectionDeleteRequest{
		ID: id,
	}

	url := *c.protocol + "://" + *c.host + "/v1/book"
	err := makeRequest[any](&data, nil, url, http.MethodDelete, c.httpClient)
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

	host     *string
	protocol *string
}

func NewCmdDeleteCollection(httpClient *http.Client, host *string, protocol *string) Cmd {
	return &cmdDeleteCollection{
		httpClient: httpClient,
		host:       host,
		protocol:   protocol,
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

	return cmd
}

func (c *cmdDeleteCollection) Run(cmd *cobra.Command, args []string) error {

	id := args[0]
	if ok := shared.IsValidID(id); !ok {
		return errors.New("the id specified is not a valid id")
	}

	data := shared.CollectionDeleteRequest{
		ID: id,
	}

	url := *c.protocol + "://" + *c.host + "/v1/collection"
	if err := makeRequest[any](&data, nil, url, http.MethodDelete, c.httpClient); err != nil {
		return err
	}

	cmd.Println("Successfully deleted the book collection")

	return nil
}
