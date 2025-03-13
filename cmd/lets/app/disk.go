package app

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/wassef911/lets/pkg"
)

var diskService = pkg.NewDisk()

var ShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Display Disk Information",
}

var DiskUsageCmd = &cobra.Command{
	Use:   "disk space",
	Short: "Display disk usage for all mounts",
	Run: func(cmd *cobra.Command, args []string) {
		diskService.ShowDiskSpace()
	},
}

var FolderUsageCmd = &cobra.Command{
	Use:   "folder size for [directory]",
	Short: "Display directory disk usage",
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		directory := args[len(args)-1]
		diskService.ShowFolderSize(directory)
	},
}

var LimitedFolderUsageCmd = &cobra.Command{
	Use:   "files over [size] in [directory]",
	Short: "Display directory disk usage with a min size limit",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		size, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			panic(err)
		}
		directory := args[len(args)-1]
		diskService.ShowFolderSizeWithLimit(directory, size)
	},
}

func init() {
	ShowCmd.AddCommand(DiskUsageCmd, FolderUsageCmd, LimitedFolderUsageCmd)
	rootCmd.AddCommand(ShowCmd)
}
