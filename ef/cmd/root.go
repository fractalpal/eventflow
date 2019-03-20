package cmd

import (
	"context"
	"github.com/fractalpal/eventflow/ef/cmd/from"
	"github.com/spf13/cobra"
)

type Root struct {
	*cobra.Command
	ctx context.Context
}

func RootCommand(ctx context.Context) *Root {
	r := &Root{
		ctx: ctx,
	}
	cmd := &cobra.Command{
		Use:              "ef",
		Short:            "Eventflow event sourcing CLI support tool",
		Long:             `Send events from given source to given target`,
		TraverseChildren: true,
	}
	cmd.Version = "0.1.0"
	cmd.AddCommand(from.Command())
	r.Command = cmd
	return r
}

func (r *Root) Execute() error {
	return r.Command.Execute()
}
