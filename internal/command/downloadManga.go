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
		return errMangaId
	}

	languages := c.CobraHelper.HandleLanguages(cmd)

	color.Green(mangaId)

	for _, chapter := range chapters {
		color.Green(fmt.Sprintf("%d", chapter))
	}

	for _, language := range languages {
		color.Green(fmt.Sprintf("%v", language))
	}

	return nil
}
