package main

import (
	"fmt"
	. "github.com/spf13/cobra"
	"strconv"
)

func cmdShow(rootCmd *Command) {
	rootCmd.AddCommand(&Command{
		Use:   "show [item id]",
		Short: "Show item",
		Long:  "Show item",
		Run: func(cmd *Command, args []string) {
			if len(args) != 1 {
				fatal("item id is required")
			}

			itemId, err := strconv.Atoi(args[0])
			failOnError(err)

			item, err := itemRepository.ItemOfId(itemId)
			failOnError(err)

			fmt.Print(item.Markdown)
		},
	})
}
