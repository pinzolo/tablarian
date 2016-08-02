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
	cnv    Converter
	errors []error
}

func newMarkdownPublisher(config *Config, converter Converter) markdownPublisher {
	return markdownPublisher{
		cfg:    config,
		cnv:    converter,
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
		md := p.convertToMarkdown(tbl)
		if f, err := os.Create(filepath.Join(path, tbl.Name()+".md")); err != nil {
			p.errors = append(p.errors, err)
		} else {
			f.Write([]byte(md))
			err = f.Close()
			if err != nil {
				p.errors = append(p.errors, err)
			}
		}
	}

	idxMd := p.convertToIndexMarkdown(tables)
	if f, err := os.Create(filepath.Join(path, "00_index.md")); err != nil {
		p.errors = append(p.errors, err)
	} else {
		f.Write([]byte(idxMd))
		err = f.Close()
		if err != nil {
			p.errors = append(p.errors, err)
		}
	}
}

func (p markdownPublisher) Errors() []error {
	return p.errors
}

func (p markdownPublisher) convertToIndexMarkdown(tables []*dbmodel.Table) string {
	buf := &bytes.Buffer{}
	nl := fmt.Sprintln()

	fmt.Fprintln(buf, "# Index"+nl)
	for _, tbl := range tables {
		if tbl.Comment() == "" {
			fmt.Fprintf(buf, "* [%s](%s.md)%s", tbl.Name(), tbl.Name(), nl)
		} else {
			fmt.Fprintf(buf, "* [%s](%s.md) : %s%s", tbl.Name(), tbl.Name(), tbl.Comment(), nl)
		}
	}
	return buf.String()
}

func (p markdownPublisher) convertToMarkdown(table *dbmodel.Table) string {
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
		w.Append(p.cnv.ConvertColumn(col))
	}
	w.Render()

	if len(table.Indices()) > 0 {
		fmt.Fprintf(buf, "%s## Indices%s%s", nl, nl, nl)
		w = newMdTableWriter(buf)
		w.SetHeader([]string{"NAME", "COLUMNS", "UNIQUE"})
		for _, idx := range table.Indices() {
			w.Append(p.cnv.ConvertIndex(idx))
		}
		w.Render()
	}

	if len(table.Constraints()) > 0 {
		fmt.Fprintf(buf, "%s## Constants%s%s", nl, nl, nl)
		w = newMdTableWriter(buf)
		w.SetHeader([]string{"NAME", "KIND", "CONTENT"})
		for _, con := range table.Constraints() {
			w.Append(p.cnv.ConvertConstraint(con))
		}
		w.Render()
	}

	if len(table.ForeignKeys()) > 0 {
		fmt.Fprintf(buf, "%s## Foreign keys%s%s", nl, nl, nl)
		w = newMdTableWriter(buf)
		w.SetHeader([]string{"NAME", "COLUMNS", "FOREIGN TABLE", "FOREIGN COLUMNS"})
		for _, fk := range table.ForeignKeys() {
			w.Append(p.cnv.ConvertForeignKey(fk))
		}
		w.Render()
	}

	if len(table.ReferencedKeys()) > 0 {
		fmt.Fprintf(buf, "%s## Referenced keys%s%s", nl, nl, nl)
		w = newMdTableWriter(buf)
		w.SetHeader([]string{"NAME", "SOURCE TABLE", "SOURCE COLUMNS", "COLUMNS"})
		for _, rk := range table.ReferencedKeys() {
			w.Append(p.cnv.ConvertReferencedKey(rk))
		}
		w.Render()
	}

	return buf.String()
}

func newMdTableWriter(w io.Writer) *tablewriter.Table {
	tw := tablewriter.NewWriter(w)
	tw.SetAutoWrapText(false)
	tw.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	tw.SetCenterSeparator("|")
	return tw
}

func cleanDir(path string) error {
	if err := os.RemoveAll(path); err != nil {
		return err
	}
	if err := os.MkdirAll(path, 0777); err != nil {
		return err
	}
	return nil
}
