package helper_test

import (
	"testing"

	"github.com/Otarossoni/mangadex-downloader/helper"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_HandleMangaId(t *testing.T) {
	helper := helper.NewCobraHelper()

	tests := []struct {
		name          string
		mangaIdFlag   string
		urlFlag       string
		expectedId    string
		expectedError bool
	}{
		{
			name:          "Valid Manga ID",
			mangaIdFlag:   "6a1d1cb1-ecd5-40d9-89ff-9d88e40b136b",
			urlFlag:       "",
			expectedId:    "6a1d1cb1-ecd5-40d9-89ff-9d88e40b136b",
			expectedError: false,
		},
		{
			name:          "Valid URL with Manga ID",
			mangaIdFlag:   "",
			urlFlag:       "https://mangadex.org/title/6a1d1cb1-ecd5-40d9-89ff-9d88e40b136b/tokyo-ghoul",
			expectedId:    "6a1d1cb1-ecd5-40d9-89ff-9d88e40b136b",
			expectedError: false,
		},
		{
			name:          "Invalid Manga ID and URL",
			mangaIdFlag:   "invalid-id",
			urlFlag:       "https://invalid-url.com",
			expectedId:    "",
			expectedError: true,
		},
		{
			name:          "Empty Flags",
			mangaIdFlag:   "",
			urlFlag:       "",
			expectedId:    "",
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{}
			cmd.Flags().String("mangaId", tt.mangaIdFlag, "Manga ID")
			cmd.Flags().String("url", tt.urlFlag, "URL")

			id, err := helper.HandleMangaId(cmd)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expectedId, id)
		})
	}
}

func Test_HandleChapters(t *testing.T) {
	helper := helper.NewCobraHelper()

	tests := []struct {
		name             string
		chaptersFlag     string
		expectedChapters []int
		expectedError    bool
	}{
		{
			name:             "Valid Single Chapter",
			chaptersFlag:     "5",
			expectedChapters: []int{5},
			expectedError:    false,
		},
		{
			name:             "Valid Multiple Chapters",
			chaptersFlag:     "1,3,5",
			expectedChapters: []int{1, 3, 5},
			expectedError:    false,
		},
		{
			name:             "Valid Range of Chapters",
			chaptersFlag:     "2-5",
			expectedChapters: []int{2, 3, 4, 5},
			expectedError:    false,
		},
		{
			name:             "Mixed Single Chapters and Ranges",
			chaptersFlag:     "1,3-5,7",
			expectedChapters: []int{1, 3, 4, 5, 7},
			expectedError:    false,
		},
		{
			name:             "Invalid Range Format",
			chaptersFlag:     "5-3",
			expectedChapters: nil,
			expectedError:    true,
		},
		{
			name:             "Invalid Characters in Input",
			chaptersFlag:     "1,a,3-5",
			expectedChapters: []int{1, 3, 4, 5},
			expectedError:    false,
		},
		{
			name:             "Empty Input",
			chaptersFlag:     "",
			expectedChapters: nil,
			expectedError:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{}
			cmd.Flags().String("chapters", tt.chaptersFlag, "Chapters flag")

			chapters, err := helper.HandleChapters(cmd)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expectedChapters, chapters)
		})
	}
}

func Test_HandleLanguage(t *testing.T) {
	helper := helper.NewCobraHelper()

	tests := []struct {
		name           string
		languageFlag   string
		expectedOutput string
		expectedError  bool
	}{
		{
			name:           "Default Language (Empty Flag)",
			languageFlag:   "",
			expectedOutput: "en",
			expectedError:  false,
		},
		{
			name:           "Valid Language Code",
			languageFlag:   "en",
			expectedOutput: "en",
			expectedError:  false,
		},
		{
			name:           "Valid Language and Region Code",
			languageFlag:   "en-us",
			expectedOutput: "en-us",
			expectedError:  false,
		},
		{
			name:           "Invalid Language Code (Too Long)",
			languageFlag:   "english",
			expectedOutput: "",
			expectedError:  true,
		},
		{
			name:           "Invalid Language Code (Special Characters)",
			languageFlag:   "en@us",
			expectedOutput: "",
			expectedError:  true,
		},
		{
			name:           "Invalid Language Code (Capital Letters)",
			languageFlag:   "EN",
			expectedOutput: "",
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{}
			cmd.Flags().String("language", tt.languageFlag, "Language flag")

			output, err := helper.HandleLanguage(cmd)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expectedOutput, output)
		})
	}
}

func Test_HandlePackExtension(t *testing.T) {
	helper := helper.NewCobraHelper()

	tests := []struct {
		name           string
		extensionFlag  string
		expectedOutput string
		expectedError  bool
	}{
		{
			name:           "Default Extension (Empty Flag)",
			extensionFlag:  "",
			expectedOutput: ".zip",
			expectedError:  false,
		},
		{
			name:           "Valid Extension .zip",
			extensionFlag:  ".zip",
			expectedOutput: ".zip",
			expectedError:  false,
		},
		{
			name:           "Valid Extension .cbz",
			extensionFlag:  ".cbz",
			expectedOutput: ".cbz",
			expectedError:  false,
		},
		{
			name:           "Invalid Extension .rar",
			extensionFlag:  ".rar",
			expectedOutput: "",
			expectedError:  true,
		},
		{
			name:           "Invalid Extension Without Dot",
			extensionFlag:  "cbz",
			expectedOutput: "",
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{}
			cmd.Flags().String("extension", tt.extensionFlag, "Extension flag")

			output, err := helper.HandlePackExtension(cmd)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expectedOutput, output)
		})
	}
}

func Test_HandleOutPath(t *testing.T) {
	cobraHelper := helper.NewCobraHelper()

	cmd := &cobra.Command{}
	cmd.Flags().String("outPath", "", "output path")

	t.Run("empty outPath", func(t *testing.T) {
		outPath, err := cobraHelper.HandleOutPath(cmd)
		assert.NoError(t, err)
		assert.Equal(t, "", outPath)
	})

	t.Run("existing outPath", func(t *testing.T) {
		tempDir := t.TempDir()

		cmd.Flags().Set("outPath", tempDir)

		outPath, err := cobraHelper.HandleOutPath(cmd)
		assert.NoError(t, err)
		assert.Equal(t, tempDir+"\\", outPath)
	})

	t.Run("non-existing outPath", func(t *testing.T) {
		cmd.Flags().Set("outPath", "/invalid/path")

		outPath, err := cobraHelper.HandleOutPath(cmd)
		assert.Error(t, err)
		assert.Equal(t, "", outPath)
		assert.Equal(t, "\nout path provided not exist", err.Error())
	})
}
