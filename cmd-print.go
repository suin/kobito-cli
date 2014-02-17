package main

import (
	. "github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

func cmdPrint(rootCmd *Command) {
	rootCmd.AddCommand(&Command{
		Use:   "print [item id]",
		Short: "Print out item",
		Long:  "Print out item",
		Run: func(cmd *Command, args []string) {
			if len(args) != 1 {
				fatal("item id is required")
			}

			itemId, err := strconv.Atoi(args[0])
			failOnError(err)

			item, err := itemRepository.ItemOfId(itemId)
			failOnError(err)

			tempDir, err := ioutil.TempDir("/tmp", "kobito")
			failOnError(err)

			// HTML に印刷JSを埋め込む
			html := item.Html + `<script>window.print();</script>`

			// HTML を書き込む
			ioutil.WriteFile(tempDir+"/index.html", []byte(html), os.ModePerm)

			// CSS などをコピーする
			assets := []string{
				"highlight.pack.js",
				"github.min.css",
				"markdown.css",
			}

			for _, asset := range assets {
				err := os.Symlink("/Applications/Kobito.app/Contents/Resources/"+asset, tempDir+"/"+asset)
				failOnError(err)
			}

			// ブラウザで開く
			_, err = exec.Command("open", "-a", "Safari", tempDir+"/index.html").Output()
			failOnError(err)
		},
	})
}
