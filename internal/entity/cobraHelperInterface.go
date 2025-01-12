package entity

import "github.com/spf13/cobra"

type CobraHelper interface {
	HandleMangaId(cmd *cobra.Command) (string, error)
	HandleChapters(cmd *cobra.Command) ([]int, error)
	HandleLanguages(cmd *cobra.Command) []string
}
