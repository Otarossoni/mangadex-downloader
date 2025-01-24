package util

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func GetPageName(pageIdentification string) string {
	if pageIdentification == "" {
		return uuid.NewString()
	}

	splitIdentification := strings.Split(pageIdentification, "-")

	return strings.Trim(splitIdentification[0], " ")
}

func GetPageExtension(pageIdentification string) string {
	if pageIdentification == "" {
		return ".png"
	}

	splitIdentification := strings.Split(pageIdentification, ".")

	if len(splitIdentification) == 1 {
		return ""
	}

	return strings.Trim(
		fmt.Sprintf(".%v", splitIdentification[len(splitIdentification)-1]),
		" ",
	)
}
