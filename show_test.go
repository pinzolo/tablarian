package main

import (
	"fmt"
	"os"
)

func Example_cmdShow() {
	setupTestConfigFile("tablarian-aw")
	args := []string{"currency"}
	cmdShow.Run(args)
	// Output:
	// +----+---------------+-------------+------+------+---------+--------------------------------+
	// | PK |     NAME      |    TYPE     | SIZE | NULL | DEFAULT |            COMMENT             |
	// +----+---------------+-------------+------+------+---------+--------------------------------+
	// |  1 | currency_code | bpchar      |    3 | NO   |         | The ISO code for the currency. |
	// |    | name          | public.Name |   50 | NO   |         |                                |
	// |    | memo          | text        |      |      |         |                                |
	// |    | modified_date | timestamp   |    6 | NO   | now()   |                                |
	// +----+---------------+-------------+------+------+---------+--------------------------------+
}

func Example_cmdShow_with_other_config() {
	setupTestConfigFile("tablarian-test")
	configFile = "test/tablarian-aw.config"
	args := []string{"currency"}
	cmdShow.Run(args)
	// Output:
	// +----+---------------+-------------+------+------+---------+--------------------------------+
	// | PK |     NAME      |    TYPE     | SIZE | NULL | DEFAULT |            COMMENT             |
	// +----+---------------+-------------+------+------+---------+--------------------------------+
	// |  1 | currency_code | bpchar      |    3 | NO   |         | The ISO code for the currency. |
	// |    | name          | public.Name |   50 | NO   |         |                                |
	// |    | memo          | text        |      |      |         |                                |
	// |    | modified_date | timestamp   |    6 | NO   | now()   |                                |
	// +----+---------------+-------------+------+------+---------+--------------------------------+
}

func Example_cmdShow_with_other_config_on_absolute_file_path() {
	setupTestConfigFile("tablarian-aw")
	absPath, err := testConfigFilePath()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failure config loading: %v", err)
	}
	configFile = "@" + absPath
	args := []string{"currency"}
	cmdShow.Run(args)
	// Output:
	// +----+---------------+-------------+------+------+---------+--------------------------------+
	// | PK |     NAME      |    TYPE     | SIZE | NULL | DEFAULT |            COMMENT             |
	// +----+---------------+-------------+------+------+---------+--------------------------------+
	// |  1 | currency_code | bpchar      |    3 | NO   |         | The ISO code for the currency. |
	// |    | name          | public.Name |   50 | NO   |         |                                |
	// |    | memo          | text        |      |      |         |                                |
	// |    | modified_date | timestamp   |    6 | NO   | now()   |                                |
	// +----+---------------+-------------+------+------+---------+--------------------------------+
}
