package main

import (
	"fmt"
	"io"

	"github.com/pinzolo/dbmodel"
)

type publishOption struct {
	baseOption
	format  string
	locale  string
	verbose bool
}

var (
	cmdPublish = &Command{
		Run:       runPublish,
		UsageLine: "publish ",
		Short:     "Output definition of tables to file.",
		Long: `Output definition of tables to file.

Options:
    -c CONGIG_FILE, --config CONFIG_FILE
        use config file instead of default config file(.tablarian.config)
        if CONFIG_FILE starts with '@', it is treated as absolute file path.

    -p, --pretty
        convert data type to usually name.
        this option is author's personal option. (only PostgreSQL)

    -f, --format
        file format for saving table definitions.
        formats:
            markdown (default)

    -l LOCALE, --locale LOCALE
        use LOCALE instead of default locale(en).
        currently acceptable locales are 'en', 'ja'.

    -v, --verbose
        print verbose log to console.
	`,
	}
	publishOpt = publishOption{}
)

func init() {
	cmdPublish.Flag.StringVar(&publishOpt.configFile, "config", DefaultConfigFileName, "Config file path")
	cmdPublish.Flag.StringVar(&publishOpt.configFile, "c", DefaultConfigFileName, "Config file path")
	cmdPublish.Flag.BoolVar(&publishOpt.prettyPrint, "pretty", false, "Pretty print")
	cmdPublish.Flag.BoolVar(&publishOpt.prettyPrint, "p", false, "Pretty print")
	cmdPublish.Flag.StringVar(&publishOpt.format, "format", "markdown", "File format")
	cmdPublish.Flag.StringVar(&publishOpt.format, "f", "markdown", "File format")
	cmdPublish.Flag.StringVar(&publishOpt.locale, "locale", "en", "Locale")
	cmdPublish.Flag.StringVar(&publishOpt.locale, "l", "en", "Locale")
	cmdPublish.Flag.BoolVar(&publishOpt.verbose, "v", false, "Print log")
	cmdPublish.Flag.BoolVar(&publishOpt.verbose, "verbose", false, "Print log")
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
	var logger io.Writer
	if publishOpt.verbose {
		logger = o.out
	}
	pub, err := findPublisher(publishOpt.format, cfg, conv, l(publishOpt.locale), logger)
	if err != nil {
		fmt.Fprintln(o.err, err)
		return 1
	}
	pub.Publish(tables)
	if len(pub.Errors()) > 0 {
		fmt.Fprintln(o.err, "Error occured!! ==========")
		for _, err = range pub.Errors() {
			fmt.Fprintln(o.err, err)
		}
		return 1
	}

	return 0
}
