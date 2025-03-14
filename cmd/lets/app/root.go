package app

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/wassef911/lets/pkg"
)

var rootCmd = &cobra.Command{
	Use:   "lets",
	Short: "Human-friendly system administration toolkit",
	Long:  `Human-friendly system administration toolkit.`,
}

func Execute() {
	// create services
	diskService := pkg.NewDisk()
	inputOutputService := pkg.NewInputOutput()
	procService := pkg.NewProc()
	searchService := pkg.NewSearch(true)

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

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
