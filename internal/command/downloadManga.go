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
			color.Red(fmt.Sprintf("\nerror with status code %v getting chapter %v details", errResponse.Status, chapter))
			continue
		}
		if errChapterDetails != nil {
			color.Red(fmt.Sprintf("\nerror fetching chapter %v details", chapter))
			continue
		}

		for _, chapterData := range chapterDetails.Data {
			fmt.Printf(
				color.HiBlackString("\nchapter fetched with success: chapter %v - %v"),
				color.YellowString(util.GetChapterNumber(chapter)),
				color.CyanString(util.GetChapterName(chapterData.Attributes.Title, chapter)),
			)

			chapterPagesList, errResponse, errChapterPagesList := c.MangadexApi.GetChapterPages(chapterData.Id)
			if errResponse != nil {
				color.Red(fmt.Sprintf("\nerror with status code %v fetching chapter %v pages list", errResponse.Status, chapter))
				continue
			}
			if errChapterPagesList != nil {
				color.Red(fmt.Sprintf("\nerror fetching chapter %v pages list", chapter))
				continue
			}

			fmt.Printf(
				color.HiBlackString("\n%v fetching chapter %v - %v pages...\n"),
				color.HiCyanString("+"),
				color.YellowString(util.GetChapterNumber(chapter)),
				color.CyanString(util.GetChapterName(chapterData.Attributes.Title, chapter)),
			)

			for _, chapterPageIdentification := range chapterPagesList.Chapter.Data {
				chapterPage, errResponse, errChapterPage := c.MangadexApi.GetPage(
					chapterPagesList.BaseUrl,
					chapterPagesList.Chapter.Hash,
					chapterPageIdentification,
				)
				if errResponse != nil {
					chapterPageNumber := util.GetChapterPageNumber(chapterPageIdentification)
					color.Red(fmt.Sprintf("\nerror with status code %v getting chapter %v page %v", errResponse.Status, chapter, chapterPageNumber))
					continue
				}
				if errChapterPage != nil {
					color.Red(fmt.Sprintf("\nerror fetching chapter %v pages list", chapter))
					continue
				}

				color.Magenta(fmt.Sprintf("%v", chapterPage[1000]))
			}
		}
	}

	return nil
}
