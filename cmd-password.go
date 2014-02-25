package main

import (
	"fmt"
	. "github.com/spf13/cobra"
	"os/exec"
)

func cmdPassword(rootCmd *Command) {
	rootCmd.AddCommand(&Command{
		Use:   "password",
		Short: "Show Kobito password",
		Long:  "Show Kobito password",
		Run: func(cmd *Command, args []string) {
			output, _ := exec.Command("security", "find-internet-password", "-w", "-s", "com.qiita").Output()
			fmt.Printf("%s", output)
		},
	})
}
