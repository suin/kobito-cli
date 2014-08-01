package main

import (
	. "github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

func cmdPdf(rootCmd *Command) {
	rootCmd.AddCommand(&Command{
		Use:   "pdf [item id] | pdf [item id] [pdf file name]",
		Short: "Save item as PDF",
		Long:  "Save item as PDF",
		Run: func(cmd *Command, args []string) {
			if len(args) < 1 {
				fatal("not enough arguments")
			}

			itemId, err := strconv.Atoi(args[0])
			failOnError(err)

			item, err := itemRepository.ItemOfId(itemId)
			failOnError(err)

			tempDir, err := ioutil.TempDir("/tmp", "kobito")
			failOnError(err)
			defer os.RemoveAll(tempDir)

			// HTML を書き込む
			ioutil.WriteFile(tempDir+"/index.html", []byte(item.Html), os.ModePerm)

			// CSS などをコピーする
			assets := []string{
				"highlight.pack.js",
				"github.min.css",
				"markdown.css",
			}

			kobitoPath, err := kobitoPath()
			failOnError(err)

			for _, asset := range assets {
				err := os.Symlink(kobitoPath+"/Contents/Resources/"+asset, tempDir+"/"+asset)
				failOnError(err)
			}

			// 保存先ファイルを作る
			var pdfFilename string

			if len(args) >= 2 {
				pdfFilename = args[1]
			} else {
				pdfFilename = item.Title + ".pdf"
			}

			pdfFile, err := os.OpenFile(pdfFilename, os.O_WRONLY|os.O_CREATE, 0600)
			failOnError(err)
			defer pdfFile.Close()

			// PDFに変換する
			cupsfilter := exec.Command("/usr/sbin/cupsfilter", tempDir+"/index.html")
			cupsfilter.Stdout = pdfFile
			cupsfilter.Run()
		},
	})
}
