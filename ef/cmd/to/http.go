package to

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

type HttpCmd struct {
	*cobra.Command
	client http.Client
	toUrl  string
}

func HttpCommand(httpClient http.Client) *HttpCmd {
	httpCmd := &HttpCmd{}
	cmd := &cobra.Command{
		Use:   "http",
		Short: "send events using http post requests",
		Long:  `to http Long Description goes here.....`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside to.HttpCmd Run with args: %v and toUrl %s\n", args, httpCmd.toUrl)
		},
		TraverseChildren: true,
	}
	cmd.PersistentFlags().StringVarP(&httpCmd.toUrl, "to-url", "", "", "to http server toUrl")
	httpCmd.Command = cmd
	httpCmd.client = httpClient

	return httpCmd
}
