package main

import (
	. "github.com/spf13/cobra"
	"os/exec"
)

func cmdLink(rootCmd *Command) {
	rootCmd.AddCommand(&Command{
		Use:   "link [markdown file]",
		Short: "Link markdown file to Kobito",
		Long:  "Link markdown file to Kobito",
		Run: func(cmd *Command, args []string) {
			if len(args) != 1 {
				fatal("arg 1 is required")
			}

			filename := args[0]

			_, err := exec.Command("open", "-a", "Kobito", filename).Output()
			failOnError(err)
		},
	})
}
