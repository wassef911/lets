package app

import (
	"github.com/spf13/cobra"

	"github.com/wassef911/lets/pkg"
)

var procService = pkg.NewProc()

var ProcessesCmd = &cobra.Command{
	Use:   "show processes",
	Short: "List all running processes",
	Run: func(cmd *cobra.Command, args []string) {
		procService.Processes()
	},
}

var ResourcesCmd = &cobra.Command{
	Use:   "show resources",
	Short: "Live system resource view",
	Run: func(cmd *cobra.Command, args []string) {
		procService.Resources()
	},
}

var TerminateCmd = &cobra.Command{
	Use:   "kill process [name]",
	Short: "Terminate by process name",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[1]
		procService.KillProcessByName(name)
	},
}

func init() {
	rootCmd.AddCommand(ProcessesCmd, ResourcesCmd, TerminateCmd)
}
