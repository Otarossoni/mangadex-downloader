package command

import (
	"fmt"

	"github.com/Otarossoni/mangadex-downloader/internal/entity"
	"github.com/Otarossoni/mangadex-downloader/internal/util"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type DownloadMangaCommand struct {
	CobraHelper entity.CobraHelper
	MangadexApi entity.MangadexApi
}

func NewDownloadMangaCommand(cobraHelper entity.CobraHelper, mangadexApi entity.MangadexApi) *DownloadMangaCommand {
	return &DownloadMangaCommand{
		CobraHelper: cobraHelper,
		MangadexApi: mangadexApi,
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

	for _, chapter := range chapters {
		chapterDetails, errResponse, errChapterDetails := c.MangadexApi.GetChapter(
			mangaId,
			language,
			chapter,
		)
		if errResponse != nil {
			color.Red(fmt.Sprintf("error with status code %v in chapter %v", errResponse.Status, chapter))
			continue
		}
		if errChapterDetails != nil {
			color.Red(fmt.Sprintf("error fetching chapter %v", chapter))
			continue
		}

		for _, chapterData := range chapterDetails.Data {
			fmt.Printf(
				color.HiBlackString("\nchapter fetched with success: %v - %v"),
				color.YellowString(util.GetChapterNumber(chapter)),
				color.CyanString(util.GetChapterName(chapterData.Attributes.Title, chapter)),
			)
		}
	}

	return nil
}
