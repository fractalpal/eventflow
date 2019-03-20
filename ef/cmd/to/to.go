package to

import (
	"github.com/spf13/cobra"
	"net/http"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:              "to",
		Short:            "send events to target",
		Long:             `events to Long Description goes here.....`,
		TraverseChildren: true,
	}
	k := KafkaCommand()
	p := PostgresCommand()
	h := HttpCommand(http.Client{})
	cmd.AddCommand(k.Command, p.Command, h.Command)
	return cmd
}
