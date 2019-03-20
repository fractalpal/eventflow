package from

import (
	"fmt"
	"github.com/fractalpal/eventflow/ef/cmd/to"
	"github.com/spf13/cobra"
)

type PostgresCmd struct {
	*cobra.Command
	sourceUrl string
}

func PostgresCommand() *PostgresCmd {
	pCmd := &PostgresCmd{}
	cmd := &cobra.Command{
		Use:   "postgres",
		Short: "read events from postgres database",
		Long:  `postgres from Long Description goes here.....`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside from.PostgresCmd PersistentPreRun with args: %v and sourceUrl %s\n", args, pCmd.sourceUrl)
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside from.PostgresCmd PersistentPostRun with args: %v and sourceUrl %s\n", args, pCmd.sourceUrl)
		},
		TraverseChildren: true,
	}
	cmd.PersistentFlags().StringVarP(&pCmd.sourceUrl, "from-url", "", "", "from postgres database source url")
	cmd.AddCommand(to.Cmd())
	pCmd.Command = cmd
	return pCmd
}
