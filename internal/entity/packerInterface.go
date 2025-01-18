package entity

type Packer interface {
	CreateZipFile(pages []Page, outputPath string) error
}
