package cmd

import (
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

type cmdDocs struct {
	host *string
}

func NewCmdDocs() Cmd {
	return &cmdDocs{}
}

func (c *cmdDocs) Command() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = "docs"
	cmd.Short = "opens the API docs in the default browser."
	cmd.Long = "Opens the swagger API generated docs in the default browser."
	cmd.RunE = c.Run

	// Hostname
	c.host = cmd.Flags().String("host", "localhost:8080", "The hostname of the server to connect to, this must include the port")

	cmd.Args = cobra.ExactArgs(0)

	return cmd
}

func (c *cmdDocs) Run(cmd *cobra.Command, args []string) error {
	open("http://" + *c.host + "/swagger/index.html")
	return nil
}

// Gotten from: https://gist.github.com/sevkin/9798d67b2cb9d07cb05f89f14ba682f8
// https://stackoverflow.com/questions/39320371/how-start-web-server-to-open-page-in-browser-in-golang
// open opens the specified URL in the default browser of the user.
func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
