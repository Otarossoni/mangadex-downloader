package util_test

import (
	"fmt"
	"testing"

	"github.com/Otarossoni/mangadex-downloader/internal/util"
)

func Test_GetChapterNumber(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{input: 1, expected: "00001"},
		{input: 12, expected: "00012"},
		{input: 123, expected: "00123"},
		{input: 1234, expected: "01234"},
		{input: 12345, expected: "12345"},
		{input: 0, expected: "00000"},
		{input: -1, expected: "-0001"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Input%d", tt.input), func(t *testing.T) {
			result := util.GetChapterNumber(tt.input)
			if result != tt.expected {
				t.Errorf("GetChapterNumber(%d) = %s; want %s", tt.input, result, tt.expected)
			}
		})
	}
}

func Test_GetChapterName(t *testing.T) {
	tests := []struct {
		title          string
		chapterNumber  int
		expectedOutput string
	}{
		{title: "", chapterNumber: 1, expectedOutput: "Chapter_00001"},
		{title: " Chapter Title ", chapterNumber: 12, expectedOutput: "Chapter Title"},
		{title: "Chapter 1", chapterNumber: 123, expectedOutput: "Chapter 1"},
		{title: "", chapterNumber: 0, expectedOutput: "Chapter_00000"},
		{title: "", chapterNumber: -1, expectedOutput: "Chapter_-0001"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Title:%q_Chapter:%d", tt.title, tt.chapterNumber), func(t *testing.T) {
			result := util.GetChapterName(tt.title, tt.chapterNumber)
			if result != tt.expectedOutput {
				t.Errorf("GetChapterName(%q, %d) = %q; want %q", tt.title, tt.chapterNumber, result, tt.expectedOutput)
			}
		})
	}
}
