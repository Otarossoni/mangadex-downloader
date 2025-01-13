package command

import (
	"fmt"

	"github.com/Otarossoni/mangadex-downloader/internal/entity"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type DownloadMangaCommand struct {
	CobraHelper entity.CobraHelper
}

func NewDownloadMangaCommand(cobraHelper entity.CobraHelper) *DownloadMangaCommand {
	return &DownloadMangaCommand{
		CobraHelper: cobraHelper,
	}
}

func (c *DownloadMangaCommand) Execute(cmd *cobra.Command) error {
	mangaId, errMangaId := c.CobraHelper.HandleMangaId(cmd)
	if errMangaId != nil {
		return errMangaId
	}

	chapters, errChapters := c.CobraHelper.HandleChapters(cmd)
	if errChapters != nil {
		return errChapters
	}

	language, errLanguage := c.CobraHelper.HandleLanguage(cmd)
	if errLanguage != nil {
		return errLanguage
	}

	color.Green(mangaId)

	for _, chapter := range chapters {
		color.Green(fmt.Sprintf("%d", chapter))
	}

	color.Green(language)

	return nil
}
