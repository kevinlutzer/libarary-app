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

	//
	// Create commands
	//

	httpClient := &http.Client{}

	createCmd := cmd.NewCmdCreate().Command()
	createBookCmd := cmd.NewCmdCreateBook(httpClient).Command()
	createCollection := cmd.NewCmdCreateCollection(httpClient).Command()

	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(createBookCmd)
	createCmd.AddCommand(createCollection)

	//
	// Delete Commands
	//

	deleteCmd := cmd.NewCmdDelete().Command()
	deleteBookCmd := cmd.NewCmdDeleteBook(httpClient).Command()
	deleteCollectionCmd := cmd.NewCmdDeleteCollection(httpClient).Command()

	rootCmd.AddCommand(deleteCmd)
	deleteCmd.AddCommand(deleteBookCmd)
	deleteCmd.AddCommand(deleteCollectionCmd)

	//
	// Get Commands
	//

	getCmd := cmd.NewCmdGet().Command()
	getBookCmd := cmd.NewCmdGetBook(httpClient).Command()
	getCollectionCmd := cmd.NewCmdGetCollection(httpClient).Command()

	getCmd.AddCommand(getBookCmd)
	getCmd.AddCommand(getCollectionCmd)
	rootCmd.AddCommand(getCmd)

	//
	// Update Commands
	//

	updateCmd := cmd.NewCmdUpdate().Command()
	updateBookCmd := cmd.NewCmdUpdateBook(httpClient).Command()
	updateCollectionCmd := cmd.NewCmdUpdateCollection(httpClient).Command()

	updateCmd.AddCommand(updateBookCmd)
	updateCmd.AddCommand(updateCollectionCmd)
	rootCmd.AddCommand(updateCmd)

	//
	// Build the command structure
	//

	rootCmd.Execute()
}
