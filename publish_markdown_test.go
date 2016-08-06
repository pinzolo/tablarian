package main

import (
	"bytes"
	"crypto/md5"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var salesTblNames = []string{
	"country_region_currency",
	"credit_card",
	"currency",
	"currency_rate",
	"customer",
	"person_credit_card",
	"sales_order_detail",
	"sales_order_header",
	"sales_order_header_sales_reason",
	"sales_person",
	"sales_person_quota_history",
	"sales_reason",
	"sales_tax_rate",
	"sales_territory",
	"sales_territory_history",
	"shopping_cart_item",
	"special_offer",
	"special_offer_product",
	"store",
}

func TestCmdPublishDefault(t *testing.T) {
	initPublishOpt()
	buf := &bytes.Buffer{}
	o.out = buf
	setupTestConfigFile("tablarian-aw")
	stat := cmdPublish.Run([]string{})
	if stat != 0 {
		t.Error("Publish subcommand should finish normally.")
	}

	path, err := resolvePath("out")
	if err != nil {
		t.Error("Output path should be able to resolve.")
	}
	if fi, err := os.Stat(path); err != nil || !fi.IsDir() {
		t.Error("Publish subcommand should make output directory.")
	}

	idxPath := filepath.Join(path, "00_index.md")
	_, err = os.Stat(idxPath)
	if err != nil {
		t.Error("Default publish subcommand should make index file.")
	}

	if !isSameFile(idxPath, "sales_00_index.md") {
		t.Error("Index file content is not expected.")
	}

	for _, n := range salesTblNames {
		_, err = os.Stat(filepath.Join(path, n+".md"))
		if err != nil {
			t.Errorf("Default publish subcommand should make %s.", n+".md")
		}
	}

	if !isSameFile(filepath.Join(path, "sales_order_header.md"), "default_sales_order_header.md") {
		t.Errorf("File: %v content is not expected.", "sales_order_header.md")
	}
}

func TestCmdPublishMarkdownPostgresPretty(t *testing.T) {
	initPublishOpt()
	buf := &bytes.Buffer{}
	o.out = buf
	publishOpt.prettyPrint = true
	setupTestConfigFile("tablarian-aw")
	stat := cmdPublish.Run([]string{})
	if stat != 0 {
		t.Error("Publish subcommand should finish normally.")
	}

	path, err := resolvePath("out")
	if err != nil {
		t.Error("Output path should be able to resolve.")
	}
	if fi, err := os.Stat(path); err != nil || !fi.IsDir() {
		t.Error("Publish subcommand should make output directory.")
	}

	idxPath := filepath.Join(path, "00_index.md")
	_, err = os.Stat(idxPath)
	if err != nil {
		t.Error("Default publish subcommand should make index file.")
	}

	if !isSameFile(idxPath, "sales_00_index.md") {
		t.Error("Index file content is not expected.")
	}

	for _, n := range salesTblNames {
		_, err = os.Stat(filepath.Join(path, n+".md"))
		if err != nil {
			t.Errorf("Default publish subcommand should make %s.", n+".md")
		}
	}

	if !isSameFile(filepath.Join(path, "sales_order_header.md"), "default_sales_order_header_pretty.md") {
		t.Errorf("File: %v content is not expected.", "sales_order_header.md")
	}
}

func TestCmdPublishMarkdownWhenOutDirNotExists(t *testing.T) {
	initPublishOpt()
	buf := &bytes.Buffer{}
	o.out = buf
	setupTestConfigFile("tablarian-aw")
	path, err := resolvePath("out")
	if err != nil {
		t.Error("Output path should be able to resolve.")
		return
	}

	if fi, err := os.Stat(path); err == nil && fi.IsDir() {
		err = os.RemoveAll(path)
		if err != nil {
			t.Error("Output directory cannot remove.")
			return
		}
	}

	stat := cmdPublish.Run([]string{})
	if stat != 0 {
		t.Error("Publish subcommand should finish normally.")
	}

	if fi, err := os.Stat(path); err != nil || !fi.IsDir() {
		t.Error("Publish subcommand should make output directory.")
	}

	idxPath := filepath.Join(path, "00_index.md")
	_, err = os.Stat(idxPath)
	if err != nil {
		t.Error("Default publish subcommand should make index file.")
	}

	for _, n := range salesTblNames {
		_, err = os.Stat(filepath.Join(path, n+".md"))
		if err != nil {
			t.Errorf("Default publish subcommand should make %s.", n+".md")
		}
	}
}

func TestCmdPublishMarkdownOutDirCleaning(t *testing.T) {
	initPublishOpt()
	buf := &bytes.Buffer{}
	o.out = buf
	setupTestConfigFile("tablarian-aw")
	path, err := resolvePath("out")
	if err != nil {
		t.Error("Output path should be able to resolve.")
		return
	}
	_, err = os.Create(filepath.Join(path, "unconcerned.md"))
	if err != nil {
		t.Error("Unconcernd file should be maid.")
		return
	}

	if fi, err := os.Stat(path); err == nil && fi.IsDir() {
		err = os.RemoveAll(path)
		if err != nil {
			t.Error("Output directory cannot remove.")
			return
		}
	}

	stat := cmdPublish.Run([]string{})
	if stat != 0 {
		t.Error("Publish subcommand should finish normally.")
	}

	if fi, err := os.Stat(path); err != nil || !fi.IsDir() {
		t.Error("Publish subcommand should make output directory.")
	}

	if _, err := os.Stat(filepath.Join(path, "unconcernd.md")); err == nil {
		t.Error("Unconcernd file should be deleted.")
	}

	idxPath := filepath.Join(path, "00_index.md")
	_, err = os.Stat(idxPath)
	if err != nil {
		t.Error("Default publish subcommand should make index file.")
	}

	for _, n := range salesTblNames {
		_, err = os.Stat(filepath.Join(path, n+".md"))
		if err != nil {
			t.Errorf("Default publish subcommand should make %s.", n+".md")
		}
	}
}

func TestCmdPublishWithInvalidJson(t *testing.T) {
	initPublishOpt()
	buf := &bytes.Buffer{}
	o.err = buf
	setupTestConfigFile("invalid-json")
	stat := cmdPublish.Run([]string{})
	if stat == 0 {
		t.Error("Publish command should not finish normally on invalid json.")
	}
	if actual, expected := strings.TrimSpace(buf.String()), "unexpected end of JSON input"; actual != expected {
		t.Errorf("Error masseage is not expected. actual: %v, expected: %v", actual, expected)
	}
}

func TestCmdPublishWithDbError(t *testing.T) {
	initPublishOpt()
	buf := &bytes.Buffer{}
	o.err = buf
	setupTestConfigFile("db-error")
	stat := cmdPublish.Run([]string{})
	if stat == 0 {
		t.Error("Publish command should not finish normally on invalid schema.")
	}
	if actual, expected := strings.TrimSpace(buf.String()), `pq: role "foobar" does not exist`; actual != expected {
		t.Errorf("Error masseage is not expected. actual: %v, expected: %v", actual, expected)
	}
}

func TestCmdPublishWithInvalidFormat(t *testing.T) {
	initPublishOpt()
	buf := &bytes.Buffer{}
	o.err = buf
	setupTestConfigFile("tablarian-aw")
	publishOpt.format = "invalid"
	stat := cmdPublish.Run([]string{})
	if stat == 0 {
		t.Error("Publish command should not finish normally on invalid format.")
	}
	if actual, expected := strings.TrimSpace(buf.String()), "Format 'invalid' is invalid format."; actual != expected {
		t.Errorf("Error masseage is not expected. actual: %v, expected: %v", actual, expected)
	}
}

func initPublishOpt() {
	publishOpt.configFile = "tablarian.config"
	publishOpt.prettyPrint = false
	publishOpt.format = "markdown"
}

func isSameFile(path string, testFile string) bool {
	cs1, err := ioutil.ReadFile(path)
	if err != nil {
		return false
	}
	wd, err := os.Getwd()
	if err != nil {
		return false
	}
	cs2, err := ioutil.ReadFile(filepath.Join(wd, "test", testFile))
	if err != nil {
		return false
	}
	return md5.Sum(cs1) == md5.Sum(cs2)
}
