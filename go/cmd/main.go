package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cobra.EnableCommandSorting = true
	if err := cliCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
