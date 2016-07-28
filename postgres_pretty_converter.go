package main

import (
	"fmt"

	"github.com/pinzolo/dbmodel"
)

const (
	maxSeqLength = 63
	threshold    = 29
)

type postgresPrettyConverter struct {
	defaultConverter
}

func (ppc postgresPrettyConverter) ConvertColumn(col *dbmodel.Column) []string {
	base := ppc.defaultConverter.ConvertColumn(col)
	dataType := base[2]
	size := base[3]
	defVal := base[5]
	switch dataType {
	case "int2":
		if isSerial(col) {
			dataType = "smallserial"
			defVal = ""
		} else {
			dataType = "smallint"
		}
		size = ""
	case "int4":
		if isSerial(col) {
			dataType = "serial"
			defVal = ""
		} else {
			dataType = "integer"
		}
		size = ""
	case "int8":
		if isSerial(col) {
			dataType = "bigserial"
			defVal = ""
		} else {
			dataType = "bigint"
		}
		size = ""
	case "float4":
		dataType = "real"
		size = ""
	case "float8":
		dataType = "double precision"
		size = ""
	}

	return []string{base[0], base[1], dataType, size, base[4], defVal, base[6]}
}

func isSerial(col *dbmodel.Column) bool {
	tblName := col.TableName()
	colName := col.Name()
	if len(tblName)+len(colName)+5 > maxSeqLength {
		if len(tblName) > threshold {
			if len(colName) > threshold {
				tblName = tblName[0:threshold]
				colName = colName[0:threshold]
			} else {
				len := threshold + threshold - len(colName)
				tblName = tblName[0:len]
			}
		} else {
			len := threshold + threshold - len(tblName)
			colName = colName[0:len]
		}
	}
	return col.DefaultValue() == fmt.Sprintf("nextval('%s_%s_seq'::regclass)", tblName, colName)
}
