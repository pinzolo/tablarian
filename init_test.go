package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestCmdInitWithPostgres(t *testing.T) {
	deleteTestConfigFile()
	stat := cmdInit.Run([]string{"postgres"})
	if stat != 0 {
		t.Error("Init subcommand should finish normally.")
	}

	path, err := testConfigFilePath()
	if err != nil {
		t.Error(err)
	}

	f, err := os.Open(path)
	if err != nil {
		t.Error(err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Error(err)
	}

	if string(b) != postgresConfigContent() {
		t.Error("Init subcommand wrote invalid content.")
	}
}

func TestCmdInitWithoutArgument(t *testing.T) {
	buf := &bytes.Buffer{}
	o.err = buf
	stat := cmdInit.Run([]string{})
	if stat == 0 {
		t.Error("Init command should not finish normally without argument.")
	}
	if actual, expected := strings.TrimSpace(buf.String()), "Database system driver name is required."; actual != expected {
		t.Errorf("Error masseage is not expected. actual: %v, expected: %v", actual, expected)
	}
}

func TestCmdInitWithInvalidDriverName(t *testing.T) {
	buf := &bytes.Buffer{}
	o.err = buf
	stat := cmdInit.Run([]string{"foobar"})
	if stat == 0 {
		t.Error("Init command should not finish normally with invalid driver name.")
	}
	if actual, expected := strings.TrimSpace(buf.String()), "Driver name 'foobar' is unknow."; actual != expected {
		t.Errorf("Error masseage is not expected. actual: %v, expected: %v", actual, expected)
	}
}

func TestCmdInitOnConfigFileAlreadyExists(t *testing.T) {
	buf := &bytes.Buffer{}
	o.err = buf
	setupTestConfigFile("tablarian-aw")
	stat := cmdInit.Run([]string{"postgres"})
	if stat == 0 {
		t.Error("Init command should not finish normally when config file already exists.")
	}
	if actual, expected := strings.TrimSpace(buf.String()), "Config file 'tablarian.config' already exists."; actual != expected {
		t.Errorf("Error masseage is not expected. actual: %v, expected: %v", actual, expected)
	}
}
