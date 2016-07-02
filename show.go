package main

var cmdShow = &Command{
	Run:       runShow,
	UsageLine: "show ",
	Short:     "Print a table definition.(Console only)",
	Long: `

	`,
}

func init() {
	// Set your flag here like below.
	// cmdShow.Flag.BoolVar(&flagA, "a", false, "")
}

// runShow executes show command and return exit code.
func runShow(args []string) int {

	return 0
}
