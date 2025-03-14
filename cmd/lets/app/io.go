package app

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/wassef911/lets/pkg"
)

var inputOutputService = pkg.NewInputOutput()

var getCmd = &cobra.Command{
	Use:   "get column [col] from [csvPath]",
	Short: "Prints the values of a column in a CSV file",
	Args:  cobra.MinimumNArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		csvPath := args[len(args)-1]
		col, err := strconv.Atoi(args[1])
		if err != nil {
			panic(err)
		}
		inputOutputService.GetColumn(csvPath, col)
	},
}

var ReplaceCmd = &cobra.Command{
	Use:   "replace [foo] with [bar] in [filename]",
	Short: "In-place text replacement",
	Args:  cobra.MinimumNArgs(5),
	Run: func(cmd *cobra.Command, args []string) {
		foo := args[0]
		bar := args[2]
		filename := args[len(args)-1]
		inputOutputService.ReplaceText(filename, foo, bar)
	},
}
