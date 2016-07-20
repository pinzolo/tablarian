package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestCmdIndex(t *testing.T) {
	buf := &bytes.Buffer{}
	o.out = buf
	setupTestConfigFile("tablarian-aw")
	cmdIndex.Run([]string{})
	expected := `
country_region_currency
currency`
	actual := buf.String()
	if strings.TrimSpace(expected) != strings.TrimSpace(actual) {
		t.Errorf("\nactual:\n%v\nexpected:%v\n", actual, expected)
	}
}

func TestCmdIndexWithOtherConfig(t *testing.T) {
	buf := &bytes.Buffer{}
	o.out = buf
	setupTestConfigFile("tablarian-test")
	configFile = "test/tablarian-aw.config"
	cmdIndex.Run([]string{})
	expected := `
country_region_currency
currency`
	actual := buf.String()
	if strings.TrimSpace(expected) != strings.TrimSpace(actual) {
		t.Errorf("\nactual:\n%v\nexpected:%v\n", actual, expected)
	}
}

func TestCmdIndexWithOtherConfigByAbsPath(t *testing.T) {
	buf := &bytes.Buffer{}
	o.out = buf
	setupTestConfigFile("tablarian-aw")
	absPath, err := testConfigFilePath()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failure config loading: %v", err)
	}
	configFile = "@" + absPath
	cmdIndex.Run([]string{})
	expected := `
country_region_currency
currency`
	actual := buf.String()
	if strings.TrimSpace(expected) != strings.TrimSpace(actual) {
		t.Errorf("\nactual:\n%v\nexpected:%v\n", actual, expected)
	}
}
