package app

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/wassef911/lets/pkg"
)

var caseSensitive = true
var searchService = pkg.NewSearch(caseSensitive)

var searchFilesCmd = &cobra.Command{
	Use:   "search files for <pattern> in [directory]",
	Short: "Search files containing pattern",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pattern := args[0]
		directory := "."
		if len(args) > 1 {
			directory = args[1]
		}
		searchService.SearchFiles(pattern, directory)
	},
}

var countMatchesCmd = &cobra.Command{
	Use:   "count matches <pattern> in <file>",
	Short: "Count pattern occurrences in file",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		searchService.CountMatches(args[0], args[1])
	},
}

var findFilesCmd = &cobra.Command{
	Use:   "find files named <glob> in [directory] older than <days> days",
	Short: "Find files by name and age",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		glob := args[0]
		directory := "."
		days := 7

		if len(args) > 1 {
			directory = args[1]
			if len(args) > 2 {
				fmt.Sscanf(args[2], "%d", &days)
			}
		}
		searchService.FindFiles(glob, directory, days)
	},
}

func init() {
	searchFilesCmd.Flags().BoolVarP(&caseSensitive, "case-sensitive", "c", caseSensitive, "Case-sensitive search")
	rootCmd.AddCommand(searchFilesCmd, countMatchesCmd, findFilesCmd)
}
