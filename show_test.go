package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestCmdShow(t *testing.T) {
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
	buf := &bytes.Buffer{}
	o.out = buf
	setupTestConfigFile("tablarian-test")
	configFile = "test/tablarian-aw.config"
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
	buf := &bytes.Buffer{}
	o.out = buf
	setupTestConfigFile("tablarian-aw")
	absPath, err := testConfigFilePath()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failure config loading: %v", err)
	}
	configFile = "@" + absPath
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
