package helper

import (
	"errors"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

type CobraHelper struct{}

func NewCobraHelper() *CobraHelper {
	return &CobraHelper{}
}

// HandleMangaId is the function that identifies the manga ID, either directly by parameter, or by URL
// Input example from flag:
//   - "https://mangadex.org/title/6a1d1cb1-ecd5-40d9-89ff-9d88e40b136b/tokyo-ghoul"
//   - "6a1d1cb1-ecd5-40d9-89ff-9d88e40b136b"s
func (c *CobraHelper) HandleMangaId(cmd *cobra.Command) (string, error) {
	id, _ := cmd.Flags().GetString("mangaId")

	if id != "" && isValidUUID(id) {
		return id, nil
	}

	url, _ := cmd.Flags().GetString("url")

	if isValidURL(url) {
		if mangaId := extractUUIDFromURL(url); mangaId != "" {
			return mangaId, nil
		}
	}

	return "", errors.New("\ncould not identify a valid Manga ID from the provided flags")
}

// HandleChapters is the function that handles chapter numbers
// Input example from flag: "1-20,25,30,40-50"
func (c *CobraHelper) HandleChapters(cmd *cobra.Command) ([]int, error) {
	var chapters []int

	rawChapters, _ := cmd.Flags().GetString("chapters")

	splitChapterRanges := strings.Split(rawChapters, ",")

	for _, splitChapterRange := range splitChapterRanges {

		if hasDashInString(splitChapterRange) {
			rawRanges := strings.SplitN(splitChapterRange, "-", 2)

			if len(rawRanges) > 2 {
				return chapters, errors.New("\nrange provided is invalid")
			}

			minRange, errMinRange := convertStringToInt(strings.TrimSpace(rawRanges[0]))
			if errMinRange != nil {
				return chapters, errors.New("\nrange provided is invalid")
			}

			maxRange, errMaxRange := convertStringToInt(strings.TrimSpace(rawRanges[1]))
			if errMaxRange != nil {
				return chapters, errors.New("\nrange provided is invalid")
			}

			if minRange > maxRange {
				return nil, errors.New("\nrange start cannot be greater than range end")
			}

			rangeSize := maxRange - minRange + 1
			for i := 0; i < rangeSize; i++ {
				chapters = append(chapters, minRange+i)
			}

			continue
		}

		if newChapter, err := convertStringToInt(splitChapterRange); err == nil {
			chapters = append(chapters, newChapter)
		}
	}

	return chapters, nil
}

// HandleLanguage is the function that handles the chapters languages
// Input example from flag: "en", "pt-br"
func (c *CobraHelper) HandleLanguage(cmd *cobra.Command) (string, error) {
	rawLanguage, _ := cmd.Flags().GetString("language")

	if rawLanguage == "" {
		return "en", nil
	}

	pattern := `^[a-z]{2}(-[a-z]{2})?$`

	regex := regexp.MustCompile(pattern)
	if regex.MatchString(rawLanguage) {
		return rawLanguage, nil
	}

	return "", errors.New("\nformat of the language provided is invalid")
}

// HandlePackExtension is the function that handles the pack extension, ensuring a default value
// Input example from flag: ".zip", ".cbz"
func (c *CobraHelper) HandlePackExtension(cmd *cobra.Command) (string, error) {
	rawExtension, _ := cmd.Flags().GetString("extension")

	switch rawExtension {
	case "":
		return ".zip", nil
	case ".cbz", ".zip":
		return rawExtension, nil
	default:
		return "", errors.New("\npack extension provided is invalid")
	}
}
