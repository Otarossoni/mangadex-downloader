package service

import (
	"encoding/json"

	request "github.com/Otarossoni/mangadex-downloader/http"
	"github.com/Otarossoni/mangadex-downloader/internal/entity"
)

const BASE_URL = "https://api.mangadex.org"

type MangadexApi struct{}

func NewMangadexApi() *MangadexApi {
	return &MangadexApi{}
}

func (m *MangadexApi) GetChapter(mangaId, language string, chapterNumber int) (*entity.GetChapterResponse, error) {
	params := request.Params{
		Method: "GET",
		URL:    BASE_URL + "/chapter",
		QueryParams: map[string]interface{}{
			"manga":                 mangaId,
			"translatedLanguage[0]": language,
			"chapter[0]":            chapterNumber,
			"contentRating[0]":      "safe",
			"contentRating[1]":      "suggestive",
			"contentRating[2]":      "erotica",
			"contentRating[3]":      "pornographic",
		},
	}

	response, err := request.New(params)
	if err != nil {
		return nil, err
	}

	var getChapterResponse entity.GetChapterResponse
	err = json.Unmarshal(response.RawBody, &getChapterResponse)
	return &getChapterResponse, err
}

func (m *MangadexApi) GetChapterPages(chapterId string) (*entity.GetChapterPagesResponse, error) {
	params := request.Params{
		Method: "GET",
		URL:    BASE_URL + "/at-home/server/" + chapterId,
	}

	response, err := request.New(params)
	if err != nil {
		return nil, err
	}

	var getChapterPagesResponse entity.GetChapterPagesResponse
	err = json.Unmarshal(response.RawBody, &getChapterPagesResponse)
	return &getChapterPagesResponse, err
}

func (m *MangadexApi) GetPage(baseUrl, chapterHash, pageIdentification string) ([]byte, error) {
	params := request.Params{
		Method: "GET",
		URL:    baseUrl + "/data/" + chapterHash + "/" + pageIdentification,
	}

	response, err := request.New(params)
	_, ok := err.(*json.SyntaxError)

	if err != nil && !ok {
		return nil, err
	}

	return response.RawBody, nil
}
