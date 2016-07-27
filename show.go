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
	`,
	}
	showOpt = showOption{}
)

func init() {
	cmdShow.Flag.StringVar(&showOpt.configFile, "config", "tablarian.config", "Config file path")
	cmdShow.Flag.StringVar(&showOpt.configFile, "c", "tablarian.config", "Config file path")
	cmdShow.Flag.BoolVar(&showOpt.showAll, "all", false, "Show all metadata of table.")
	cmdShow.Flag.BoolVar(&showOpt.showAll, "a", false, "Show all metadata of table.")
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

	printTable(tbl)
	return 0
}

func printTable(tbl *dbmodel.Table) {
	printColumns(tbl.Columns())
	printIndices(tbl.Indices())
	printConstraints(tbl.Constraints())
	printForeignKeys(tbl.ForeignKeys())
	printReferencedKyes(tbl.ReferencedKeys())
}

func printColumns(cols []*dbmodel.Column) {
	w := tablewriter.NewWriter(o.out)
	w.SetHeader([]string{"PK", "NAME", "TYPE", "SIZE", "NULL", "DEFAULT", "COMMENT"})
	w.SetAutoWrapText(false)
	for _, col := range cols {
		w.Append(convertColumn(col))
	}
	w.Render()
}

func printIndices(idxs []*dbmodel.Index) {
	if len(idxs) == 0 {
		return
	}
	fmt.Fprintln(o.out)
	fmt.Fprintln(o.out, "### Indices")
	w := tablewriter.NewWriter(o.out)
	w.SetHeader([]string{"NAME", "COLUMNS", "UNIQUE"})
	w.SetAutoWrapText(false)
	for _, idx := range idxs {
		w.Append(convertIndex(idx))
	}
	w.Render()
}

func printConstraints(cons []*dbmodel.Constraint) {
	if len(cons) == 0 {
		return
	}
	fmt.Fprintln(o.out)
	fmt.Fprintln(o.out, "### Constraints")
	w := tablewriter.NewWriter(o.out)
	w.SetHeader([]string{"NAME", "KIND", "CONTENT"})
	w.SetAutoWrapText(false)
	for _, con := range cons {
		w.Append(convertConstraint(con))
	}
	w.Render()
}

func printForeignKeys(fks []*dbmodel.ForeignKey) {
	if len(fks) == 0 {
		return
	}
	fmt.Fprintln(o.out)
	fmt.Fprintln(o.out, "### Foreign keys")
	w := tablewriter.NewWriter(o.out)
	w.SetHeader([]string{"NAME", "COLUMNS", "FOREIGN TABLE", "FOREIGN COLUMNS"})
	w.SetAutoWrapText(false)
	for _, fk := range fks {
		w.Append(convertForeignKey(fk))
	}
	w.Render()
}

func printReferencedKyes(rks []*dbmodel.ForeignKey) {
	if len(rks) == 0 {
		return
	}
	fmt.Fprintln(o.out)
	fmt.Fprintln(o.out, "### Referenced keys")
	w := tablewriter.NewWriter(o.out)
	w.SetHeader([]string{"NAME", "SOURCE TABLE", "SOURCE COLUMNS", "COLUMNS"})
	w.SetAutoWrapText(false)
	for _, rk := range rks {
		w.Append(convertReferencedKey(rk))
	}
	w.Render()
}
