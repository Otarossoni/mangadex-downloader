package util

import (
	"fmt"
	"strings"
)

func GetChapterNumber(chapterNumber int) string {
	return fmt.Sprintf("%05d", chapterNumber)
}

func GetChapterName(title string, chapterNumber int) string {
	if title == "" {
		return "Chapter_" + GetChapterNumber(chapterNumber)
	}
	return strings.Trim(title, " ")
}
