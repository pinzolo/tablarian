package main

import (
	"fmt"
	"os"
)

func Example_cmdIndex() {
	setupTestConfigFile("tablarian-aw")
	cmdIndex.Run([]string{})
	// Output:
	// country_region_currency
	// currency
}

func Example_cmdIndex_with_other_config() {
	setupTestConfigFile("tablarian-test")
	configFile = "test/tablarian-aw.config"
	cmdIndex.Run([]string{})
	// Output:
	// country_region_currency
	// currency
}

func Example_cmdIndex_with_other_config_on_absolute_file_path() {
	setupTestConfigFile("tablarian-aw")
	absPath, err := testConfigFilePath()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failure config loading: %v", err)
	}
	configFile = "@" + absPath
	cmdIndex.Run([]string{})
	// Output:
	// country_region_currency
	// currency
}
