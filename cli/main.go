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

	rootCmd.Flags().String("host", "localhost:8080", "The host the server is on, including the port")
	rootCmd.Flags().String("protocol", "http", "The protocol to use when making the request")

	//
	// Instiate root commands
	//

	httpClient := &http.Client{}

	createCmd := cmd.NewCmdCreate().Command()
	createBookCmd := cmd.NewCmdCreateBook(httpClient).Command()
	deleteCmd := cmd.NewCmd(httpClient).Command()

	//
	// Build the command structure
	//

	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(createBookCmd)
	rootCmd.Execute()
}
