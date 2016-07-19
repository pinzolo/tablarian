package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/pinzolo/dbmodel"
	"github.com/pinzolo/tablarian/converter"
)

var (
	cmdShow = &Command{
		Run:       runShow,
		UsageLine: "show ",
		Short:     "Print a table definition.(Console only)",
		Long: `

	`,
	}
	configFile string
)

func init() {
	cmdShow.Flag.StringVar(&configFile, "config", "tablarian.config", "Config file path")
	cmdShow.Flag.StringVar(&configFile, "c", "tablarian.config", "Config file path")
}

// runShow executes show command and return exit code.
func runShow(args []string) int {
	cfg, err := loadConfig(configFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	db := dbClientFor(cfg)
	db.Connect()
	defer db.Disconnect()

	tbl, err := db.Table(cfg.Schema, args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	err = printColumns(tbl.Columns(), cfg)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	return 0
}

func printColumns(cols []*dbmodel.Column, cfg *Config) error {
	conv, err := converter.FindConverter(cfg.Driver)
	if err != nil {
		return err
	}
	w := tablewriter.NewWriter(os.Stdout)
	w.SetHeader([]string{"PK", "NAME", "TYPE", "SIZE", "NULL", "DEFAULT", "COMMENT"})
	w.SetAutoWrapText(false)
	for _, col := range cols {
		w.Append(conv.Convert(col))
	}
	w.Render()
	return nil
}
