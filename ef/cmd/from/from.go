package from

import (
	"github.com/spf13/cobra"
	"net/http"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:              "from",
		Short:            "fetch events from source",
		Long:             `fetch events from Long Description goes here.....`,
		TraverseChildren: true,
	}
	httpCmd := HttpCommand(http.Client{})
	KafkaCmd := KafkaCommand()
	PostgresCmd := PostgresCommand()
	cmd.AddCommand(KafkaCmd.Command, PostgresCmd.Command, httpCmd.Command)
	return cmd
}
