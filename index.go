package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/pinzolo/dbmodel"
)

type indexOption struct {
	baseOption
	withoutTableComment bool
}

var (
	cmdIndex = &Command{
		Run:       runIndex,
		UsageLine: "index ",
		Short:     "Print table names to console.",
		Long: `Print table names to console.

Options:
    -c CONGIG_FILE, --config CONFIG_FILE
        use config file instead of default config file(tablarian.config)
        if CONFIG_FILE starts with '@', it is treated as absolute file path.

    -C, --no-comment
        Not print table comment. (default: false)
	`,
	}
	idxOpt = indexOption{}
)

func init() {
	cmdIndex.Flag.StringVar(&idxOpt.configFile, "config", "tablarian.config", "Config file path")
	cmdIndex.Flag.StringVar(&idxOpt.configFile, "c", "tablarian.config", "Config file path")
	cmdIndex.Flag.BoolVar(&idxOpt.withoutTableComment, "no-comment", false, "Without table comment")
	cmdIndex.Flag.BoolVar(&idxOpt.withoutTableComment, "C", false, "Without table comment")
}

// runIndex executes index command and return exit code.
func runIndex(args []string) int {
	cfg, err := loadConfig(idxOpt.configFile)
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

	printTableNames(tables)
	return 0
}

func printTableNames(tables []*dbmodel.Table) {
	for _, line := range tableNamesLines(tables) {
		fmt.Fprintln(o.out, line)
	}
}

func tableNamesLines(tables []*dbmodel.Table) []string {
	buf := &bytes.Buffer{}
	w := tablewriter.NewWriter(buf)
	w.SetBorder(false)
	w.SetColumnSeparator("")
	w.SetAutoWrapText(false)
	for _, tbl := range tables {
		data := []string{tbl.Name()}
		if !idxOpt.withoutTableComment {
			data = append(data, tbl.Comment())
		}
		w.Append(data)
	}
	w.Render()
	return trimEachLines(buf.String())
}

func trimEachLines(tableString string) []string {
	lines := strings.Split(tableString, "\n")
	newLines := make([]string, 0, len(lines))
	for _, line := range lines {
		if l := strings.TrimSpace(line); len(l) > 0 {
			newLines = append(newLines, l)
		}
	}
	return newLines
}
