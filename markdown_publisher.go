package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/olekukonko/tablewriter"
	"github.com/pinzolo/dbmodel"
)

type markdownPublisher struct {
	cfg    *Config
	conv   Converter
	errors []error
}

func newMarkdownPublisher(config *Config, converter Converter) markdownPublisher {
	return markdownPublisher{
		cfg:    config,
		conv:   converter,
		errors: make([]error, 0, 0),
	}
}

func (p markdownPublisher) Publish(tables []*dbmodel.Table) {
	path, err := resolvePath(p.cfg.Out)
	if err != nil {
		p.errors = append(p.errors, err)
		return
	}
	if err := cleanDir(path); err != nil {
		p.errors = append(p.errors, err)
		return
	}

	for _, tbl := range tables {
		md := convertToMarkdown(tbl, p.conv)
		if f, err := os.Create(filepath.Join(path, tbl.Name()+".md")); err != nil {
			p.errors = append(p.errors, err)
		} else {
			f.Write(md)
			err = f.Close()
			if err != nil {
				p.errors = append(p.errors, err)
			}
		}
	}

	idxMd := convertToIndexMarkdown(tables)
	if f, err := os.Create(filepath.Join(path, "00_index.md")); err != nil {
		p.errors = append(p.errors, err)
	} else {
		f.Write(idxMd)
		err = f.Close()
		if err != nil {
			p.errors = append(p.errors, err)
		}
	}
}

func (p markdownPublisher) Errors() []error {
	return p.errors
}

func convertToMarkdown(table *dbmodel.Table, conv Converter) []byte {
	buf := &bytes.Buffer{}
	nl := fmt.Sprintln()

	fmt.Fprintf(buf, "# %s%s", table.Name(), nl)
	if table.Comment() != "" {
		fmt.Fprintln(buf, nl+table.Comment())
	}

	fmt.Fprintf(buf, "%s## Columns%s%s", nl, nl, nl)
	w := newMdTableWriter(buf)
	w.SetHeader([]string{"PK", "NAME", "TYPE", "SIZE", "NULL", "DEFAULT", "COMMENT"})
	for _, col := range table.Columns() {
		w.Append(conv.ConvertColumn(col))
	}
	w.Render()

	if len(table.Indices()) > 0 {
		fmt.Fprintf(buf, "%s## Indices%s%s", nl, nl, nl)
		w = newMdTableWriter(buf)
		w.SetHeader([]string{"NAME", "COLUMNS", "UNIQUE"})
		for _, idx := range table.Indices() {
			w.Append(conv.ConvertIndex(idx))
		}
		w.Render()
	}

	if len(table.Constraints()) > 0 {
		fmt.Fprintf(buf, "%s## Constants%s%s", nl, nl, nl)
		w = newMdTableWriter(buf)
		w.SetHeader([]string{"NAME", "KIND", "CONTENT"})
		for _, con := range table.Constraints() {
			w.Append(conv.ConvertConstraint(con))
		}
		w.Render()
	}

	if len(table.ForeignKeys()) > 0 {
		fmt.Fprintf(buf, "%s## Foreign keys%s%s", nl, nl, nl)
		w = newMdTableWriter(buf)
		w.SetHeader([]string{"NAME", "COLUMNS", "FOREIGN TABLE", "FOREIGN COLUMNS"})
		for _, fk := range table.ForeignKeys() {
			w.Append(conv.ConvertForeignKey(fk))
		}
		w.Render()
	}

	if len(table.ReferencedKeys()) > 0 {
		fmt.Fprintf(buf, "%s## Referenced keys%s%s", nl, nl, nl)
		w = newMdTableWriter(buf)
		w.SetHeader([]string{"NAME", "SOURCE TABLE", "SOURCE COLUMNS", "COLUMNS"})
		for _, rk := range table.ReferencedKeys() {
			w.Append(conv.ConvertReferencedKey(rk))
		}
		w.Render()
	}

	return buf.Bytes()
}

func convertToIndexMarkdown(tables []*dbmodel.Table) []byte {
	buf := &bytes.Buffer{}
	nl := fmt.Sprintln()

	fmt.Fprintln(buf, "# Index"+nl)
	w := newMdTableWriter(buf)
	w.SetHeader([]string{"TABLE", "COMMENT"})
	for _, tbl := range tables {
		w.Append([]string{fmt.Sprintf("[%s](%s.md)", tbl.Name(), tbl.Name()), tbl.Comment()})
	}
	w.Render()

	return buf.Bytes()
}

func newMdTableWriter(w io.Writer) *tablewriter.Table {
	tw := tablewriter.NewWriter(w)
	tw.SetAutoWrapText(false)
	tw.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	tw.SetCenterSeparator("|")
	return tw
}

func cleanDir(path string) error {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			if err := os.RemoveAll(path); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("Path: %s is file.", path)
		}
	}
	if err := os.MkdirAll(path, 0777); err != nil {
		return err
	}
	return nil
}
