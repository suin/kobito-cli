package main

import (
	"fmt"
	. "github.com/spf13/cobra"
)

const (
	KOBITO_CLI_VERSION = "__KOBITO_CLI_VERSION__"
)

func cmdVersion(rootCmd *Command) {
	rootCmd.AddCommand(&Command{
		Use:   "version",
		Short: "Print kobito cli tools version",
		Long:  "Print kobito cli tools version",
		Run: func(cmd *Command, args []string) {
			fmt.Println("kobito clit tools version " + KOBITO_CLI_VERSION)
		},
	})
}
