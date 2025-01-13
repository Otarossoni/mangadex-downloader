package entity

import "github.com/spf13/cobra"

type CobraHelper interface {
	HandleMangaId(cmd *cobra.Command) (string, error)
	HandleChapters(cmd *cobra.Command) ([]int, error)
	HandleLanguage(cmd *cobra.Command) (string, error)
}
