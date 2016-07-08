package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Config stores loaded config file content
type Config struct {
	FilePath string            `json:"-"`
	Driver   string            `json:"driver"`
	Host     string            `json:"host"`
	Port     int               `json:"port"`
	User     string            `json:"user"`
	Password string            `json:"password"`
	Database string            `json:"database"`
	Schema   string            `json:"schema"`
	Options  map[string]string `json:"options"`
}

// LoadConfig returns loaded config file content as *Config
func LoadConfig(path string) (*Config, error) {
	rpath, err := resolvePath(path)
	if err != nil {
		return nil, err
	}

	cfg := &Config{FilePath: rpath}
	content, err := ioutil.ReadFile(rpath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(content, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func resolvePath(path string) (string, error) {
	if strings.HasPrefix(path, "@") {
		return strings.TrimPrefix(path, "@"), nil
	}

	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(wd, path), nil
}
