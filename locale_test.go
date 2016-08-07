package main

import "testing"

func TestL(t *testing.T) {
	loc := l("ja")
	if loc.t("column", "title") != "列一覧" {
		t.Error("Locale ja should be Japanese locale.")
	}
	loc = l("en")
	if loc.t("column", "title") != "Columns" {
		t.Error("Locale en should be English locale.")
	}
	loc = l("de")
	if loc.t("column", "title") != "Columns" {
		t.Error("Unknown locale should be English as default local.")
	}
}

func TestT(t *testing.T) {
	loc := l("ja")
	if loc.t("column", "name") != "列名" {
		t.Error("Invalid name is returned.")
	}
	if loc.t("foo", "name") != "" {
		t.Error("Translation function should return empty string when category is unknown.")
	}
	if loc.t("index", "foo") != "" {
		t.Error("Translation function should return empty string when key is unknown.")
	}
}
