package app

import (
	"github.com/spf13/cobra"

	"github.com/wassef911/lets/pkg"
)

func newProcCmd(procService pkg.ProcServiceInterface) *cobra.Command {
	ProcessesCmd := &cobra.Command{
		Use:   "processes",
		Short: "List all running processes",
		Run: func(cmd *cobra.Command, args []string) {
			procService.Processes()
		},
	}

	ResourcesCmd := &cobra.Command{
		Use:   "resources",
		Short: "Live system resource view",
		Run: func(cmd *cobra.Command, args []string) {
			procService.Resources()
		},
	}

	InspectCmd := &cobra.Command{
		Use:   "inspect",
		Short: "Inspect Processes and Resources",
	}
	InspectCmd.AddCommand(ProcessesCmd, ResourcesCmd)
	return InspectCmd
}

func newTerminatedCmd(procService pkg.ProcServiceInterface) *cobra.Command {
	TerminateCmd := &cobra.Command{
		Use:   "kill process [name]",
		Short: "Terminate by process name",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[len(args)-1]
			procService.KillProcessByName(name)
		},
	}
	return TerminateCmd
}
