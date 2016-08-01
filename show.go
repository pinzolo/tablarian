package main

import (
	"fmt"

	"github.com/olekukonko/tablewriter"
	"github.com/pinzolo/dbmodel"
)

type showOption struct {
	baseOption
	showAll bool
}

var (
	cmdShow = &Command{
		Run:       runShow,
		UsageLine: "show [-c] table_name",
		Short:     "Print a table definition.(Console only)",
		Long: `Print table definition to console.

Options:
    -c 'config gile', --config 'config file'
        use config file instead of default config file(tablarian.config)
        if 'config file' starts with '@', it is treated as absolute file path.

    -a, --all
        show all metadata of table.(indices, foreign keys, referenced keys, constraints)
        without this option, print only column definitions.

    -p, --pretty
        convert data type to usually name.
        this option is author's personal option. (only PostgreSQL)
	`,
	}
	showOpt = showOption{}
)

func init() {
	cmdShow.Flag.StringVar(&showOpt.configFile, "config", "tablarian.config", "Config file path")
	cmdShow.Flag.StringVar(&showOpt.configFile, "c", "tablarian.config", "Config file path")
	cmdShow.Flag.BoolVar(&showOpt.showAll, "all", false, "Show all metadata of table")
	cmdShow.Flag.BoolVar(&showOpt.showAll, "a", false, "Show all metadata of table")
	cmdShow.Flag.BoolVar(&showOpt.prettyPrint, "pretty", false, "Pretty print")
	cmdShow.Flag.BoolVar(&showOpt.prettyPrint, "p", false, "Pretty print")
}

// runShow executes show command and return exit code.
func runShow(args []string) int {
	cfg, err := loadConfig(showOpt.configFile)
	if err != nil {
		fmt.Fprintln(o.err, err)
		return 1
	}
	db := dbClientFor(cfg)
	db.Connect()
	defer db.Disconnect()

	opt := dbmodel.RequireNone
	if showOpt.showAll {
		opt = dbmodel.RequireAll
	}
	tbl, err := db.Table(cfg.Schema, args[0], opt)
	if err != nil {
		fmt.Fprintln(o.err, err)
		return 1
	}

	conv := findConverter(showOpt.prettyPrint, cfg.Driver)
	printTable(tbl, conv)
	return 0
}

func printTable(tbl *dbmodel.Table, conv Converter) {
	printColumns(tbl.Columns(), conv)
	printIndices(tbl.Indices(), conv)
	printConstraints(tbl.Constraints(), conv)
	printForeignKeys(tbl.ForeignKeys(), conv)
	printReferencedKyes(tbl.ReferencedKeys(), conv)
}

func printColumns(cols []*dbmodel.Column, conv Converter) {
	w := tablewriter.NewWriter(o.out)
	w.SetHeader([]string{"PK", "NAME", "TYPE", "SIZE", "NULL", "DEFAULT", "COMMENT"})
	w.SetAutoWrapText(false)
	for _, col := range cols {
		w.Append(conv.ConvertColumn(col))
	}
	w.Render()
}

func printIndices(idxs []*dbmodel.Index, conv Converter) {
	if len(idxs) == 0 {
		return
	}
	fmt.Fprintln(o.out)
	fmt.Fprintln(o.out, "### Indices")
	w := tablewriter.NewWriter(o.out)
	w.SetHeader([]string{"NAME", "COLUMNS", "UNIQUE"})
	w.SetAutoWrapText(false)
	for _, idx := range idxs {
		w.Append(conv.ConvertIndex(idx))
	}
	w.Render()
}

func printConstraints(cons []*dbmodel.Constraint, conv Converter) {
	if len(cons) == 0 {
		return
	}
	fmt.Fprintln(o.out)
	fmt.Fprintln(o.out, "### Constraints")
	w := tablewriter.NewWriter(o.out)
	w.SetHeader([]string{"NAME", "KIND", "CONTENT"})
	w.SetAutoWrapText(false)
	for _, con := range cons {
		w.Append(conv.ConvertConstraint(con))
	}
	w.Render()
}

func printForeignKeys(fks []*dbmodel.ForeignKey, conv Converter) {
	if len(fks) == 0 {
		return
	}
	fmt.Fprintln(o.out)
	fmt.Fprintln(o.out, "### Foreign keys")
	w := tablewriter.NewWriter(o.out)
	w.SetHeader([]string{"NAME", "COLUMNS", "FOREIGN TABLE", "FOREIGN COLUMNS"})
	w.SetAutoWrapText(false)
	for _, fk := range fks {
		w.Append(conv.ConvertForeignKey(fk))
	}
	w.Render()
}

func printReferencedKyes(rks []*dbmodel.ForeignKey, conv Converter) {
	if len(rks) == 0 {
		return
	}
	fmt.Fprintln(o.out)
	fmt.Fprintln(o.out, "### Referenced keys")
	w := tablewriter.NewWriter(o.out)
	w.SetHeader([]string{"NAME", "SOURCE TABLE", "SOURCE COLUMNS", "COLUMNS"})
	w.SetAutoWrapText(false)
	for _, rk := range rks {
		w.Append(conv.ConvertReferencedKey(rk))
	}
	w.Render()
}
