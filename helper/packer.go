package helper

import (
	"archive/zip"
	"fmt"
	"os"

	"github.com/Otarossoni/mangadex-downloader/internal/entity"
)

type Packer struct{}

func NewPacker() *Packer {
	return &Packer{}
}

func (c *Packer) CreateZipFile(pages []entity.Page, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("\nerror creating file: %w", err)
	}
	defer file.Close()

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	for i, page := range pages {
		fileName := fmt.Sprintf("page_%d%v", i+1, page.Extension)
		writer, err := zipWriter.Create(fileName)
		if err != nil {
			return fmt.Errorf("\nerror creating file in zip: %w", err)
		}

		_, err = writer.Write(page.Data)
		if err != nil {
			return fmt.Errorf("\nerror creating file in zip: %w", err)
		}
	}

	return nil
}
