package command_test

import (
	"testing"

	"github.com/Otarossoni/mangadex-downloader/helper"
	"github.com/Otarossoni/mangadex-downloader/internal/command"
	"github.com/Otarossoni/mangadex-downloader/service"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_DownloadMangaCommand(t *testing.T) {
	t.Run("it should be able to download a manga chapter", func(t *testing.T) {
		cmd := &cobra.Command{}
		cmd.Flags().String("url", "", "Manga URL")
		cmd.Flags().String("chapters", "", "Chapters to download")
		cmd.Flags().String("language", "", "Language")
		cmd.Flags().String("extension", "", "File extension")
		cmd.Flags().String("outPath", "", "Output path")

		cobraHelper := helper.NewCobraHelper()
		mangadexApi := service.NewMangadexApi()
		packer := helper.NewPacker()
		downloadMangaCommand := command.NewDownloadMangaCommand(cobraHelper, mangadexApi, packer)

		tempDir := t.TempDir()

		cmd.Flags().Set("url", "https://mangadex.org/title/6a1d1cb1-ecd5-40d9-89ff-9d88e40b136b/tokyo-ghoul")
		cmd.Flags().Set("chapters", "1")
		cmd.Flags().Set("language", "en")
		cmd.Flags().Set("extension", ".zip")
		cmd.Flags().Set("outPath", tempDir)

		err := downloadMangaCommand.Execute(cmd)

		assert.NoError(t, err)
	})

	t.Run("it should not be able to download a manga chapter with wrong flags", func(t *testing.T) {
		cmd := &cobra.Command{}
		cmd.Flags().String("url", "", "Manga URL")
		cmd.Flags().String("chapters", "", "Chapters to download")
		cmd.Flags().String("language", "", "Language")
		cmd.Flags().String("extension", "", "File extension")
		cmd.Flags().String("outPath", "", "Output path")

		cobraHelper := helper.NewCobraHelper()
		mangadexApi := service.NewMangadexApi()
		packer := helper.NewPacker()
		downloadMangaCommand := command.NewDownloadMangaCommand(cobraHelper, mangadexApi, packer)

		cmd.Flags().Set("url", "")

		err := downloadMangaCommand.Execute(cmd)

		assert.Error(t, err)
	})
}
