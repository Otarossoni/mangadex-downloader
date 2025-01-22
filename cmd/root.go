package cmd

import (
	"fmt"

	"github.com/Otarossoni/mangadex-downloader/helper"
	"github.com/Otarossoni/mangadex-downloader/internal/command"
	"github.com/Otarossoni/mangadex-downloader/service"
	"github.com/fatih/color"
	coloredCobra "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "mangadex-downloader [flags]",
	Short:   "Download manga from Mangadex, compress and make available chapters in .zip and .cbz",
	Example: helper.GetExamplesDescription(),
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

	rootCmd.Flags().StringVarP(&url, "url", "u", "", "URL of the Manga to be downloaded - Mangadex manga URL")
	rootCmd.Flags().StringVarP(&mangaId, "mangaId", "m", "", "ID of the Manga to be downloaded - UUID from Mangadex manga URL")
	rootCmd.Flags().StringVarP(&chapters, "chapters", "c", "", "Chapter range (ranged by \"-\", segmented by \";\") - See the examples")
	rootCmd.Flags().StringVarP(&language, "language", "l", "", "Chapter language accepted by Mangadex - en(default), pt-br, es-la, pl, cs, uk, it, vi, hu, and others")
	rootCmd.Flags().StringVarP(&extension, "extension", "e", "", "Extension for chapter pack - .zip(default) or .cbz")
	rootCmd.Flags().StringVarP(&outPath, "outPath", "o", "", "Directory for saving files - The default path is that of the executable")
}

func ExecuteRoot() {
	coloredCobra.Init(&coloredCobra.Config{
		RootCmd:       rootCmd,
		Headings:      coloredCobra.Green + coloredCobra.Bold,
		ExecName:      coloredCobra.Bold,
		Flags:         coloredCobra.Bold,
		FlagsDescr:    coloredCobra.HiBlue,
		FlagsDataType: coloredCobra.Italic,
	})

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
