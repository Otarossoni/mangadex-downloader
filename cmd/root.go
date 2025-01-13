package cmd

import (
	"fmt"

	"github.com/Otarossoni/mangadex-downloader/helper"
	"github.com/Otarossoni/mangadex-downloader/internal/command"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mangadex-downloader [flags]",
	Short: "Download manga from Mangadex",
	Run: func(cmd *cobra.Command, args []string) {
		cobraHelper := helper.NewCobraHelper()
		downloadMangaCommand := command.NewDownloadMangaCommand(cobraHelper)

		errMangaId := downloadMangaCommand.Execute(cmd)
		if errMangaId != nil {
			color.Red(errMangaId.Error())
			return
		}
	},
}

func init() {
	rootCmd.Flags().String("url", "", "URL of the Manga to be downloaded")
	rootCmd.Flags().String("mangaId", "", "ID of the Manga to be downloaded")
	rootCmd.Flags().String("chapters", "", "Chapters range")
	rootCmd.Flags().String("language", "", "Chapter language")
}

func ExecuteRoot() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
