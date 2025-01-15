package util

import "fmt"

func GetChapterNumber(chapterNumber int) string {
	return fmt.Sprintf("%05d", chapterNumber)
}

func GetChapterName(title string, chapterNumber int) string {
	formattedChapterNumber := GetChapterNumber(chapterNumber)

	if title == "" {
		return "Chapter_" + formattedChapterNumber
	}

	return title
}
