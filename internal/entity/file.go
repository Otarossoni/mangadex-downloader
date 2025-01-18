package entity

type Page struct {
	Name      string
	Extension string
	Data      []byte
}

type FilePack struct {
	Name      string
	Extension string
	Data      []Page
}
