package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var cmdInit = &Command{
	Run:       runInit,
	UsageLine: "version ",
	Short:     "Create config file template.",
	Long: `Create config template as 'tablarian.config' to working directory.
Argument is your database sysytem driver name. (current acceptable 'postgres' only.)
	`,
}

func init() {
	// Set your flag here like below.
	// cmdVersion.Flag.BoolVar(&flagA, "a", false, "")
}

// runInit executes version command and return exit code.
func runInit(args []string) int {
	if len(args) == 0 {
		fmt.Fprintln(o.err, "Database system driver name is required.")
		return 1
	}

	cfg, err := findConfigContent(args[0])
	if err != nil {
		fmt.Fprintln(o.err, err)
		return 1
	}

	err = writeConfigContent(cfg)
	if err != nil {
		fmt.Fprintln(o.err, err)
		return 1
	}

	fmt.Println("Config file template is created as 'tablarian.config'.")
	return 0
}

func findConfigContent(d string) (string, error) {
	if d == "postgres" {
		return postgresConfigContent(), nil
	}
	return "", fmt.Errorf("Driver name '%s' is unknow.", d)
}

func postgresConfigContent() string {
	return `{
  "driver": "postgres",
  "version": "9.4",
  "host": "localhost",
  "port": 5432,
  "user": "postgres",
  "password": "your-password",
  "database": "postgres",
  "schema": "public",
  "options": {
    "sslmode": "disable"
  },
  "out" : "out"
}`
}

func writeConfigContent(cfg string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	path := filepath.Join(wd, "tablarian.config")
	if _, err = os.Stat(path); err == nil {
		return errors.New("Config file 'tablarian.config' already exists.")
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	f.Write([]byte(cfg))
	return nil
}
