package util_test

import (
	"regexp"
	"testing"

	"github.com/Otarossoni/mangadex-downloader/internal/util"
)

func Test_GetPageName(t *testing.T) {
	const uuidRegexp = `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`

	tests := []struct {
		name               string
		pageIdentification string
		wantPageName       string
	}{
		{
			name:               "Empty pageIdentification",
			pageIdentification: "",
			wantPageName:       uuidRegexp,
		},
		{
			name:               "Valid pageIdentification with hyphen",
			pageIdentification: "home-page-123",
			wantPageName:       "home",
		},
		{
			name:               "Valid pageIdentification with leading/trailing spaces",
			pageIdentification: "  about  - v1",
			wantPageName:       "about",
		},
		{
			name:               "Invalid pageIdentification (no hyphen)",
			pageIdentification: "homepage",
			wantPageName:       "homepage",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPageName := util.GetPageName(tt.pageIdentification)

			if tt.wantPageName == uuidRegexp {
				if matched, _ := regexp.MatchString(uuidRegexp, gotPageName); !matched {
					t.Errorf("GetPageName() = %v, want a UUID-like string", gotPageName)
				}
			} else {
				if gotPageName != tt.wantPageName {
					t.Errorf("GetPageName() = %v, want %v", gotPageName, tt.wantPageName)
				}
			}
		})
	}
}

func Test_GetPageExtension(t *testing.T) {
	tests := []struct {
		name               string
		pageIdentification string
		wantExtension      string
	}{
		{
			name:               "Empty pageIdentification",
			pageIdentification: "",
			wantExtension:      ".png",
		},
		{
			name:               "Valid pageIdentification with extension",
			pageIdentification: "home-page.png",
			wantExtension:      ".png",
		},
		{
			name:               "Invalid pageIdentification (no extension)",
			pageIdentification: "contact",
			wantExtension:      "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotExtension := util.GetPageExtension(tt.pageIdentification)

			if gotExtension != tt.wantExtension {
				t.Errorf("GetPageExtension() = %v, want %v", gotExtension, tt.wantExtension)
			}
		})
	}
}
