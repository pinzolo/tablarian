package main

var cmdPublish = &Command{
	Run:       runPublish,
	UsageLine: "publish ",
	Short:     "Output table(s) definition(s) as file.",
	Long: `

	`,
}

func init() {
	// Set your flag here like below.
	// cmdPublish.Flag.BoolVar(&flagA, "a", false, "")
}

// runPublish executes out command and return exit code.
func runPublish(args []string) int {

	return 0
}
