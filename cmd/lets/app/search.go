package app

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/wassef911/lets/pkg"
)

func newSearchCmd(searchService pkg.SearchServiceInterface) *cobra.Command {
	searchCmd := &cobra.Command{
		Use:   "search files for <pattern> in [directory]",
		Short: "Search files containing pattern",
		Args:  cobra.MinimumNArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			pattern := args[3]
			directory := "."

			if len(args) > 4 {
				directory = args[5]
			}

			if _, err := os.Stat(directory); os.IsNotExist(err) {
				fmt.Printf("Error: Directory '%s' does not exist\n", directory)
				return
			}

			err := searchService.SearchFiles(pattern, directory)
			if err != nil {
				panic(err)
			}
		},
	}
	return searchCmd
}

func newCountMatchesCmd(searchService pkg.SearchServiceInterface) *cobra.Command {
	countMatchesCmd := &cobra.Command{
		Use:   "count matches <pattern> in <file>",
		Short: "Count pattern occurrences in file",
		Args:  cobra.ExactArgs(4),
		Run: func(cmd *cobra.Command, args []string) {
			pattern := args[1]
			file := args[3]
			err := searchService.CountMatches(pattern, file)
			if err != nil {
				panic(err)
			}
		},
	}
	return countMatchesCmd
}

func newFindFilesCmd(searchService pkg.SearchServiceInterface) *cobra.Command {
	findFilesCmd := &cobra.Command{
		Use:   "find files named <glob> in <directory> older than <days> days",
		Short: "Find files by name and age",
		Args:  cobra.MinimumNArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			glob, directory, days, err := pkg.ParseFind(args)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(glob, directory, days)
			cmdErr := searchService.FindFiles(glob, directory, days)
			if cmdErr != nil {
				panic(err)
			}
		},
	}
	return findFilesCmd
}
