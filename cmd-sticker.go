package main

import (
	. "github.com/spf13/cobra"
	"os/exec"
)

func cmdSticker(rootCmd *Command) {
	rootCmd.AddCommand(&Command{
		Use:   "sticker",
		Short: "Do you want Kobito sticker?",
		Long:  "Do you want Kobito sticker?",
		Run: func(cmd *Command, args []string) {
			// ブラウザで開く
			_, err := exec.Command("open", "http://shop.qiita.com/items/173021").Output()
			failOnError(err)
		},
	})
}
