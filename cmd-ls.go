package main

import (
	"fmt"
	. "github.com/spf13/cobra"
)

func cmdLs(rootCmd *Command) {
	rootCmd.AddCommand(&Command{
		Use:   "ls",
		Short: "List all items",
		Long:  "List all items",
		Run: func(cmd *Command, args []string) {
			items, err := itemRepository.Items()
			failOnError(err)

			fmt.Println("ITEM_ID    TITLE")
			for _, item := range items {
				fmt.Printf("%7d    %s\n", item.Id, item.Title)
			}
		},
	})
}
