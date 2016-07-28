package main

import (
	"database/sql"
	"testing"

	"github.com/pinzolo/dbmodel"
)

func TestInt4ToSerial(t *testing.T) {
	conv := postgresPrettyConverter{}
	size := dbmodel.NewSize(
		sql.NullInt64{Valid: false},
		sql.NullInt64{Int64: 32, Valid: true},
		sql.NullInt64{Int64: 0, Valid: true},
	)
	col := dbmodel.NewColumn("foo", "items", "id", "Primary key of items", "int4", size, false, "nextval('items_id_seq'::regclass)", 0)
	data := conv.ConvertColumn(&col)
	if a, e := data[0], ""; a != e {
		t.Errorf("Primary key should be empty when PrimaryKeyPosition is 0. actual: %v", a)
	}
	if a, e := data[1], "id"; a != e {
		t.Errorf("Second value should be column name. expected: %v, actual: %v", e, a)
	}
	if a, e := data[2], "serial"; a != e {
		t.Errorf("Third value should be serial. expected: %v, actual: %v", e, a)
	}
	if a, e := data[3], ""; a != e {
		t.Errorf("Fourth value should be empty on data type is serial. expected: %v, actual: %v", e, a)
	}
	if a, e := data[4], "NO"; a != e {
		t.Errorf("Fifth value should be nullable and value should be 'NO' when column is not nullable. expected: %v, actual: %v", e, a)
	}
	if a, e := data[5], ""; a != e {
		t.Errorf("Sixth value should be empty on data type is serial. expected: %v, actual: %v", e, a)
	}
	if a, e := data[6], "Primary key of items"; a != e {
		t.Errorf("Seventh value should be comment. expected: %v, actual: %v", e, a)
	}
}

func TestInt4ToSerialLongTableAndColumn(t *testing.T) {
	conv := postgresPrettyConverter{}
	size := dbmodel.NewSize(
		sql.NullInt64{Valid: false},
		sql.NullInt64{Int64: 32, Valid: true},
		sql.NullInt64{Int64: 0, Valid: true},
	)
	col := dbmodel.NewColumn("foo", "bank_output_record_user_sales_detail", "bank_output_record_user_sales_detail_id", "Primary key of items", "int4", size, false, "nextval('bank_output_record_user_sales_bank_output_record_user_sales_seq'::regclass)", 0)
	data := conv.ConvertColumn(&col)
	if a, e := data[2], "serial"; a != e {
		t.Errorf("Third value should be serial. expected: %v, actual: %v", e, a)
	}
	if a, e := data[5], ""; a != e {
		t.Errorf("Sixth value should be empty on data type is serial. expected: %v, actual: %v", e, a)
	}
}

func TestInt4ToSerialLongTable(t *testing.T) {
	conv := postgresPrettyConverter{}
	size := dbmodel.NewSize(
		sql.NullInt64{Valid: false},
		sql.NullInt64{Int64: 32, Valid: true},
		sql.NullInt64{Int64: 0, Valid: true},
	)
	col := dbmodel.NewColumn("foo", "long_long_table_name_abcd_efgh_ijkl_mnop_qrst_uvwx_yz12_3456_7890", "id", "Primary key of items", "int4", size, false, "nextval('long_long_table_name_abcd_efgh_ijkl_mnop_qrst_uvwx_yz12__id_seq'::regclass)", 0)
	data := conv.ConvertColumn(&col)
	if a, e := data[2], "serial"; a != e {
		t.Errorf("Third value should be serial. expected: %v, actual: %v", e, a)
	}
	if a, e := data[5], ""; a != e {
		t.Errorf("Sixth value should be empty on data type is serial. expected: %v, actual: %v", e, a)
	}
}

func TestInt4ToSerialLongColumn(t *testing.T) {
	conv := postgresPrettyConverter{}
	size := dbmodel.NewSize(
		sql.NullInt64{Valid: false},
		sql.NullInt64{Int64: 32, Valid: true},
		sql.NullInt64{Int64: 0, Valid: true},
	)
	col := dbmodel.NewColumn("foo", "st", "long_long_table_name_abcd_efgh_ijkl_mnop_qrst_uvwx_yz12_3456_7890", "Primary key of items", "int4", size, false, "nextval('st_long_long_table_name_abcd_efgh_ijkl_mnop_qrst_uvwx_yz12__seq'::regclass)", 0)
	data := conv.ConvertColumn(&col)
	if a, e := data[2], "serial"; a != e {
		t.Errorf("Third value should be serial. expected: %v, actual: %v", e, a)
	}
	if a, e := data[5], ""; a != e {
		t.Errorf("Sixth value should be empty on data type is serial. expected: %v, actual: %v", e, a)
	}
}

func TestInt4ToInteger(t *testing.T) {
	conv := postgresPrettyConverter{}
	size := dbmodel.NewSize(
		sql.NullInt64{Valid: false},
		sql.NullInt64{Int64: 32, Valid: true},
		sql.NullInt64{Int64: 0, Valid: true},
	)
	col := dbmodel.NewColumn("foo", "items", "id", "Primary key of items", "int4", size, false, "", 0)
	data := conv.ConvertColumn(&col)
	if a, e := data[0], ""; a != e {
		t.Errorf("Primary key should be empty when PrimaryKeyPosition is 0. actual: %v", a)
	}
	if a, e := data[1], "id"; a != e {
		t.Errorf("Second value should be column name. expected: %v, actual: %v", e, a)
	}
	if a, e := data[2], "integer"; a != e {
		t.Errorf("Third value should be serial. expected: %v, actual: %v", e, a)
	}
	if a, e := data[3], ""; a != e {
		t.Errorf("Fourth value should be empty on data type is integer. expected: %v, actual: %v", e, a)
	}
	if a, e := data[4], "NO"; a != e {
		t.Errorf("Fifth value should be nullable and value should be 'NO' when column is not nullable. expected: %v, actual: %v", e, a)
	}
	if a, e := data[5], ""; a != e {
		t.Errorf("Sixth value should be default value. expected: %v, actual: %v", e, a)
	}
	if a, e := data[6], "Primary key of items"; a != e {
		t.Errorf("Seventh value should be comment. expected: %v, actual: %v", e, a)
	}
}

func TestInt4ToIntegerWithOtherSeq(t *testing.T) {
	conv := postgresPrettyConverter{}
	size := dbmodel.NewSize(
		sql.NullInt64{Valid: false},
		sql.NullInt64{Int64: 32, Valid: true},
		sql.NullInt64{Int64: 0, Valid: true},
	)
	col := dbmodel.NewColumn("foo", "items", "id", "Primary key of items", "int4", size, false, "nextval('other_seq'::regclass)", 0)
	data := conv.ConvertColumn(&col)
	if a, e := data[2], "integer"; a != e {
		t.Errorf("Third value should be serial. expected: %v, actual: %v", e, a)
	}
	if a, e := data[5], "nextval('other_seq'::regclass)"; a != e {
		t.Errorf("Sixth value should be default value. expected: %v, actual: %v", e, a)
	}
}

func TestInt2ToSmallserial(t *testing.T) {
	conv := postgresPrettyConverter{}
	size := dbmodel.NewSize(
		sql.NullInt64{Valid: false},
		sql.NullInt64{Int64: 16, Valid: true},
		sql.NullInt64{Int64: 0, Valid: true},
	)
	col := dbmodel.NewColumn("foo", "items", "id", "Primary key of items", "int2", size, false, "nextval('items_id_seq'::regclass)", 0)
	data := conv.ConvertColumn(&col)
	if a, e := data[0], ""; a != e {
		t.Errorf("Primary key should be empty when PrimaryKeyPosition is 0. actual: %v", a)
	}
	if a, e := data[1], "id"; a != e {
		t.Errorf("Second value should be column name. expected: %v, actual: %v", e, a)
	}
	if a, e := data[2], "smallserial"; a != e {
		t.Errorf("Third value should be smallserial. expected: %v, actual: %v", e, a)
	}
	if a, e := data[3], ""; a != e {
		t.Errorf("Fourth value should be empty on data type is smallserial. expected: %v, actual: %v", e, a)
	}
	if a, e := data[4], "NO"; a != e {
		t.Errorf("Fifth value should be nullable and value should be 'NO' when column is not nullable. expected: %v, actual: %v", e, a)
	}
	if a, e := data[5], ""; a != e {
		t.Errorf("Sixth value should be empty on data type is serial. expected: %v, actual: %v", e, a)
	}
	if a, e := data[6], "Primary key of items"; a != e {
		t.Errorf("Seventh value should be comment. expected: %v, actual: %v", e, a)
	}
}

func TestInt2ToSmallserialLongTableAndColumn(t *testing.T) {
	conv := postgresPrettyConverter{}
	size := dbmodel.NewSize(
		sql.NullInt64{Valid: false},
		sql.NullInt64{Int64: 16, Valid: true},
		sql.NullInt64{Int64: 0, Valid: true},
	)
	col := dbmodel.NewColumn("foo", "bank_output_record_user_sales_detail", "bank_output_record_user_sales_detail_id", "Primary key of items", "int2", size, false, "nextval('bank_output_record_user_sales_bank_output_record_user_sales_seq'::regclass)", 0)
	data := conv.ConvertColumn(&col)
	if a, e := data[2], "smallserial"; a != e {
		t.Errorf("Third value should be smallserial. expected: %v, actual: %v", e, a)
	}
	if a, e := data[5], ""; a != e {
		t.Errorf("Sixth value should be empty on data type is smallserial. expected: %v, actual: %v", e, a)
	}
}

func TestInt2ToSmallserialLongTable(t *testing.T) {
	conv := postgresPrettyConverter{}
	size := dbmodel.NewSize(
		sql.NullInt64{Valid: false},
		sql.NullInt64{Int64: 16, Valid: true},
		sql.NullInt64{Int64: 0, Valid: true},
	)
	col := dbmodel.NewColumn("foo", "long_long_table_name_abcd_efgh_ijkl_mnop_qrst_uvwx_yz12_3456_7890", "id", "Primary key of items", "int2", size, false, "nextval('long_long_table_name_abcd_efgh_ijkl_mnop_qrst_uvwx_yz12__id_seq'::regclass)", 0)
	data := conv.ConvertColumn(&col)
	if a, e := data[2], "smallserial"; a != e {
		t.Errorf("Third value should be smallserial. expected: %v, actual: %v", e, a)
	}
	if a, e := data[5], ""; a != e {
		t.Errorf("Sixth value should be empty on data type is smallserial. expected: %v, actual: %v", e, a)
	}
}

func TestInt2ToSmallserialLongColumn(t *testing.T) {
	conv := postgresPrettyConverter{}
	size := dbmodel.NewSize(
		sql.NullInt64{Valid: false},
		sql.NullInt64{Int64: 16, Valid: true},
		sql.NullInt64{Int64: 0, Valid: true},
	)
	col := dbmodel.NewColumn("foo", "st", "long_long_table_name_abcd_efgh_ijkl_mnop_qrst_uvwx_yz12_3456_7890", "Primary key of items", "int2", size, false, "nextval('st_long_long_table_name_abcd_efgh_ijkl_mnop_qrst_uvwx_yz12__seq'::regclass)", 0)
	data := conv.ConvertColumn(&col)
	if a, e := data[2], "smallserial"; a != e {
		t.Errorf("Third value should be smallserial. expected: %v, actual: %v", e, a)
	}
	if a, e := data[5], ""; a != e {
		t.Errorf("Sixth value should be empty on data type is smallserial. expected: %v, actual: %v", e, a)
	}
}

func TestInt2ToSmallint(t *testing.T) {
	conv := postgresPrettyConverter{}
	size := dbmodel.NewSize(
		sql.NullInt64{Valid: false},
		sql.NullInt64{Int64: 16, Valid: true},
		sql.NullInt64{Int64: 0, Valid: true},
	)
	col := dbmodel.NewColumn("foo", "items", "id", "Primary key of items", "int2", size, false, "", 0)
	data := conv.ConvertColumn(&col)
	if a, e := data[0], ""; a != e {
		t.Errorf("Primary key should be empty when PrimaryKeyPosition is 0. actual: %v", a)
	}
	if a, e := data[1], "id"; a != e {
		t.Errorf("Second value should be column name. expected: %v, actual: %v", e, a)
	}
	if a, e := data[2], "smallint"; a != e {
		t.Errorf("Third value should be smallint. expected: %v, actual: %v", e, a)
	}
	if a, e := data[3], ""; a != e {
		t.Errorf("Fourth value should be empty on data type is smallint. expected: %v, actual: %v", e, a)
	}
	if a, e := data[4], "NO"; a != e {
		t.Errorf("Fifth value should be nullable and value should be 'NO' when column is not nullable. expected: %v, actual: %v", e, a)
	}
	if a, e := data[5], ""; a != e {
		t.Errorf("Sixth value should be default value. expected: %v, actual: %v", e, a)
	}
	if a, e := data[6], "Primary key of items"; a != e {
		t.Errorf("Seventh value should be comment. expected: %v, actual: %v", e, a)
	}
}

func TestInt2ToSmallintWithOtherSeq(t *testing.T) {
	conv := postgresPrettyConverter{}
	size := dbmodel.NewSize(
		sql.NullInt64{Valid: false},
		sql.NullInt64{Int64: 16, Valid: true},
		sql.NullInt64{Int64: 0, Valid: true},
	)
	col := dbmodel.NewColumn("foo", "items", "id", "Primary key of items", "int2", size, false, "nextval('other_seq'::regclass)", 0)
	data := conv.ConvertColumn(&col)
	if a, e := data[2], "smallint"; a != e {
		t.Errorf("Third value should be smallint. expected: %v, actual: %v", e, a)
	}
	if a, e := data[5], "nextval('other_seq'::regclass)"; a != e {
		t.Errorf("Sixth value should be default value. expected: %v, actual: %v", e, a)
	}
}

func TestInt8ToBigserial(t *testing.T) {
	conv := postgresPrettyConverter{}
	size := dbmodel.NewSize(
		sql.NullInt64{Valid: false},
		sql.NullInt64{Int64: 64, Valid: true},
		sql.NullInt64{Int64: 0, Valid: true},
	)
	col := dbmodel.NewColumn("foo", "items", "id", "Primary key of items", "int8", size, false, "nextval('items_id_seq'::regclass)", 0)
	data := conv.ConvertColumn(&col)
	if a, e := data[0], ""; a != e {
		t.Errorf("Primary key should be empty when PrimaryKeyPosition is 0. actual: %v", a)
	}
	if a, e := data[1], "id"; a != e {
		t.Errorf("Second value should be column name. expected: %v, actual: %v", e, a)
	}
	if a, e := data[2], "bigserial"; a != e {
		t.Errorf("Third value should be bigserial. expected: %v, actual: %v", e, a)
	}
	if a, e := data[3], ""; a != e {
		t.Errorf("Fourth value should be empty on data type is bigserial. expected: %v, actual: %v", e, a)
	}
	if a, e := data[4], "NO"; a != e {
		t.Errorf("Fifth value should be nullable and value should be 'NO' when column is not nullable. expected: %v, actual: %v", e, a)
	}
	if a, e := data[5], ""; a != e {
		t.Errorf("Sixth value should be empty on data type is serial. expected: %v, actual: %v", e, a)
	}
	if a, e := data[6], "Primary key of items"; a != e {
		t.Errorf("Seventh value should be comment. expected: %v, actual: %v", e, a)
	}
}

func TestInt8ToBigserialLongTableAndColumn(t *testing.T) {
	conv := postgresPrettyConverter{}
	size := dbmodel.NewSize(
		sql.NullInt64{Valid: false},
		sql.NullInt64{Int64: 64, Valid: true},
		sql.NullInt64{Int64: 0, Valid: true},
	)
	col := dbmodel.NewColumn("foo", "bank_output_record_user_sales_detail", "bank_output_record_user_sales_detail_id", "Primary key of items", "int8", size, false, "nextval('bank_output_record_user_sales_bank_output_record_user_sales_seq'::regclass)", 0)
	data := conv.ConvertColumn(&col)
	if a, e := data[2], "bigserial"; a != e {
		t.Errorf("Third value should be bigserial. expected: %v, actual: %v", e, a)
	}
	if a, e := data[5], ""; a != e {
		t.Errorf("Sixth value should be empty on data type is bigserial. expected: %v, actual: %v", e, a)
	}
}

func TestInt8ToBigserialLongTable(t *testing.T) {
	conv := postgresPrettyConverter{}
	size := dbmodel.NewSize(
		sql.NullInt64{Valid: false},
		sql.NullInt64{Int64: 64, Valid: true},
		sql.NullInt64{Int64: 0, Valid: true},
	)
	col := dbmodel.NewColumn("foo", "long_long_table_name_abcd_efgh_ijkl_mnop_qrst_uvwx_yz12_3456_7890", "id", "Primary key of items", "int8", size, false, "nextval('long_long_table_name_abcd_efgh_ijkl_mnop_qrst_uvwx_yz12__id_seq'::regclass)", 0)
	data := conv.ConvertColumn(&col)
	if a, e := data[2], "bigserial"; a != e {
		t.Errorf("Third value should be bigserial. expected: %v, actual: %v", e, a)
	}
	if a, e := data[5], ""; a != e {
		t.Errorf("Sixth value should be empty on data type is bigserial. expected: %v, actual: %v", e, a)
	}
}

func TestInt8ToBigserialLongColumn(t *testing.T) {
	conv := postgresPrettyConverter{}
	size := dbmodel.NewSize(
		sql.NullInt64{Valid: false},
		sql.NullInt64{Int64: 64, Valid: true},
		sql.NullInt64{Int64: 0, Valid: true},
	)
	col := dbmodel.NewColumn("foo", "st", "long_long_table_name_abcd_efgh_ijkl_mnop_qrst_uvwx_yz12_3456_7890", "Primary key of items", "int8", size, false, "nextval('st_long_long_table_name_abcd_efgh_ijkl_mnop_qrst_uvwx_yz12__seq'::regclass)", 0)
	data := conv.ConvertColumn(&col)
	if a, e := data[2], "bigserial"; a != e {
		t.Errorf("Third value should be bigserial. expected: %v, actual: %v", e, a)
	}
	if a, e := data[5], ""; a != e {
		t.Errorf("Sixth value should be empty on data type is bigserial. expected: %v, actual: %v", e, a)
	}
}

func TestInt8ToBigint(t *testing.T) {
	conv := postgresPrettyConverter{}
	size := dbmodel.NewSize(
		sql.NullInt64{Valid: false},
		sql.NullInt64{Int64: 64, Valid: true},
		sql.NullInt64{Int64: 0, Valid: true},
	)
	col := dbmodel.NewColumn("foo", "items", "id", "Primary key of items", "int8", size, false, "", 0)
	data := conv.ConvertColumn(&col)
	if a, e := data[0], ""; a != e {
		t.Errorf("Primary key should be empty when PrimaryKeyPosition is 0. actual: %v", a)
	}
	if a, e := data[1], "id"; a != e {
		t.Errorf("Second value should be column name. expected: %v, actual: %v", e, a)
	}
	if a, e := data[2], "bigint"; a != e {
		t.Errorf("Third value should be bigint. expected: %v, actual: %v", e, a)
	}
	if a, e := data[3], ""; a != e {
		t.Errorf("Fourth value should be empty on data type is bigint. expected: %v, actual: %v", e, a)
	}
	if a, e := data[4], "NO"; a != e {
		t.Errorf("Fifth value should be nullable and value should be 'NO' when column is not nullable. expected: %v, actual: %v", e, a)
	}
	if a, e := data[5], ""; a != e {
		t.Errorf("Sixth value should be default value. expected: %v, actual: %v", e, a)
	}
	if a, e := data[6], "Primary key of items"; a != e {
		t.Errorf("Seventh value should be comment. expected: %v, actual: %v", e, a)
	}
}

func TestInt8ToBigintWithOtherSeq(t *testing.T) {
	conv := postgresPrettyConverter{}
	size := dbmodel.NewSize(
		sql.NullInt64{Valid: false},
		sql.NullInt64{Int64: 64, Valid: true},
		sql.NullInt64{Int64: 0, Valid: true},
	)
	col := dbmodel.NewColumn("foo", "items", "id", "Primary key of items", "int8", size, false, "nextval('other_seq'::regclass)", 0)
	data := conv.ConvertColumn(&col)
	if a, e := data[2], "bigint"; a != e {
		t.Errorf("Third value should be bigint. expected: %v, actual: %v", e, a)
	}
	if a, e := data[5], "nextval('other_seq'::regclass)"; a != e {
		t.Errorf("Sixth value should be default value. expected: %v, actual: %v", e, a)
	}
}

func TestFloat4ToReal(t *testing.T) {
	conv := postgresPrettyConverter{}
	size := dbmodel.NewSize(
		sql.NullInt64{},
		sql.NullInt64{Int64: 24, Valid: true},
		sql.NullInt64{},
	)
	col := dbmodel.NewColumn("foo", "items", "price", "Price of item", "float4", size, false, "0.0", 0)
	data := conv.ConvertColumn(&col)
	if a, e := data[0], ""; a != e {
		t.Errorf("Primary key should be empty when PrimaryKeyPosition is 0. actual: %v", a)
	}
	if a, e := data[1], "price"; a != e {
		t.Errorf("Second value should be column name. expected: %v, actual: %v", e, a)
	}
	if a, e := data[2], "real"; a != e {
		t.Errorf("Third value should be real. expected: %v, actual: %v", e, a)
	}
	if a, e := data[3], ""; a != e {
		t.Errorf("Fourth value should be empty on data type is real. expected: %v, actual: %v", e, a)
	}
	if a, e := data[4], "NO"; a != e {
		t.Errorf("Fifth value should be nullable and value should be 'NO' when column is not nullable. expected: %v, actual: %v", e, a)
	}
	if a, e := data[5], "0.0"; a != e {
		t.Errorf("Sixth value should be default value. expected: %v, actual: %v", e, a)
	}
	if a, e := data[6], "Price of item"; a != e {
		t.Errorf("Seventh value should be comment. expected: %v, actual: %v", e, a)
	}
}

func TestFloat8ToDoublePrecision(t *testing.T) {
	conv := postgresPrettyConverter{}
	size := dbmodel.NewSize(
		sql.NullInt64{},
		sql.NullInt64{Int64: 53, Valid: true},
		sql.NullInt64{},
	)
	col := dbmodel.NewColumn("foo", "items", "price", "Price of item", "float8", size, false, "0.0", 0)
	data := conv.ConvertColumn(&col)
	if a, e := data[0], ""; a != e {
		t.Errorf("Primary key should be empty when PrimaryKeyPosition is 0. actual: %v", a)
	}
	if a, e := data[1], "price"; a != e {
		t.Errorf("Second value should be column name. expected: %v, actual: %v", e, a)
	}
	if a, e := data[2], "double precision"; a != e {
		t.Errorf("Third value should be double precision. expected: %v, actual: %v", e, a)
	}
	if a, e := data[3], ""; a != e {
		t.Errorf("Fourth value should be empty on data type is real. expected: %v, actual: %v", e, a)
	}
	if a, e := data[4], "NO"; a != e {
		t.Errorf("Fifth value should be nullable and value should be 'NO' when column is not nullable. expected: %v, actual: %v", e, a)
	}
	if a, e := data[5], "0.0"; a != e {
		t.Errorf("Sixth value should be default value. expected: %v, actual: %v", e, a)
	}
	if a, e := data[6], "Price of item"; a != e {
		t.Errorf("Seventh value should be comment. expected: %v, actual: %v", e, a)
	}
}
