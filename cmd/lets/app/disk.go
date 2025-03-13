package app

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/wassef911/lets/pkg"
)

var diskService = pkg.NewDisk()

var DiskUsageCmd = &cobra.Command{
	Use:   "show disk space",
	Short: "Display disk usage for all mounts",
	Run: func(cmd *cobra.Command, args []string) {
		diskService.ShowDiskSpace()
	},
}

var FolderUsageCmd = &cobra.Command{
	Use:   "show folder size for [directory]",
	Short: "Display directory disk usage",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		directory := args[0]
		diskService.ShowFolderSize(directory)
	},
}

var LimitedFolderUsageCmd = &cobra.Command{
	Use:   "show files over [size] in [directory]",
	Short: "Display directory disk usage with a min size limit",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		size, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			panic(err)
		}
		directory := args[1]
		diskService.ShowFolderSizeWithLimit(directory, size)
	},
}

func init() {
	rootCmd.AddCommand(DiskUsageCmd, FolderUsageCmd)
}
