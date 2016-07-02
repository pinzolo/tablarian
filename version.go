package main

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

	return 0
}
