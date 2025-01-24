package service_test

import (
	"testing"

	"github.com/Otarossoni/mangadex-downloader/mock"
	"github.com/Otarossoni/mangadex-downloader/service"
	"github.com/stretchr/testify/assert"
)

func Test_GetChapter(t *testing.T) {
	t.Run("it should be able to get a single chapter", func(t *testing.T) {
		mangadexApi := service.NewMangadexApi()

		chapterDetails, errChapterDetails := mangadexApi.GetChapter(
			mock.ValidMangaUUID,
			mock.ValidMangaLanguage,
			1,
		)

		assert.NotNil(t, chapterDetails)
		assert.Equal(t, "ok", chapterDetails.Result)
		assert.NotEmpty(t, chapterDetails.Data)
		assert.Nil(t, errChapterDetails)
	})

	t.Run("it should not be able to get an non-existent manga", func(t *testing.T) {
		mangadexApi := service.NewMangadexApi()

		chapterDetails, errChapterDetails := mangadexApi.GetChapter(
			mock.InvalidMangaUUID,
			mock.ValidMangaLanguage,
			1,
		)

		assert.Nil(t, chapterDetails)
		assert.NotNil(t, errChapterDetails)
	})

	t.Run("it should not be able to get a manga with invalid language", func(t *testing.T) {
		mangadexApi := service.NewMangadexApi()

		chapterDetails, errChapterDetails := mangadexApi.GetChapter(
			mock.ValidMangaUUID,
			mock.InvalidMangaLanguage,
			1,
		)

		assert.Nil(t, chapterDetails)
		assert.NotNil(t, errChapterDetails)
	})
}

func Test_GetChapterPages(t *testing.T) {
	t.Run("it should be able to get pages from a chapter", func(t *testing.T) {
		mangadexApi := service.NewMangadexApi()

		chapterDetails, _ := mangadexApi.GetChapter(
			mock.ValidMangaUUID,
			mock.ValidMangaLanguage,
			1,
		)

		chapterPagesList, errChapterPagesList := mangadexApi.GetChapterPages(chapterDetails.Data[0].Id)

		assert.NotNil(t, chapterPagesList)
		assert.Equal(t, "ok", chapterPagesList.Result)
		assert.NotEqual(t, "", chapterPagesList.Result)
		assert.NotEmpty(t, chapterPagesList.Chapter.Data)
		assert.Nil(t, errChapterPagesList)
	})

	t.Run("it should not be able to get pages from non-existent chapter", func(t *testing.T) {
		mangadexApi := service.NewMangadexApi()

		chapterPagesList, errChapterPagesList := mangadexApi.GetChapterPages(mock.NonExistentChapterId)

		assert.Nil(t, chapterPagesList)
		assert.NotNil(t, errChapterPagesList)
	})
}

func Test_GetPage(t *testing.T) {
	t.Run("it should be able to get a single manga page", func(t *testing.T) {
		mangadexApi := service.NewMangadexApi()

		chapterDetails, _ := mangadexApi.GetChapter(
			mock.ValidMangaUUID,
			mock.ValidMangaLanguage,
			1,
		)

		chapterPagesList, _ := mangadexApi.GetChapterPages(chapterDetails.Data[0].Id)

		chapterPage, errChapterPage := mangadexApi.GetPage(
			chapterPagesList.BaseUrl,
			chapterPagesList.Chapter.Hash,
			chapterPagesList.Chapter.Data[0],
		)

		assert.NotNil(t, chapterPage)
		assert.NotEmpty(t, chapterPage)
		assert.Nil(t, errChapterPage)
	})

	t.Run("it should not be able to get a manga page with wrong identification", func(t *testing.T) {
		mangadexApi := service.NewMangadexApi()

		chapterDetails, _ := mangadexApi.GetChapter(
			mock.ValidMangaUUID,
			mock.ValidMangaLanguage,
			1,
		)

		chapterPagesList, _ := mangadexApi.GetChapterPages(chapterDetails.Data[0].Id)

		chapterPage, errChapterPage := mangadexApi.GetPage(
			chapterPagesList.BaseUrl,
			chapterPagesList.Chapter.Hash,
			mock.NonExistentPageIdentification,
		)

		assert.Nil(t, chapterPage)
		assert.Empty(t, chapterPage)
		assert.NotNil(t, errChapterPage)
	})
}
