package main

import (
	"strconv"
	"strings"

	"github.com/pinzolo/dbmodel"
)

// Converter is data Converter that Convert to table data from dbmodel.
type Converter interface {
	ConvertColumn(*dbmodel.Column) []string
	ConvertIndex(*dbmodel.Index) []string
	ConvertConstraint(*dbmodel.Constraint) []string
	ConvertForeignKey(*dbmodel.ForeignKey) []string
	ConvertReferencedKey(*dbmodel.ForeignKey) []string
}

func findConverter(pretty bool, driver string) Converter {
	if !pretty {
		return defaultConverter{}
	}

	if driver == "postgres" {
		return postgresPrettyConverter{}
	}

	return defaultConverter{}
}

type defaultConverter struct{}

func (dc defaultConverter) ConvertColumn(c *dbmodel.Column) []string {
	null := "NO"
	if c.IsNullable() {
		null = ""
	}
	var pkPosition string
	if c.PrimaryKeyPosition() > 0 {
		pkPosition = strconv.FormatInt(c.PrimaryKeyPosition(), 10)
	}
	return []string{pkPosition, c.Name(), c.DataType(), c.Size().String(), null, c.DefaultValue(), c.Comment()}
}

func (dc defaultConverter) ConvertIndex(idx *dbmodel.Index) []string {
	cols := make([]string, 0, len(idx.Columns()))
	for _, col := range idx.Columns() {
		cols = append(cols, col.Name())
	}
	uniq := ""
	if idx.IsUnique() {
		uniq = "YES"
	}
	return []string{idx.Name(), strings.Join(cols, ", "), uniq}
}

func (dc defaultConverter) ConvertConstraint(con *dbmodel.Constraint) []string {
	return []string{con.Name(), con.Kind(), con.Content()}
}

func (dc defaultConverter) ConvertForeignKey(fk *dbmodel.ForeignKey) []string {
	fCols := make([]string, 0, len(fk.ColumnReferences()))
	tCols := make([]string, 0, len(fk.ColumnReferences()))
	for _, ref := range fk.ColumnReferences() {
		fCols = append(fCols, ref.From().Name())
		tCols = append(tCols, ref.To().Name())
	}
	from := fk.ColumnReferences()[0].From()
	to := fk.ColumnReferences()[0].To()
	tbl := to.TableName()
	if from.Schema() != to.Schema() {
		tbl = to.Schema() + "." + tbl
	}
	return []string{fk.Name(), strings.Join(fCols, ", "), tbl, strings.Join(tCols, ", ")}
}

func (dc defaultConverter) ConvertReferencedKey(rk *dbmodel.ForeignKey) []string {
	fCols := make([]string, 0, len(rk.ColumnReferences()))
	tCols := make([]string, 0, len(rk.ColumnReferences()))
	for _, ref := range rk.ColumnReferences() {
		fCols = append(fCols, ref.From().Name())
		tCols = append(tCols, ref.To().Name())
	}
	from := rk.ColumnReferences()[0].From()
	to := rk.ColumnReferences()[0].To()
	tbl := from.TableName()
	if from.Schema() != to.Schema() {
		tbl = from.Schema() + "." + tbl
	}
	return []string{rk.Name(), tbl, strings.Join(fCols, ", "), strings.Join(tCols, ", ")}
}
