package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mangadex-downloader [flags]",
	Short: "Download manga from Mangadex",
	Run: func(cmd *cobra.Command, args []string) {
		color.Green(args[0])
	},
}

func init() {
	rootCmd.Flags().String("url", "", "URL of the Manga to be downloaded")
	rootCmd.Flags().String("mangaId", "", "ID of the Manga to be downloaded")
}

func ExecuteRoot() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
