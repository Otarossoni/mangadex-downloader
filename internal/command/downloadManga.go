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
	Packer      entity.Packer
}

func NewDownloadMangaCommand(cobraHelper entity.CobraHelper, mangadexApi entity.MangadexApi, packer entity.Packer) *DownloadMangaCommand {
	return &DownloadMangaCommand{
		CobraHelper: cobraHelper,
		MangadexApi: mangadexApi,
		Packer:      packer,
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

	packExtension, errPackExtension := c.CobraHelper.HandlePackExtension(cmd)
	if errPackExtension != nil {
		return errPackExtension
	}

	outPath, errOutPath := c.CobraHelper.HandleOutPath(cmd)
	if errOutPath != nil {
		return errOutPath
	}

	for _, chapter := range chapters {
		chapterDetails, errChapterDetails := c.MangadexApi.GetChapter(
			mangaId,
			language,
			chapter,
		)
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

			chapterPagesList, errChapterPagesList := c.MangadexApi.GetChapterPages(chapterData.Id)
			if errChapterPagesList != nil {
				color.Red(fmt.Sprintf("\nerror fetching chapter %v pages list", chapter))
				continue
			}

			fmt.Printf(
				color.HiBlackString("\n%v fetching chapter %v - %v pages..."),
				color.HiCyanString("+"),
				color.YellowString(util.GetChapterNumber(chapter)),
				color.CyanString(util.GetChapterName(chapterData.Attributes.Title, chapter)),
			)

			filePack := entity.FilePack{
				Name: fmt.Sprintf(
					"%v - %v",
					util.GetChapterNumber(chapter),
					util.GetChapterName(chapterData.Attributes.Title, chapter),
				),
				Extension: packExtension,
			}

			for _, chapterPageIdentification := range chapterPagesList.Chapter.Data {
				chapterPage, errChapterPage := c.MangadexApi.GetPage(
					chapterPagesList.BaseUrl,
					chapterPagesList.Chapter.Hash,
					chapterPageIdentification,
				)
				if errChapterPage != nil {
					color.Red(fmt.Sprintf("\nerror fetching chapter %v pages list", chapter))
					continue
				}

				filePack.Data = append(
					filePack.Data,
					entity.Page{
						Name:      util.GetPageName(chapterPageIdentification),
						Extension: util.GetPageExtension(chapterPageIdentification),
						Data:      chapterPage,
					},
				)
			}

			zipName := fmt.Sprintf("%v%v", filePack.Name, filePack.Extension)

			c.Packer.CreateZipFile(
				filePack.Data,
				outPath+zipName,
			)

			fmt.Printf(
				color.HiBlackString("\n%v file %v saved\n"),
				color.HiGreenString("+"),
				color.CyanString(zipName),
			)
		}
	}

	return nil
}
