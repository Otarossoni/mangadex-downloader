package helper_test

import (
	"path/filepath"
	"testing"

	"github.com/Otarossoni/mangadex-downloader/helper"
	"github.com/Otarossoni/mangadex-downloader/mock"
	"github.com/stretchr/testify/assert"
)

func Test_HasDashInString(t *testing.T) {
	t.Run("it should be able to verify if has dash in string", func(t *testing.T) {
		stringTest := "has-dash-in-string"

		result := helper.HasDashInString(stringTest)

		assert.Equal(t, result, true)
	})

	t.Run("it should be able to verify if has no dash in string", func(t *testing.T) {
		stringTest := "has no dash in string"

		result := helper.HasDashInString(stringTest)

		assert.Equal(t, result, false)
	})
}

func Test_ConvertStringToInt(t *testing.T) {
	t.Run("it should be able to convert string to int", func(t *testing.T) {
		numberStringTest := "1"

		result, err := helper.ConvertStringToInt(numberStringTest)

		assert.Equal(t, result, 1)
		assert.Nil(t, err)
	})

	t.Run("it should not be able to convert a invalid string to int", func(t *testing.T) {
		invalidNumberStringTest := "A"

		result, err := helper.ConvertStringToInt(invalidNumberStringTest)

		assert.Equal(t, result, 0)
		assert.NotNil(t, err)
	})
}

func Test_IsValidUUID(t *testing.T) {
	t.Run("it should be able to verify if is a valid UUID", func(t *testing.T) {
		result := helper.IsValidUUID(mock.ValidMangaUUID)

		assert.Equal(t, result, true)
	})

	t.Run("it should be able to verify if is a invalid UUID", func(t *testing.T) {
		result := helper.IsValidUUID(mock.InvalidMangaLanguage)

		assert.Equal(t, result, false)
	})
}

func Test_IsValidURL(t *testing.T) {
	t.Run("it should be able to verify if URL is valid", func(t *testing.T) {
		pathTest := "https://mangadex.org/"

		result := helper.IsValidURL(pathTest)

		assert.Equal(t, result, true)
	})

	t.Run("it should be able to verify if URL is invalid", func(t *testing.T) {
		pathTest := "is a invalid URL"

		result := helper.IsValidURL(pathTest)

		assert.Equal(t, result, false)
	})
}

func Test_ExtractUUIDFromURL(t *testing.T) {
	t.Run("it should be able to extract UUID from URL", func(t *testing.T) {
		urlTest := "https://mangadex.org/title/a1f9462a-a340-4098-81de-1f2536b2f782/steins-gate"

		result := helper.ExtractUUIDFromURL(urlTest)

		assert.Equal(t, result, "a1f9462a-a340-4098-81de-1f2536b2f782")
	})

	t.Run("should not be able to extract UUID from a non-Mangadex URL", func(t *testing.T) {
		urlTest := "https://www.google.com/"

		result := helper.ExtractUUIDFromURL(urlTest)

		assert.Equal(t, result, "")
	})
}

func Test_ExistPath(t *testing.T) {
	tempDir := t.TempDir()
	t.Run("it should be able to verify if path exist", func(t *testing.T) {
		if !helper.ExistPath(tempDir) {
			t.Errorf("Expected path to exist: %s", tempDir)
		}
	})

	t.Run("it should not be able to verify a non-existent path", func(t *testing.T) {
		nonExistentPath := tempDir + "/nonexistent"
		if helper.ExistPath(nonExistentPath) {
			t.Errorf("Expected path not to exist: %s", nonExistentPath)
		}
	})
}

func Test_AddBackslash(t *testing.T) {
	t.Run("should be able to add backslash if the string does not have it", func(t *testing.T) {
		stringTest := "test"

		result := helper.AddBackslash(stringTest)

		assert.Equal(t, result, "test"+string(filepath.Separator))
	})

	t.Run("should be able to add backslash if the string has", func(t *testing.T) {
		stringTest := "test" + string(filepath.Separator)

		result := helper.AddBackslash(stringTest)

		assert.Equal(t, result, stringTest)
	})
}
