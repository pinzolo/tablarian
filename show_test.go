package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestCmdShow(t *testing.T) {
	initShowOpt()
	buf := &bytes.Buffer{}
	o.out = buf
	setupTestConfigFile("tablarian-aw")
	args := []string{"currency"}
	cmdShow.Run(args)
	expected := strings.TrimSpace(`
+----+---------------+-------------+------+------+---------+--------------------------------+
| PK |     NAME      |    TYPE     | SIZE | NULL | DEFAULT |            COMMENT             |
+----+---------------+-------------+------+------+---------+--------------------------------+
|  1 | currency_code | bpchar      |    3 | NO   |         | The ISO code for the Currency. |
|    | name          | public.Name |   50 | NO   |         | Currency name.                 |
|    | modified_date | timestamp   |    6 | NO   | now()   |                                |
+----+---------------+-------------+------+------+---------+--------------------------------+`)
	if actual := strings.TrimSpace(buf.String()); expected != actual {
		t.Errorf("\nactual:\n%v\nexpected:\n%v\n", actual, expected)
	}
}

func TestCmdShowWithOtherConfig(t *testing.T) {
	initShowOpt()
	buf := &bytes.Buffer{}
	o.out = buf
	setupTestConfigFile("tablarian-test")
	showOpt.configFile = "test/tablarian-aw.config"
	args := []string{"currency"}
	cmdShow.Run(args)
	expected := strings.TrimSpace(`
+----+---------------+-------------+------+------+---------+--------------------------------+
| PK |     NAME      |    TYPE     | SIZE | NULL | DEFAULT |            COMMENT             |
+----+---------------+-------------+------+------+---------+--------------------------------+
|  1 | currency_code | bpchar      |    3 | NO   |         | The ISO code for the Currency. |
|    | name          | public.Name |   50 | NO   |         | Currency name.                 |
|    | modified_date | timestamp   |    6 | NO   | now()   |                                |
+----+---------------+-------------+------+------+---------+--------------------------------+`)
	if actual := strings.TrimSpace(buf.String()); expected != actual {
		t.Errorf("\nactual:\n%v\nexpected:\n%v\n", actual, expected)
	}
}

func TestCmdShowWithOtherConfigByAbsPath(t *testing.T) {
	initShowOpt()
	buf := &bytes.Buffer{}
	o.out = buf
	setupTestConfigFile("tablarian-aw")
	absPath, err := testConfigFilePath()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failure config loading: %v", err)
	}
	showOpt.configFile = "@" + absPath
	args := []string{"currency"}
	cmdShow.Run(args)
	expected := strings.TrimSpace(`
+----+---------------+-------------+------+------+---------+--------------------------------+
| PK |     NAME      |    TYPE     | SIZE | NULL | DEFAULT |            COMMENT             |
+----+---------------+-------------+------+------+---------+--------------------------------+
|  1 | currency_code | bpchar      |    3 | NO   |         | The ISO code for the Currency. |
|    | name          | public.Name |   50 | NO   |         | Currency name.                 |
|    | modified_date | timestamp   |    6 | NO   | now()   |                                |
+----+---------------+-------------+------+------+---------+--------------------------------+`)
	if actual := strings.TrimSpace(buf.String()); expected != actual {
		t.Errorf("\nactual:\n%v\nexpected:\n%v\n", actual, expected)
	}
}

func TestCmdShowWithAllOption(t *testing.T) {
	initShowOpt()
	buf := &bytes.Buffer{}
	o.out = buf
	setupTestConfigFile("tablarian-aw")
	showOpt.showAll = true
	args := []string{"sales_person"}
	cmdShow.Run(args)
	expected := strings.TrimSpace(`
+----+--------------------+-----------+-------+------+--------------------+-------------------------------------------------------------------------------------+
| PK |        NAME        |   TYPE    | SIZE  | NULL |      DEFAULT       |                                       COMMENT                                       |
+----+--------------------+-----------+-------+------+--------------------+-------------------------------------------------------------------------------------+
|  1 | business_entity_id | int4      | 32, 0 | NO   |                    | Primary key for SalesPerson records. Foreign key to employee.business_entity_id     |
|    | territory_id       | int4      | 32, 0 |      |                    | Territory currently assigned to. Foreign key to sales_territory.sales_territory_id. |
|    | sales_quota        | numeric   |       |      |                    | Projected yearly sales.                                                             |
|    | bonus              | numeric   |       | NO   |               0.00 | Bonus due if quota is met.                                                          |
|    | commission_pct     | numeric   |       | NO   |               0.00 | Commision percent received per sale.                                                |
|    | sales_ytd          | numeric   |       | NO   |               0.00 | Sales total year to date.                                                           |
|    | sales_last_year    | numeric   |       | NO   |               0.00 | Sales total of previous year.                                                       |
|    | rowguid            | uuid      |       | NO   | uuid_generate_v1() |                                                                                     |
|    | modified_date      | timestamp |     6 | NO   | now()              |                                                                                     |
+----+--------------------+-----------+-------+------+--------------------+-------------------------------------------------------------------------------------+

### Indices
+------------------------------------+--------------------+--------+
|                NAME                |      COLUMNS       | UNIQUE |
+------------------------------------+--------------------+--------+
| pk_sales_person_business_entity_id | business_entity_id | YES    |
+------------------------------------+--------------------+--------+

### Constraints
+---------------------------------+-------+---------------------------+
|              NAME               | KIND  |          CONTENT          |
+---------------------------------+-------+---------------------------+
| ck_sales_person_bonus           | CHECK | (bonus >= 0.00)           |
| ck_sales_person_commission_pct  | CHECK | (commission_pct >= 0.00)  |
| ck_sales_person_sales_last_year | CHECK | (sales_last_year >= 0.00) |
| ck_sales_person_sales_quota     | CHECK | (sales_quota > 0.00)      |
| ck_sales_person_sales_ytd       | CHECK | (sales_ytd >= 0.00)       |
+---------------------------------+-------+---------------------------+

### Foreign keys
+----------------------------------------------+--------------------+--------------------------+--------------------+
|                     NAME                     |      COLUMNS       |      FOREIGN TABLE       |  FOREIGN COLUMNS   |
+----------------------------------------------+--------------------+--------------------------+--------------------+
| fk_sales_person_employee_business_entity_id  | business_entity_id | human_resources.employee | business_entity_id |
| fk_sales_person_sales_territory_territory_id | territory_id       | sales_territory          | territory_id       |
+----------------------------------------------+--------------------+--------------------------+--------------------+

### Referenced keys
+---------------------------------------------------------------+----------------------------+--------------------+--------------------+
|                             NAME                              |        SOURCE TABLE        |   SOURCE COLUMNS   |      COLUMNS       |
+---------------------------------------------------------------+----------------------------+--------------------+--------------------+
| fk_sales_order_header_sales_person_sales_person_id            | sales_order_header         | sales_person_id    | business_entity_id |
| fk_sales_person_quota_history_sales_person_business_entity_id | sales_person_quota_history | business_entity_id | business_entity_id |
| fk_sales_territory_history_sales_person_business_entity_id    | sales_territory_history    | business_entity_id | business_entity_id |
| fk_store_sales_person_sales_person_id                         | store                      | sales_person_id    | business_entity_id |
+---------------------------------------------------------------+----------------------------+--------------------+--------------------+`)
	if actual := strings.TrimSpace(buf.String()); expected != actual {
		t.Errorf("\nactual:\n%v\nexpected:\n%v\n", actual, expected)
	}
}

func TestCmdShowWithInvalidJson(t *testing.T) {
	initShowOpt()
	buf := &bytes.Buffer{}
	o.err = buf
	setupTestConfigFile("invalid-json")
	args := []string{"sales_person"}
	stat := cmdShow.Run(args)
	if stat == 0 {
		t.Error("Show command should not finish normally on invalid json.")
	}
	if actual, expected := strings.TrimSpace(buf.String()), "unexpected end of JSON input"; actual != expected {
		t.Errorf("Error masseage is not expected. actual: %v, expected: %v", actual, expected)
	}
}

func TestCmdShowWithDbError(t *testing.T) {
	initShowOpt()
	buf := &bytes.Buffer{}
	o.err = buf
	setupTestConfigFile("db-error")
	args := []string{"sales_person"}
	stat := cmdShow.Run(args)
	if stat == 0 {
		t.Error("Show command should not finish normally on invalid schema.")
	}
	if actual, expected := strings.TrimSpace(buf.String()), `pq: role "foobar" does not exist`; actual != expected {
		t.Errorf("Error masseage is not expected. actual: %v, expected: %v", actual, expected)
	}
}

func TestCmdShowWithoutTableName(t *testing.T) {
	initShowOpt()
	buf := &bytes.Buffer{}
	o.err = buf
	setupTestConfigFile("tablarian-aw")
	args := []string{}
	stat := cmdShow.Run(args)
	if stat == 0 {
		t.Error("Show command should not finish normally without table name argument.")
	}
	if actual, expected := strings.TrimSpace(buf.String()), "require table name as argument."; actual != expected {
		t.Errorf("Error masseage is not expected. actual: %v, expected: %v", actual, expected)
	}
}

func TestCmdShowWithPrettyOption(t *testing.T) {
	initShowOpt()
	buf := &bytes.Buffer{}
	o.out = buf
	setupTestConfigFile("tablarian-aw")
	showOpt.prettyPrint = true
	args := []string{"customer"}
	cmdShow.Run(args)
	expected := strings.TrimSpace(`
+----+---------------+-----------+------+------+--------------------+----------------------------------------------------------------------------------------------------------+
| PK |     NAME      |   TYPE    | SIZE | NULL |      DEFAULT       |                                                 COMMENT                                                  |
+----+---------------+-----------+------+------+--------------------+----------------------------------------------------------------------------------------------------------+
|  1 | customer_id   | serial    |      | NO   |                    | Primary key.                                                                                             |
|    | person_id     | integer   |      |      |                    | Foreign key to person.business_entity_id                                                                 |
|    | store_id      | integer   |      |      |                    | Foreign key to Store.business_entity_id                                                                  |
|    | territory_id  | integer   |      |      |                    | ID of the territory in which the customer is located. Foreign key to sales_territory.sales_territory_id. |
|    | rowguid       | uuid      |      | NO   | uuid_generate_v1() |                                                                                                          |
|    | modified_date | timestamp |    6 | NO   | now()              |                                                                                                          |
+----+---------------+-----------+------+------+--------------------+----------------------------------------------------------------------------------------------------------+`)
	if actual := strings.TrimSpace(buf.String()); expected != actual {
		t.Errorf("\nactual:\n%v\nexpected:\n%v\n", actual, expected)
	}
}

func initShowOpt() {
	showOpt.configFile = "tablarian.config"
	showOpt.showAll = false
	showOpt.prettyPrint = false
}
