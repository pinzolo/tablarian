package main

type locale struct {
	dict map[string]map[string]string
}

var (
	en = locale{
		dict: map[string]map[string]string{
			"table_list": map[string]string{
				"title":   "Table index",
				"table":   "TABLE",
				"comment": "COMMENT",
			},
			"column": map[string]string{
				"title":         "Columns",
				"primary_key":   "PK",
				"name":          "NAME",
				"data_type":     "TYPE",
				"size":          "SIZE",
				"null":          "NULL",
				"default_value": "DEFAULT",
				"comment":       "COMMENT",
			},
			"index": map[string]string{
				"title":   "Indices",
				"name":    "NAME",
				"columns": "COLUMNS",
				"unique":  "UNIQUE",
			},
			"constraint": map[string]string{
				"title":   "Constraints",
				"name":    "NAME",
				"kind":    "KIND",
				"content": "CONTENT",
			},
			"foreign_key": map[string]string{
				"title":           "Foreign keys",
				"name":            "NAME",
				"columns":         "COLUMNS",
				"foreign_table":   "FOREIGN TABLE",
				"foreign_columns": "FOREIGN COLUMNS",
			},
			"referenced_key": map[string]string{
				"title":          "Referenced keys",
				"name":           "NAME",
				"source_table":   "SOURCE TABLE",
				"source_columns": "SOURCE COLUMNS",
				"columns":        "COLUMNS",
			},
		},
	}
	ja = locale{
		dict: map[string]map[string]string{
			"table_list": map[string]string{
				"title":   "テーブル一覧",
				"table":   "テーブル",
				"comment": "コメント",
			},
			"column": map[string]string{
				"title":         "列一覧",
				"primary_key":   "PK",
				"name":          "列名",
				"data_type":     "型",
				"size":          "サイズ",
				"null":          "NULL",
				"default_value": "初期値",
				"comment":       "コメント",
			},
			"index": map[string]string{
				"title":   "インデックス",
				"name":    "名前",
				"columns": "列",
				"unique":  "ユニーク",
			},
			"constraint": map[string]string{
				"title":   "制約",
				"name":    "製薬名",
				"kind":    "KIND",
				"content": "CONTENT",
			},
			"foreign_key": map[string]string{
				"title":           "参照キー",
				"name":            "参照名",
				"columns":         "列",
				"foreign_table":   "参照テーブル",
				"foreign_columns": "参照列",
			},
			"referenced_key": map[string]string{
				"title":          "被参照キー",
				"name":           "参照名",
				"source_table":   "参照元テーブル",
				"source_columns": "参照元列",
				"columns":        "被参照列",
			},
		},
	}
)

func (loc locale) t(cat string, key string) string {
	m, ok := loc.dict[cat]
	if !ok {
		return ""
	}
	return m[key]
}

func l(locale string) locale {
	if locale == "ja" {
		return ja
	}
	return en
}
