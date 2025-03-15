package main

import (
	"fmt"
	"os"

	"github.com/wassef911/lets/cmd/lets/app"
	"github.com/wassef911/lets/pkg"
)

func main() {
	// create services
	logger := pkg.NewLogger()
	diskService := pkg.NewDisk()
	inputOutputService := pkg.NewInputOutput()
	procService := pkg.NewProc()
	searchService := pkg.NewSearch(true)
	rootCmd := app.NewRootCmd(
		logger,
		diskService,
		inputOutputService,
		procService,
		searchService,
	)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
