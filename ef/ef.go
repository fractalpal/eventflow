package main

import (
	"context"
	"fmt"
	"github.com/fractalpal/eventflow/ef/cmd"
	"os"
)

func main() {
	appCtx := context.Background()
	rootCmd := cmd.RootCommand(appCtx)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
