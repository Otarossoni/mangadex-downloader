package helper_test

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/Otarossoni/mangadex-downloader/helper"
	"github.com/Otarossoni/mangadex-downloader/internal/entity"
	"github.com/stretchr/testify/assert"
)

type Page struct {
	Data      []byte
	Extension string
}

func Test_CreateZipFile(t *testing.T) {
	validateZipContents := func(zipPath string, expectedPages []entity.Page) error {
		file, err := os.Open(zipPath)
		if err != nil {
			return fmt.Errorf("error opening zip file: %w", err)
		}
		defer file.Close()

		stat, err := file.Stat()
		if err != nil {
			return fmt.Errorf("error reading zip file stats: %w", err)
		}

		zipReader, err := zip.NewReader(file, stat.Size())
		if err != nil {
			return fmt.Errorf("error reading zip file: %w", err)
		}

		if len(zipReader.File) != len(expectedPages) {
			return fmt.Errorf("unexpected number of files in zip: got %d, want %d", len(zipReader.File), len(expectedPages))
		}

		for i, f := range zipReader.File {
			expectedFileName := fmt.Sprintf("page_%d%s", i+1, expectedPages[i].Extension)
			if f.Name != expectedFileName {
				return fmt.Errorf("file name mismatch: got %s, want %s", f.Name, expectedFileName)
			}

			rc, err := f.Open()
			if err != nil {
				return fmt.Errorf("error opening file in zip: %w", err)
			}
			defer rc.Close()

			content, err := io.ReadAll(rc)
			if err != nil {
				return fmt.Errorf("error reading file content in zip: %w", err)
			}

			if !bytes.Equal(content, expectedPages[i].Data) {
				return fmt.Errorf("file content mismatch for %s", f.Name)
			}
		}

		return nil
	}

	t.Run("CreateValidZipFile", func(t *testing.T) {
		packer := helper.NewPacker()

		pages := []entity.Page{
			{Data: []byte("Content of page 1"), Extension: ".txt"},
			{Data: []byte("Content of page 2"), Extension: ".html"},
		}

		tempFile := t.TempDir() + "/test.zip"

		err := packer.CreateZipFile(pages, tempFile)
		if err != nil {
			t.Fatalf("CreateZipFile failed: %v", err)
		}

		err = validateZipContents(tempFile, pages)
		if err != nil {
			t.Errorf("Zip content validation failed: %v", err)
		}
	})

	t.Run("HandleEmptyPages", func(t *testing.T) {
		packer := helper.NewPacker()

		pages := []entity.Page{}

		tempFile := t.TempDir() + "/empty_test.zip"

		err := packer.CreateZipFile(pages, tempFile)
		if err != nil {
			t.Fatalf("CreateZipFile failed with empty pages: %v", err)
		}

		file, err := os.Open(tempFile)
		if err != nil {
			t.Fatalf("Failed to open zip file: %v", err)
		}
		defer file.Close()

		stat, err := file.Stat()
		if err != nil {
			t.Fatalf("Failed to get zip file stats: %v", err)
		}

		if stat.Size() == 0 {
			t.Errorf("Expected zip file not to be empty")
		}
	})
}

func Test_GetExamplesDescription(t *testing.T) {
	description := helper.GetExamplesDescription()

	assert.NotEqual(t, "", description)
}
