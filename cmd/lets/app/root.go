package app

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lets",
	Short: "Human-friendly system administration toolkit",
	Long:  `Human-friendly system administration toolkit.`,
}

func Execute() {

	rootCmd.AddCommand(getCmd, ReplaceCmd)
	ShowCmd.AddCommand(DiskUsageCmd, FolderUsageCmd, LimitedFolderUsageCmd)
	rootCmd.AddCommand(ShowCmd)
	rootCmd.AddCommand(searchFilesCmd, countMatchesCmd, findFilesCmd)
	InspectCmd.AddCommand(ProcessesCmd, ResourcesCmd)
	rootCmd.AddCommand(InspectCmd, TerminateCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
