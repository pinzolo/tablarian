package main

import (
	"fmt"
)

// Version is tablarian current version.
var Version = "0.0.1"

var cmdVersion = &Command{
	Run:       runVersion,
	UsageLine: "version ",
	Short:     "Print tablarian current version.",
	Long: `

	`,
}

func init() {
	// Set your flag here like below.
	// cmdVersion.Flag.BoolVar(&flagA, "a", false, "")
}

// runVersion executes version command and return exit code.
func runVersion(args []string) int {
	fmt.Println("tablarian:", Version)
	return 0
}
