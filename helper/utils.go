package helper

import (
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func hasDashInString(value string) bool {
	return strings.Contains(value, "-")
}

func convertStringToInt(input string) (int, error) {
	number, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return number, nil
}

func isValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func isValidURL(u string) bool {
	parsedURL, err := url.Parse(u)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return false
	}
	return true
}

func extractUUIDFromURL(url string) string {
	parts := strings.Split(url, "/")
	for _, part := range parts {
		if isValidUUID(part) {
			return part
		}
	}
	return ""
}

func ExistPath(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func AddBackslash(path string) string {
	if strings.HasSuffix(path, string(filepath.Separator)) {
		return path
	}

	return path + string(filepath.Separator)
}
