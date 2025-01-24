package entity

import (
	"time"
)

type MangadexApi interface {
	GetChapter(mangaId, language string, chapterNumber int) (*GetChapterResponse, error)
	GetChapterPages(chapterId string) (*GetChapterPagesResponse, error)
	GetPage(baseUrl, chapterHash, pageIdentification string) ([]byte, error)
}

type GetChapterResponse struct {
	Result   string                   `json:"result"`   // Request result: "ok" | "error"
	Response *string                  `json:"response"` // Type of response
	Data     []GetChapterDataResponse `json:"data"`     // Response content
}

type GetChapterPagesResponse struct {
	Result  string                          `json:"result"`  // Request result: "ok" | "error"
	BaseUrl string                          `json:"baseUrl"` // Base URL to fetch chapter individual pages
	Chapter *GetChapterPagesChapterResponse `json:"chapter"` // Chapter details
}

type GetChapterDataResponse struct {
	Id            string                            `json:"id"`            // Chapter ID
	Type          string                            `json:"type"`          // Structure type
	Attributes    GetChapterAttributesResponse      `json:"attributes"`    // Chapter attributes
	Relationships []GetChapterRelationshipsResponse `json:"relationships"` // Chapter relationships
}

type GetChapterAttributesResponse struct {
	Volume             *string    `json:"volume"`             // Chapter volume
	Chapter            string     `json:"chapter"`            // Chapter number
	Title              string     `json:"title"`              // Chapter title
	TranslatedLanguage string     `json:"translatedLanguage"` // Chapter language
	ExternalUrl        string     `json:"externalUrl"`        // Chapter view
	PublishAt          *time.Time `json:"publishAt"`          // Chapter publication date
	ReadableAt         *time.Time `json:"readableAt"`         // Chapter reading date
	CreatedAt          *time.Time `json:"createdAt"`          // Chapter creation date
	UpdatedAt          *time.Time `json:"updatedAt"`          // Chapter update date
	Pages              int        `json:"pages"`              // Number of pages in the chapter
	Version            int        `json:"version"`            // Chapter version
}

type GetChapterRelationshipsResponse struct {
	Id   string `json:"id"`   // Relationship Id
	Type string `json:"type"` // Relationship type
}

type GetChapterPagesChapterResponse struct {
	Hash      string   `json:"hash"`      // Chapter hash identification
	Data      []string `json:"data"`      // Pages identifications
	DataSaver []string `json:"dataSaver"` // Pages identifications in server
}
