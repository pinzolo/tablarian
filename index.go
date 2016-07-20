package main

var cmdIndex = &Command{
	Run:       runIndex,
	UsageLine: "index ",
	Short:     "Enumerate table names in schema.",
	Long: `

	`,
}

func init() {
	// Set your flag here like below.
	// cmdIndex.Flag.BoolVar(&flagA, "a", false, "")
}

// runIndex executes index command and return exit code.
func runIndex(args []string) int {
	return 0
}
