package main

import (
	"fmt"

	"github.com/pinzolo/dbmodel"
)

var cmdIndex = &Command{
	Run:       runIndex,
	UsageLine: "index ",
	Short:     "Enumerate all table names in schema.",
	Long: `

	`,
}

func init() {
	cmdIndex.Flag.StringVar(&configFile, "config", "tablarian.config", "Config file path")
	cmdIndex.Flag.StringVar(&configFile, "c", "tablarian.config", "Config file path")
}

// runIndex executes index command and return exit code.
func runIndex(args []string) int {
	cfg, err := loadConfig(configFile)
	if err != nil {
		fmt.Fprintln(o.err, err)
		return 1
	}
	db := dbClientFor(cfg)
	db.Connect()
	defer db.Disconnect()

	tables, err := db.AllTableNames(cfg.Schema)
	if err != nil {
		fmt.Fprintln(o.err, err)
		return 1
	}

	printTableNames(tables, cfg)
	return 0
}

func printTableNames(tables []*dbmodel.Table, cfg *Config) {
	for _, tbl := range tables {
		fmt.Fprintln(o.out, tbl.Name())
	}
}
