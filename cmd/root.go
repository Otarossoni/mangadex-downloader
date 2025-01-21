package cmd

import (
	"fmt"

	"github.com/Otarossoni/mangadex-downloader/helper"
	"github.com/Otarossoni/mangadex-downloader/internal/command"
	"github.com/Otarossoni/mangadex-downloader/service"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mangadex-downloader [flags]",
	Short: "Download manga from Mangadex",
	Run: func(cmd *cobra.Command, args []string) {
		cobraHelper := helper.NewCobraHelper()
		mangadexApi := service.NewMangadexApi()
		packer := helper.NewPacker()
		downloadMangaCommand := command.NewDownloadMangaCommand(cobraHelper, mangadexApi, packer)

		errMangaId := downloadMangaCommand.Execute(cmd)
		if errMangaId != nil {
			color.Red(errMangaId.Error())
			return
		}
	},
}

func init() {
	var url, mangaId, chapters, language, extension, outPath string

	rootCmd.Flags().StringVarP(&url, "url", "u", "", "URL of the Manga to be downloaded")
	rootCmd.Flags().StringVarP(&mangaId, "mangaId", "m", "", "ID of the Manga to be downloaded")
	rootCmd.Flags().StringVarP(&chapters, "chapters", "c", "", "Chapters range")
	rootCmd.Flags().StringVarP(&language, "language", "l", "", "Chapter language")
	rootCmd.Flags().StringVarP(&extension, "extension", "e", "", "Extension for chapter pack")
	rootCmd.Flags().StringVarP(&outPath, "outPath", "o", "", "Directory for saving files")
}

func ExecuteRoot() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
