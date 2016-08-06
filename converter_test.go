package main

import (
	"database/sql"
	"testing"

	"github.com/pinzolo/dbmodel"
)

func TestFindConverterWhenNoPretty(t *testing.T) {
	conv := findConverter(false, "postgres")
	if _, ok := conv.(defaultConverter); !ok {
		t.Error("If pretty is false, findConverter should return defaultConverter.")
	}
}

func TestFindConverterWhenPrettyPostgres(t *testing.T) {
	conv := findConverter(true, "postgres")
	if _, ok := conv.(postgresPrettyConverter); !ok {
		t.Error("If pretty is true and driver is 'postgres', findConverter should return postgresPrettyConverter.")
	}
}

func TestFindConverterWhenUnknownDriver(t *testing.T) {
	conv := findConverter(false, "unknown")
	if _, ok := conv.(defaultConverter); !ok {
		t.Error("If driver is unknows, findConverter should return defaultConverter.")
	}
}

func TestDefaultConvertColumn(t *testing.T) {
	conv := defaultConverter{}
	size := dbmodel.NewSize(
		sql.NullInt64{Valid: false},
		sql.NullInt64{Int64: 10, Valid: true},
		sql.NullInt64{Int64: 2, Valid: true},
	)
	col := dbmodel.NewColumn("foo", "items", "price", "item price", "numeric", size, false, "0.0", 0)
	data := conv.ConvertColumn(&col)
	if a, e := data[0], ""; a != e {
		t.Errorf("Primary key should be empty when PrimaryKeyPosition is 0. actual: %v", a)
	}
	if a, e := data[1], "price"; a != e {
		t.Errorf("Second value should be column name. expected: %v, actual: %v", e, a)
	}
	if a, e := data[2], "numeric"; a != e {
		t.Errorf("Third value should be data type. expected: %v, actual: %v", e, a)
	}
	if a, e := data[3], "10, 2"; a != e {
		t.Errorf("Fourth value should be size. expected: %v, actual: %v", e, a)
	}
	if a, e := data[4], "NO"; a != e {
		t.Errorf("Fifth value should be nullable and value should be 'NO' when column is not nullable. expected: %v, actual: %v", e, a)
	}
	if a, e := data[5], "0.0"; a != e {
		t.Errorf("Sixth value should be default value. expected: %v, actual: %v", e, a)
	}
	if a, e := data[6], "item price"; a != e {
		t.Errorf("Seventh value should be comment. expected: %v, actual: %v", e, a)
	}
}

func TestDefaultConvertColumnOnPrimaryKey(t *testing.T) {
	conv := defaultConverter{}
	size := dbmodel.NewSize(
		sql.NullInt64{Valid: false},
		sql.NullInt64{Int64: 10, Valid: true},
		sql.NullInt64{Int64: 2, Valid: true},
	)
	col := dbmodel.NewColumn("foo", "items", "price", "item price", "numeric", size, false, "0.0", 1)
	data := conv.ConvertColumn(&col)
	if a, e := data[0], "1"; a != e {
		t.Errorf("First value should be primary key position when PrimaryKeyPosition is not 0. expected: %v, actual: %v", e, a)
	}
}

func TestDefaultConvertColumnOnNullable(t *testing.T) {
	conv := defaultConverter{}
	size := dbmodel.NewSize(
		sql.NullInt64{Valid: false},
		sql.NullInt64{Int64: 10, Valid: true},
		sql.NullInt64{Int64: 2, Valid: true},
	)
	col := dbmodel.NewColumn("foo", "items", "price", "item price", "numeric", size, true, "0.0", 0)
	data := conv.ConvertColumn(&col)
	if a, e := data[4], ""; a != e {
		t.Errorf("Fifth value should be empty when column is nullable. expected: %v, actual: %v", e, a)
	}
}

func TestDefaultConvertIndex(t *testing.T) {
	conv := defaultConverter{}
	idx := dbmodel.NewIndex("foo", "users", "users_pk", true)
	size := dbmodel.NewSize(
		sql.NullInt64{Valid: false},
		sql.NullInt64{Int64: 32, Valid: true},
		sql.NullInt64{Int64: 0, Valid: true},
	)
	col := dbmodel.NewColumn("foo", "users", "id", "", "int4", size, true, "", 1)
	idx.AddColumn(&col)
	data := conv.ConvertIndex(&idx)
	if a, e := data[0], "users_pk"; a != e {
		t.Errorf("First value should be index name. expected: %v, actual: %v", e, a)
	}
	if a, e := data[1], "id"; a != e {
		t.Errorf("Second value should be column name using in index. expected: %v, actual: %v", e, a)
	}
	if a, e := data[2], "YES"; a != e {
		t.Errorf("Third value should be unique and value should be 'YES' when index is unique index. expected: %v, actual: %v", e, a)
	}
}

func TestDefaultConvertIndexIsUnique(t *testing.T) {
	conv := defaultConverter{}
	idx := dbmodel.NewIndex("foo", "posts", "posts_user_id_idx", false)
	size := dbmodel.NewSize(
		sql.NullInt64{Valid: false},
		sql.NullInt64{Int64: 32, Valid: true},
		sql.NullInt64{Int64: 0, Valid: true},
	)
	col := dbmodel.NewColumn("foo", "posts", "user_id", "", "int4", size, true, "", 0)
	idx.AddColumn(&col)
	data := conv.ConvertIndex(&idx)
	if a, e := data[2], ""; a != e {
		t.Errorf("Unique should be empty when index is not unique index. expected: %v, actual: %v", e, a)
	}
}

func TestDefaultConvertIndexMultiColumns(t *testing.T) {
	conv := defaultConverter{}
	idx := dbmodel.NewIndex("foo", "sales", "sales_customer_id_product_id_idx", true)
	size := dbmodel.NewSize(
		sql.NullInt64{Valid: false},
		sql.NullInt64{Int64: 32, Valid: true},
		sql.NullInt64{Int64: 0, Valid: true},
	)
	col1 := dbmodel.NewColumn("foo", "sales", "customer_id", "", "int4", size, false, "", 0)
	idx.AddColumn(&col1)
	col2 := dbmodel.NewColumn("foo", "sales", "product_id", "", "int4", size, false, "", 0)
	idx.AddColumn(&col2)
	data := conv.ConvertIndex(&idx)
	if a, e := data[1], "customer_id, product_id"; a != e {
		t.Errorf("Columns should be enumerated names when index has multi columns. expected: %v, actual: %v", e, a)
	}
}

func TestDefaultConvertConstraint(t *testing.T) {
	conv := defaultConverter{}
	con := dbmodel.NewConstraint("foo", "users", "chk_user_age_over_zero", "CHECK", "(age > 0)")
	data := conv.ConvertConstraint(&con)
	if a, e := data[0], "chk_user_age_over_zero"; a != e {
		t.Errorf("First value should be constraint name. expected: %v, actual: %v", e, a)
	}
	if a, e := data[1], "CHECK"; a != e {
		t.Errorf("Second value should be constraint kind. expected: %v, actual: %v", e, a)
	}
	if a, e := data[2], "(age > 0)"; a != e {
		t.Errorf("Third value should be constraint content. expected: %v, actual: %v", e, a)
	}
}

func TestDefaultConvertForeignKey(t *testing.T) {
	conv := defaultConverter{}
	size := dbmodel.NewSize(
		sql.NullInt64{Valid: false},
		sql.NullInt64{Int64: 32, Valid: true},
		sql.NullInt64{Int64: 0, Valid: true},
	)
	fk := dbmodel.NewForeignKey("foo", "posts", "posts_user_id_fk")
	fCol := dbmodel.NewColumn("foo", "posts", "user_id", "", "int4", size, true, "", 0)
	tCol := dbmodel.NewColumn("foo", "users", "id", "", "int4", size, true, "", 1)
	ref := dbmodel.NewColumnReference(&fCol, &tCol)
	fk.AddColumnReference(&ref)
	data := conv.ConvertForeignKey(&fk)
	if a, e := data[0], "posts_user_id_fk"; a != e {
		t.Errorf("First value should be foreign key name. expected: %v, actual: %v", e, a)
	}
	if a, e := data[1], "user_id"; a != e {
		t.Errorf("Second value should be column name in table. expected: %v, actual: %v", e, a)
	}
	if a, e := data[2], "users"; a != e {
		t.Errorf("Third value should be foreign table name. expected: %v, actual: %v", e, a)
	}
	if a, e := data[3], "id"; a != e {
		t.Errorf("Fourth value should be foreign column name. expected: %v, actual: %v", e, a)
	}
}

func TestDefaultConvertForeignKeyOtherSchema(t *testing.T) {
	conv := defaultConverter{}
	size := dbmodel.NewSize(
		sql.NullInt64{Valid: false},
		sql.NullInt64{Int64: 32, Valid: true},
		sql.NullInt64{Int64: 0, Valid: true},
	)
	fk := dbmodel.NewForeignKey("foo", "posts", "posts_user_id_fk")
	fCol := dbmodel.NewColumn("foo", "posts", "user_id", "", "int4", size, true, "", 0)
	tCol := dbmodel.NewColumn("bar", "users", "id", "", "int4", size, true, "", 1)
	ref := dbmodel.NewColumnReference(&fCol, &tCol)
	fk.AddColumnReference(&ref)
	data := conv.ConvertForeignKey(&fk)
	if a, e := data[2], "bar.users"; a != e {
		t.Errorf("Foreign table name value should be with schema prefix when foreign table is in other schema. expected: %v, actual: %v", e, a)
	}
}

func TestDefaultConvertForeignKeyMultiColumns(t *testing.T) {
	conv := defaultConverter{}
	size := dbmodel.NewSize(
		sql.NullInt64{Valid: false},
		sql.NullInt64{Int64: 32, Valid: true},
		sql.NullInt64{Int64: 0, Valid: true},
	)
	fk := dbmodel.NewForeignKey("foo", "sales_archive", "sales_archive_customer_id_product_id_fk")
	fCol1 := dbmodel.NewColumn("foo", "sales_archive", "customer_id", "", "int4", size, true, "", 0)
	tCol1 := dbmodel.NewColumn("foo", "sales", "customer_id", "", "int4", size, true, "", 0)
	fCol2 := dbmodel.NewColumn("foo", "sales_archive", "product_id", "", "int4", size, true, "", 0)
	tCol2 := dbmodel.NewColumn("foo", "sales", "product_id", "", "int4", size, true, "", 0)
	ref1 := dbmodel.NewColumnReference(&fCol1, &tCol1)
	fk.AddColumnReference(&ref1)
	ref2 := dbmodel.NewColumnReference(&fCol2, &tCol2)
	fk.AddColumnReference(&ref2)
	data := conv.ConvertForeignKey(&fk)
	if a, e := data[1], "customer_id, product_id"; a != e {
		t.Errorf("Second value should be column name in table. expected: %v, actual: %v", e, a)
	}
	if a, e := data[3], "customer_id, product_id"; a != e {
		t.Errorf("Fourth value should be foreign column name. expected: %v, actual: %v", e, a)
	}
}

func TestDefaultConvertReferencedKey(t *testing.T) {
	conv := defaultConverter{}
	size := dbmodel.NewSize(
		sql.NullInt64{Valid: false},
		sql.NullInt64{Int64: 32, Valid: true},
		sql.NullInt64{Int64: 0, Valid: true},
	)
	fk := dbmodel.NewForeignKey("foo", "posts", "posts_user_id_fk")
	fCol := dbmodel.NewColumn("foo", "posts", "user_id", "", "int4", size, true, "", 0)
	tCol := dbmodel.NewColumn("foo", "users", "id", "", "int4", size, true, "", 1)
	ref := dbmodel.NewColumnReference(&fCol, &tCol)
	fk.AddColumnReference(&ref)
	data := conv.ConvertReferencedKey(&fk)
	if a, e := data[0], "posts_user_id_fk"; a != e {
		t.Errorf("First value should be foreign key name. expected: %v, actual: %v", e, a)
	}
	if a, e := data[1], "posts"; a != e {
		t.Errorf("Third value should be foreign table name. expected: %v, actual: %v", e, a)
	}
	if a, e := data[2], "user_id"; a != e {
		t.Errorf("Second value should be column name in table. expected: %v, actual: %v", e, a)
	}
	if a, e := data[3], "id"; a != e {
		t.Errorf("Fourth value should be foreign column name. expected: %v, actual: %v", e, a)
	}
}

func TestDefaultConvertReferencedKeyOtherSchema(t *testing.T) {
	conv := defaultConverter{}
	size := dbmodel.NewSize(
		sql.NullInt64{Valid: false},
		sql.NullInt64{Int64: 32, Valid: true},
		sql.NullInt64{Int64: 0, Valid: true},
	)
	fk := dbmodel.NewForeignKey("foo", "posts", "posts_user_id_fk")
	fCol := dbmodel.NewColumn("foo", "posts", "user_id", "", "int4", size, true, "", 0)
	tCol := dbmodel.NewColumn("bar", "users", "id", "", "int4", size, true, "", 1)
	ref := dbmodel.NewColumnReference(&fCol, &tCol)
	fk.AddColumnReference(&ref)
	data := conv.ConvertReferencedKey(&fk)
	if a, e := data[1], "foo.posts"; a != e {
		t.Errorf("Referenced table name value should be with schema prefix when foreign table is in other schema. expected: %v, actual: %v", e, a)
	}
}

func TestDefaultConvertReferencedKeyMultiColumns(t *testing.T) {
	conv := defaultConverter{}
	size := dbmodel.NewSize(
		sql.NullInt64{Valid: false},
		sql.NullInt64{Int64: 32, Valid: true},
		sql.NullInt64{Int64: 0, Valid: true},
	)
	fk := dbmodel.NewForeignKey("foo", "sales_archive", "sales_archive_customer_id_product_id_fk")
	fCol1 := dbmodel.NewColumn("foo", "sales_archive", "customer_id", "", "int4", size, true, "", 0)
	tCol1 := dbmodel.NewColumn("foo", "sales", "customer_id", "", "int4", size, true, "", 0)
	fCol2 := dbmodel.NewColumn("foo", "sales_archive", "product_id", "", "int4", size, true, "", 0)
	tCol2 := dbmodel.NewColumn("foo", "sales", "product_id", "", "int4", size, true, "", 0)
	ref1 := dbmodel.NewColumnReference(&fCol1, &tCol1)
	fk.AddColumnReference(&ref1)
	ref2 := dbmodel.NewColumnReference(&fCol2, &tCol2)
	fk.AddColumnReference(&ref2)
	data := conv.ConvertReferencedKey(&fk)
	if a, e := data[2], "customer_id, product_id"; a != e {
		t.Errorf("Second value should be column name in table. expected: %v, actual: %v", e, a)
	}
	if a, e := data[3], "customer_id, product_id"; a != e {
		t.Errorf("Fourth value should be foreign column name. expected: %v, actual: %v", e, a)
	}
}
