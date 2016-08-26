package main

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadDefaultConfig(t *testing.T) {
	setupTestConfigFile("tablarian-test")
	cfg, err := loadConfig(DefaultConfigFileName)
	if err != nil {
		t.Errorf("Failure config loading: %v", err)
	}
	if !strings.HasSuffix(cfg.FilePath, DefaultConfigFileName) {
		t.Errorf("Load invalid config file. path: %s", cfg.FilePath)
	}
	if cfg.Driver != "postgres" {
		t.Errorf("Failure driver config loading. 'postgres' is expected but actual is %s", cfg.Driver)
	}
	if cfg.Version != "9.4" {
		t.Errorf("Failure version config loading. '9.4' is expected but actual is %s", cfg.Version)
	}
	if cfg.Host != "localhost" {
		t.Errorf("Failure host config loading. 'localhost' is expected but actual is %s", cfg.Host)
	}
	if cfg.Port != 5432 {
		t.Errorf("Failure port config loading. 5432 is expected but actual is %v", cfg.Port)
	}
	if cfg.User != "postgres" {
		t.Errorf("Failure user config loading. 'postgres' is expected but actual is %s", cfg.User)
	}
	if cfg.Password != "123456" {
		t.Errorf("Failure password config loading. '12345' is expected but actual is %s", cfg.Password)
	}
	if cfg.Database != "test" {
		t.Errorf("Failure database config loading. 'test' is expected but actual is %s", cfg.Database)
	}
	if cfg.Schema != "foo" {
		t.Errorf("Failure schema config loading. 'test' is expected but actual is %s", cfg.Schema)
	}
	if cfg.Options["sslmode"] != "disable" {
		t.Errorf("Failure options config loading. 'sslmode: disabled' is expected but actual is %#v", cfg.Options)
	}
	if cfg.Out != "out" {
		t.Errorf("Failure out config loading. 'out' is expected but actual is %s", cfg.Schema)
	}
}

func TestLoadOtherConfig(t *testing.T) {
	setupTestConfigFile("tablarian-other")
	cfg, err := loadConfig("test/tablarian-test.config")
	if err != nil {
		t.Errorf("Failure config loading: %v", err)
	}
	if !strings.HasSuffix(cfg.FilePath, "tablarian-test.config") {
		t.Errorf("Load invalid config file. path: %s", cfg.FilePath)
	}
	if cfg.Driver != "postgres" {
		t.Errorf("Failure driver config loading. 'postgres' is expected but actual is %s", cfg.Driver)
	}
	if cfg.Version != "9.4" {
		t.Errorf("Failure version config loading. '9.4' is expected but actual is %s", cfg.Version)
	}
	if cfg.Host != "localhost" {
		t.Errorf("Failure host config loading. 'localhost' is expected but actual is %s", cfg.Host)
	}
	if cfg.Port != 5432 {
		t.Errorf("Failure port config loading. 5432 is expected but actual is %v", cfg.Port)
	}
	if cfg.User != "postgres" {
		t.Errorf("Failure user config loading. 'postgres' is expected but actual is %s", cfg.User)
	}
	if cfg.Password != "123456" {
		t.Errorf("Failure password config loading. '12345' is expected but actual is %s", cfg.Password)
	}
	if cfg.Database != "test" {
		t.Errorf("Failure database config loading. 'test' is expected but actual is %s", cfg.Database)
	}
	if cfg.Schema != "foo" {
		t.Errorf("Failure schema config loading. 'test' is expected but actual is %s", cfg.Schema)
	}
	if cfg.Options["sslmode"] != "disable" {
		t.Errorf("Failure options config loading. 'sslmode: disabled' is expected but actual is %#v", cfg.Options)
	}
	if cfg.Out != "out" {
		t.Errorf("Failure out config loading. 'out' is expected but actual is %s", cfg.Schema)
	}
}

func TestLoadConfigWithAbsPath(t *testing.T) {
	setupTestConfigFile("tablarian-test")
	absPath, err := testConfigFilePath()
	cfg, err := loadConfig("@" + absPath)
	if err != nil {
		t.Errorf("Failure config loading: %v", err)
	}
	if !strings.HasSuffix(cfg.FilePath, DefaultConfigFileName) {
		t.Errorf("Load invalid config file. path: %s", cfg.FilePath)
	}
	if cfg.Driver != "postgres" {
		t.Errorf("Failure driver config loading. 'postgres' is expected but actual is %s", cfg.Driver)
	}
	if cfg.Version != "9.4" {
		t.Errorf("Failure version config loading. '9.4' is expected but actual is %s", cfg.Version)
	}
	if cfg.Host != "localhost" {
		t.Errorf("Failure host config loading. 'localhost' is expected but actual is %s", cfg.Host)
	}
	if cfg.Port != 5432 {
		t.Errorf("Failure port config loading. 5432 is expected but actual is %v", cfg.Port)
	}
	if cfg.User != "postgres" {
		t.Errorf("Failure user config loading. 'postgres' is expected but actual is %s", cfg.User)
	}
	if cfg.Password != "123456" {
		t.Errorf("Failure password config loading. '12345' is expected but actual is %s", cfg.Password)
	}
	if cfg.Database != "test" {
		t.Errorf("Failure database config loading. 'test' is expected but actual is %s", cfg.Database)
	}
	if cfg.Schema != "foo" {
		t.Errorf("Failure schema config loading. 'test' is expected but actual is %s", cfg.Schema)
	}
	if cfg.Options["sslmode"] != "disable" {
		t.Errorf("Failure options config loading. 'sslmode: disabled' is expected but actual is %#v", cfg.Options)
	}
	if cfg.Out != "out" {
		t.Errorf("Failure out config loading. 'out' is expected but actual is %s", cfg.Schema)
	}
}

func TestLoadConfigWithInvalidJson(t *testing.T) {
	setupTestConfigFile("invalid-json")
	_, err := loadConfig(DefaultConfigFileName)
	if err == nil {
		t.Error("Config loading should fail on invalid json.")
	}
}

func TestLoadConfigWithConfigFileNotFound(t *testing.T) {
	deleteTestConfigFile()
	_, err := loadConfig(DefaultConfigFileName)
	if err == nil {
		t.Error("Config loading should fail on config file not exists.")
	}
}

func setupTestConfigFile(fileName string) error {
	deleteTestConfigFile()
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	src, err := os.Open(filepath.Join(wd, "test", fileName+".config"))
	if err != nil {
		return err
	}
	defer src.Close()

	dest, err := os.Create(filepath.Join(wd, DefaultConfigFileName))
	if err != nil {
		return err
	}

	_, err = io.Copy(dest, src)
	if err != nil {
		return err
	}
	return nil
}

func deleteTestConfigFile() {
	path, perr := testConfigFilePath()
	_, serr := os.Stat(path)
	if perr == nil && serr == nil {
		os.Remove(path)
	}
}

func testConfigFilePath() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(wd, DefaultConfigFileName), nil
}
