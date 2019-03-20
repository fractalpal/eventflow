package to

import (
	"fmt"
	"github.com/spf13/cobra"
)

type PostgresCmd struct {
	*cobra.Command
	toUrl string
}

func PostgresCommand() *PostgresCmd {
	pCmd := &PostgresCmd{}
	cmd := &cobra.Command{
		Use:   "postgres",
		Short: "save events to postgres database",
		Long:  `to postgres Long Description goes here.....`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside to.PostgresCmd Run with args: %v and %s\n", args, pCmd.toUrl)
		},
		TraverseChildren: true,
	}
	cmd.PersistentFlags().StringVarP(&pCmd.toUrl, "to-url", "", "", "to database target url")
	pCmd.Command = cmd
	return pCmd
}
