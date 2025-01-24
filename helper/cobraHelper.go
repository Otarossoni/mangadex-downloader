package helper

import (
	"errors"
	"regexp"
	"strings"

	"github.com/fatih/color"
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

	if id != "" && IsValidUUID(id) {
		return id, nil
	}

	url, _ := cmd.Flags().GetString("url")

	if IsValidURL(url) {
		if mangaId := ExtractUUIDFromURL(url); mangaId != "" {
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

		if HasDashInString(splitChapterRange) {
			rawRanges := strings.SplitN(splitChapterRange, "-", 2)

			if len(rawRanges) > 2 {
				return chapters, errors.New("\nrange provided is invalid")
			}

			minRange, errMinRange := ConvertStringToInt(strings.TrimSpace(rawRanges[0]))
			if errMinRange != nil {
				return chapters, errors.New("\nrange provided is invalid")
			}

			maxRange, errMaxRange := ConvertStringToInt(strings.TrimSpace(rawRanges[1]))
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

		if newChapter, err := ConvertStringToInt(splitChapterRange); err == nil {
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

// HandleOutPath is the function that handles the output path, validating and formatting it
func (c *CobraHelper) HandleOutPath(cmd *cobra.Command) (string, error) {
	rawOutPath, _ := cmd.Flags().GetString("outPath")

	if rawOutPath == "" {
		return rawOutPath, nil
	}

	if ExistPath(rawOutPath) {
		return AddBackslash(rawOutPath), nil
	}

	return "", errors.New("\nout path provided not exist")
}

func GetExamplesDescription() string {
	var description string

	description += color.HiBlackString("\"I want to download from chapter 1 to chapter 100 of Chainsaw Man\"\n")
	description += color.HiMagentaString("mangadex-downloader --url https://mangadex.org/title/a77742b1-befd-49a4-bff5-1ad4e6b0ef7b/chainsaw-man --chapters 1-100\n")

	description += color.HiBlackString("\n\"I want to download from chapter 1 to chapter 25, from 50 to 75 and chapter 100, in just one command\"\n")
	description += color.HiMagentaString("mangadex-downloader --url https://mangadex.org/title/dd8a907a-3850-4f95-ba03-ba201a8399e3/fullmetal-alchemist --chapters 1-25;50-75;100\n")

	description += color.HiBlackString("\n\"I want to download from chapter 1 to chapter 100 of Tokyo Ghoul, but in Brazilian Portuguese\"\n")
	description += color.HiMagentaString("mangadex-downloader --url https://mangadex.org/title/6a1d1cb1-ecd5-40d9-89ff-9d88e40b136b/tokyo-ghoul --chapters 1-100 --language pt-br\n")

	description += color.HiBlackString("\n\"I want to download from chapter 1 to chapter 50 of Tokyo Ghoul, but in Brazilian Portuguese and packed in .cbz\"\n")
	description += color.HiMagentaString("mangadex-downloader --mangaId 6a1d1cb1-ecd5-40d9-89ff-9d88e40b136b --chapters 1-50 --language pt-br --extension .cbz\n")

	description += color.HiBlackString("\n\"I want to download from chapter 1 to chapter 10 of Sousou no Frieren, and generate final file in my desktop\"\n")
	description += color.HiMagentaString("mangadex-downloader --mangaId b0b721ff-c388-4486-aa0f-c2b0bb321512 --chapters 1-10 --outPath C:\\Users\\$user\\Desktop")

	return description
}
