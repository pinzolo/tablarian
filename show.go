package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/pinzolo/dbmodel"
	"github.com/pinzolo/tablarian/converter"

	_ "github.com/lib/pq"
)

var (
	cmdShow = &Command{
		Run:       runShow,
		UsageLine: "show ",
		Short:     "Print a table definition.(Console only)",
		Long: `

	`,
	}
)

func init() {
	// Set your flag here like below.
	// cmdShow.Flag.BoolVar(&flagA, "a", false, "")
}

// runShow executes show command and return exit code.
func runShow(args []string) int {
	cfg, err := loadConfig("tablarian.config")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	client := clientFor(cfg)
	client.Connect()
	defer client.Disconnect()

	tbl, err := client.Table(cfg.Schema, args[0])
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
	w.SetHeader([]string{"NAME", "TYPE", "SIZE", "NULL", "DEFAULT", "COMMENT"})
	w.SetBorder(false)
	for _, col := range cols {
		w.Append(conv.Convert(col))
	}
	w.Render()
	return nil
}
