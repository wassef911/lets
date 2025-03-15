package app

import (
	"github.com/spf13/cobra"

	"github.com/wassef911/lets/pkg"
)

func NewRootCmd(
	logger pkg.LoggerInterface,
	diskService pkg.DiskServiceInterface,
	inputOutputService pkg.InputOutputServiceInterface,
	procService pkg.ProcServiceInterface,
	searchService pkg.SearchServiceInterface,
) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "lets",
		Short: "Human-friendly system administration toolkit",
		Long:  `Human-friendly system administration toolkit.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	// create commands
	ShowCmd := NewShowCmd(logger, diskService)
	getCmd := NewGetCmd(logger, inputOutputService)
	ReplaceCmd := NewReplaceCmd(logger, inputOutputService)
	InspectCmd := NewProcCmd(logger, procService)
	TerminateCmd := NewTerminatedCmd(logger, procService)
	searchFilesCmd := NewSearchCmd(logger, searchService)
	countMatchesCmd := NewCountMatchesCmd(logger, searchService)
	findFilesCmd := NewFindFilesCmd(logger, searchService)

	// hook to root
	rootCmd.AddCommand(ShowCmd)
	rootCmd.AddCommand(getCmd, ReplaceCmd)
	rootCmd.AddCommand(InspectCmd, TerminateCmd)
	rootCmd.AddCommand(searchFilesCmd, countMatchesCmd, findFilesCmd)
	return rootCmd
}
