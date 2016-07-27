package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/template"

	_ "github.com/lib/pq"
)

// A Command is an implementation of a tablarian command
type Command struct {
	// Run runs the command.
	// The args are the arguments after the command name.
	Run func(args []string) int

	// UsageLine is the one-line usage message.
	// The first word in the line is taken to be the command name.
	UsageLine string

	// Short is the short description shown in the 'tablarian help' output.
	Short string

	// Long is the long message shown in the 'tablarian help <this-command>' output.
	Long string

	// Flag is a set of flags specific to this command.
	Flag flag.FlagSet
}

type out struct {
	out io.Writer
	err io.Writer
}

type baseOption struct {
	configFile  string
	prettyPrint bool
}

var o = &out{out: os.Stdout, err: os.Stderr}
var conv Converter

// Name returns the command's name: the first word in the usage line.
func (c *Command) Name() string {
	name := c.UsageLine
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

// Usage prints usage.
func (c *Command) Usage() {
	fmt.Fprintf(o.err, "usage: %s\n\n", c.UsageLine)
	fmt.Fprintf(o.err, "%s\n", strings.TrimSpace(c.Long))
	os.Exit(2)
}

// Commands lists the available commands and help topics.
// The order here is the order in which they are printed by 'tablarian help'.
var commands = []*Command{
	cmdVersion,
	cmdShow,
	cmdOut,
	cmdIndex,
}

func main() {

	flag.Usage = usage
	flag.Parse()
	log.SetFlags(0)

	args := flag.Args()
	if len(args) < 1 {
		usage()
	}

	if args[0] == "help" {
		help(args[1:])
		return
	}

	for _, cmd := range commands {
		if cmd.Name() == args[0] {
			cmd.Flag.Usage = func() { cmd.Usage() }

			cmd.Flag.Parse(args[1:])
			args = cmd.Flag.Args()

			os.Exit(cmd.Run(args))
		}
	}

	fmt.Fprintf(o.err, "tablarian: unknown subcommand %q\nRun ' tablarian help' for usage.\n", args[0])
	os.Exit(2)
}

var usageTemplate = `tablarian is librarian for tables.

Usage:

	tablarian command [arguments]

The commands are:
{{range .}}
	{{.Name | printf "%-11s"}} {{.Short}}{{end}}

Use "tablarian help [command]" for more information about a command.

`

var helpTemplate = `usage: tablarian {{.UsageLine}}

{{.Long | trim}}
`

// tmpl executes the given template text on data, writing the result to w.
func tmpl(w io.Writer, text string, data interface{}) {
	t := template.New("top")
	t.Funcs(template.FuncMap{"trim": strings.TrimSpace})
	template.Must(t.Parse(text))
	if err := t.Execute(w, data); err != nil {
		panic(err)
	}
}

func printUsage(w io.Writer) {
	bw := bufio.NewWriter(w)
	tmpl(bw, usageTemplate, commands)
	bw.Flush()
}

func usage() {
	printUsage(o.err)
	os.Exit(2)
}

// help implements the 'help' command.
func help(args []string) {
	if len(args) == 0 {
		printUsage(o.out)
		// not exit 2: succeeded at 'tablarian help'.
		return
	}
	if len(args) != 1 {
		fmt.Fprintf(o.err, "usage: tablarian help command\n\nToo many arguments given.\n")
		os.Exit(2) // failed at 'tablarian help'
	}

	arg := args[0]

	for _, cmd := range commands {
		if cmd.Name() == arg {
			tmpl(o.out, helpTemplate, cmd)
			// not exit 2: succeeded at 'tablarian help cmd'.
			return
		}
	}

	fmt.Fprintf(o.err, "Unknown help topic %#q.  Run 'tablarian help'.\n", arg)
	os.Exit(2) // failed at 'tablarian help cmd'
}
