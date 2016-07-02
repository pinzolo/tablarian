package main

var cmdOut = &Command{
	Run:       runOut,
	UsageLine: "out ",
	Short:     "Output table(s) definition(s) as file.",
	Long: `

	`,
}

func init() {
	// Set your flag here like below.
	// cmdOut.Flag.BoolVar(&flagA, "a", false, "")
}

// runOut executes out command and return exit code.
func runOut(args []string) int {

	return 0
}
