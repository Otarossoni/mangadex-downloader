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

func (m *MangadexApi) GetChapter(mangaId, language string, chapterNumber int) (*entity.GetChapterResponse, *entity.ErrorResponse, error) {
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
		return nil, nil, err
	}

	if response.StatusCode > 300 {
		resp, err := parseError(response.RawBody)
		return nil, resp, err
	}

	var getChapterResponse entity.GetChapterResponse
	err = json.Unmarshal(response.RawBody, &getChapterResponse)
	return &getChapterResponse, nil, err
}

func (m *MangadexApi) GetChapterPages(chapterId string) (*entity.GetChapterPagesResponse, *entity.ErrorResponse, error) {
	params := request.Params{
		Method: "GET",
		URL:    BASE_URL + "/at-home/server/" + chapterId,
	}

	response, err := request.New(params)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode > 300 {
		resp, err := parseError(response.RawBody)
		return nil, resp, err
	}

	var getChapterPagesResponse entity.GetChapterPagesResponse
	err = json.Unmarshal(response.RawBody, &getChapterPagesResponse)
	return &getChapterPagesResponse, nil, err
}

func (m *MangadexApi) GetPage(baseUrl, chapterHash, pageIdentification string) ([]byte, *entity.ErrorResponse, error) {
	params := request.Params{
		Method: "GET",
		URL:    baseUrl + "/data/" + chapterHash + "/" + pageIdentification,
	}

	response, err := request.New(params)
	_, ok := err.(*json.SyntaxError)

	if err != nil && !ok {
		return nil, nil, err
	}

	if response.StatusCode > 300 {
		resp, err := parseError(response.RawBody)
		return nil, resp, err
	}

	return response.RawBody, nil, nil
}

func parseError(body []byte) (*entity.ErrorResponse, error) {
	var errResponse entity.ErrorResponse
	if err := json.Unmarshal(body, &errResponse); err != nil {
		return nil, err
	}
	return &errResponse, nil
}
