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
	loc    locale
	errors []error
}

func newMarkdownPublisher(config *Config, converter Converter, locale locale) markdownPublisher {
	return markdownPublisher{
		cfg:    config,
		conv:   converter,
		loc:    locale,
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
		md := convertToMarkdown(tbl, p.conv, p.loc)
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

	idxMd := convertToIndexMarkdown(tables, p.loc)
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

func convertToMarkdown(table *dbmodel.Table, conv Converter, loc locale) []byte {
	buf := &bytes.Buffer{}

	fmt.Fprintln(buf, "#", table.Name())
	if table.Comment() != "" {
		fmt.Fprintln(buf)
		fmt.Fprintln(buf, table.Comment())
	}

	fmt.Fprintln(buf)
	fmt.Fprintln(buf, "##", loc.t("column", "title"))
	fmt.Fprintln(buf)
	w := newMdTableWriter(buf)
	w.SetHeader(tHeaders(loc, "column", "primary_key", "name", "data_type", "size", "null", "default_value", "comment"))
	for _, col := range table.Columns() {
		w.Append(conv.ConvertColumn(col))
	}
	w.Render()

	if len(table.Indices()) > 0 {
		fmt.Fprintln(buf)
		fmt.Fprintln(buf, "##", loc.t("index", "title"))
		fmt.Fprintln(buf)
		w = newMdTableWriter(buf)
		w.SetHeader(tHeaders(loc, "index", "name", "columns", "unique"))
		for _, idx := range table.Indices() {
			w.Append(conv.ConvertIndex(idx))
		}
		w.Render()
	}

	if len(table.Constraints()) > 0 {
		fmt.Fprintln(buf)
		fmt.Fprintln(buf, "##", loc.t("constraint", "title"))
		fmt.Fprintln(buf)
		w = newMdTableWriter(buf)
		w.SetHeader(tHeaders(loc, "constraint", "name", "kind", "content"))
		for _, con := range table.Constraints() {
			w.Append(conv.ConvertConstraint(con))
		}
		w.Render()
	}

	if len(table.ForeignKeys()) > 0 {
		fmt.Fprintln(buf)
		fmt.Fprintln(buf, "##", loc.t("foreign_key", "title"))
		fmt.Fprintln(buf)
		w = newMdTableWriter(buf)
		w.SetHeader(tHeaders(loc, "foreign_key", "name", "columns", "foreign_table", "foreign_columns"))
		for _, fk := range table.ForeignKeys() {
			w.Append(conv.ConvertForeignKey(fk))
		}
		w.Render()
	}

	if len(table.ReferencedKeys()) > 0 {
		fmt.Fprintln(buf)
		fmt.Fprintln(buf, "##", loc.t("referenced_key", "title"))
		fmt.Fprintln(buf)
		w = newMdTableWriter(buf)
		w.SetHeader(tHeaders(loc, "referenced_key", "name", "source_table", "source_columns", "columns"))
		for _, rk := range table.ReferencedKeys() {
			w.Append(conv.ConvertReferencedKey(rk))
		}
		w.Render()
	}

	return buf.Bytes()
}

func convertToIndexMarkdown(tables []*dbmodel.Table, loc locale) []byte {
	buf := &bytes.Buffer{}

	fmt.Fprintln(buf, "#", loc.t("table_list", "title"))
	fmt.Fprintln(buf)
	w := newMdTableWriter(buf)
	w.SetHeader(tHeaders(loc, "table_list", "table", "comment"))
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

func tHeaders(loc locale, cat string, keys ...string) []string {
	hs := make([]string, 0, len(keys))
	for _, key := range keys {
		hs = append(hs, loc.t(cat, key))
	}
	return hs
}
