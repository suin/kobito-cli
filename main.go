package main

import (
	"github.com/spf13/cobra"
	"log"
)

var (
	itemRepository ItemRepository
)

func failOnError(err error) {
	if err != nil {
		fatal(err)
	}
}

func fatal(message interface{}) {
	log.Fatal(message)
}

func main() {
	processStdin()
	rootCmd := &cobra.Command{Use: "kobito"}
	cmdLs(rootCmd)
	cmdShow(rootCmd)
	cmdHtml(rootCmd)
	cmdPrint(rootCmd)
	cmdLink(rootCmd)
	cmdVersion(rootCmd)
	cmdPdf(rootCmd)
	cmdPassword(rootCmd)
	rootCmd.Execute()
}
