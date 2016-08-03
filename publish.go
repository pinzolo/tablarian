package main

import (
	"fmt"

	"github.com/pinzolo/dbmodel"
)

type publishOption struct {
	baseOption
	format string
}

var (
	cmdPublish = &Command{
		Run:       runPublish,
		UsageLine: "publish ",
		Short:     "Output definition of tables to file.",
		Long: `Output definition of tables to file.

Options:
    -c CONGIG_FILE, --config CONFIG_FILE
        use config file instead of default config file(tablarian.config)
        if CONFIG_FILE starts with '@', it is treated as absolute file path.

    -p, --pretty
        convert data type to usually name.
        this option is author's personal option. (only PostgreSQL)

    -f, --format
        file format for saving table definitions.
        formats:
            markdown (default)
	`,
	}
	publishOpt = publishOption{}
)

func init() {
	cmdPublish.Flag.StringVar(&publishOpt.configFile, "config", "tablarian.config", "Config file path")
	cmdPublish.Flag.StringVar(&publishOpt.configFile, "c", "tablarian.config", "Config file path")
	cmdPublish.Flag.BoolVar(&publishOpt.prettyPrint, "pretty", false, "Pretty print")
	cmdPublish.Flag.BoolVar(&publishOpt.prettyPrint, "p", false, "Pretty print")
	cmdPublish.Flag.StringVar(&publishOpt.format, "format", "markdown", "File format")
	cmdPublish.Flag.StringVar(&publishOpt.format, "f", "markdown", "File format")
}

// runPublish executes out command and return exit code.
func runPublish(args []string) int {
	cfg, err := loadConfig(publishOpt.configFile)
	if err != nil {
		fmt.Fprintln(o.err, err)
		return 1
	}

	db := dbClientFor(cfg)
	db.Connect()
	defer db.Disconnect()

	tables, err := db.AllTables(cfg.Schema, dbmodel.RequireAll)
	if err != nil {
		fmt.Fprintln(o.err, err)
		return 1
	}
	conv := findConverter(publishOpt.prettyPrint, cfg.Driver)
	pub := findPublisher(publishOpt.format, cfg, conv)
	pub.Publish(tables)
	if len(pub.Errors()) > 0 {
		for _, err = range pub.Errors() {
			fmt.Fprintln(o.err, err)
		}
		return 1
	}

	return 0
}
