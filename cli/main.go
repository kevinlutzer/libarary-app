package main

import (
	"klutzer/conanical-library-app/cli/cmd"
	"net/http"

	"github.com/spf13/cobra"
)

func main() {

	//
	// Setup Root Command
	//

	rootCmd := &cobra.Command{
		Use: "libraryapp",
	}

	host := rootCmd.Flags().String("host", "localhost:8080", "The host the server is on, including the port")
	protocal := rootCmd.Flags().String("protocol", "http", "The protocol to use when making the request")

	//
	// Create commands
	//

	httpClient := &http.Client{}

	createCmd := cmd.NewCmdCreate().Command()
	createBookCmd := cmd.NewCmdCreateBook(httpClient, host, protocal).Command()
	createCollection := cmd.NewCmdCreateCollection(httpClient, host, protocal).Command()

	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(createBookCmd)
	createCmd.AddCommand(createCollection)

	//
	// Delete Commands
	//

	deleteCmd := cmd.NewCmdDelete().Command()
	deleteBookCmd := cmd.NewCmdDeleteBook(httpClient, host, protocal).Command()
	deleteCollectionCmd := cmd.NewCmdDeleteCollection(httpClient, host, protocal).Command()

	rootCmd.AddCommand(deleteCmd)
	deleteCmd.AddCommand(deleteBookCmd)
	deleteCmd.AddCommand(deleteCollectionCmd)

	//
	// Get Commands
	//

	getCmd := cmd.NewCmdGet().Command()
	getBookCmd := cmd.NewCmdGetBook(httpClient, host, protocal).Command()
	getCollectionCmd := cmd.NewCmdGetCollection(httpClient, host, protocal).Command()

	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getBookCmd)
	getCmd.AddCommand(getCollectionCmd)

	//
	// Update Commands
	//

	updateCmd := cmd.NewCmdUpdate().Command()
	updateBookCmd := cmd.NewCmdUpdateBook(httpClient, host, protocal).Command()

	rootCmd.AddCommand(updateCmd)
	updateCmd.AddCommand(updateBookCmd)

	//
	// Build the command structure
	//

	rootCmd.Execute()
}
