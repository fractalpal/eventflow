package from

import (
	"fmt"
	"github.com/fractalpal/eventflow/ef/cmd/to"
	"github.com/spf13/cobra"
	"net/http"
)

type HttpCmd struct {
	*cobra.Command
	client http.Client
	port   string
}

func HttpCommand(httpClient http.Client) *HttpCmd {
	httpCmd := &HttpCmd{}
	cmd := &cobra.Command{
		Use:   "http",
		Short: "listen for events on http server",
		Long:  `from http long Description goes here.....`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside from.HttpCmd PersistentPreRun with args: %v and port %s\n", args, httpCmd.port)
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside from.HttpCmd PersistentPostRun with args: %v and port %s\n", args, httpCmd.port)
		},
		TraverseChildren: true,
	}
	cmd.PersistentFlags().StringVarP(&httpCmd.port, "from-port", "", "", "from http server port")
	cmd.AddCommand(to.Cmd())
	httpCmd.Command = cmd

	return httpCmd
}
