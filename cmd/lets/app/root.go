package app

import (
	"github.com/spf13/cobra"

	"github.com/wassef911/lets/pkg"
)

func NewRootCmd(
	diskService pkg.DiskServiceInterface,
	inputOutputService pkg.InputOutputServiceInterface,
	procService pkg.ProcServiceInterface,
	searchService pkg.SearchServiceInterface,
) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "lets",
		Short: "Human-friendly system administration toolkit",
		Long:  `Human-friendly system administration toolkit.`,
	}

	// create commands
	ShowCmd := newShowCmd(diskService)
	getCmd := newGetCmd(inputOutputService)
	ReplaceCmd := newReplaceCmd(inputOutputService)
	InspectCmd := newProcCmd(procService)
	TerminateCmd := newTerminatedCmd(procService)
	searchFilesCmd := newSearchCmd(searchService)
	countMatchesCmd := newCountMatchesCmd(searchService)
	findFilesCmd := newFindFilesCmd(searchService)
	// hook to root
	rootCmd.AddCommand(ShowCmd)
	rootCmd.AddCommand(getCmd, ReplaceCmd)
	rootCmd.AddCommand(InspectCmd, TerminateCmd)
	rootCmd.AddCommand(searchFilesCmd, countMatchesCmd, findFilesCmd)
	return rootCmd
}
