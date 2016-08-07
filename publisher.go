package main

import (
	"fmt"

	"github.com/pinzolo/dbmodel"
)

// Publisher is interface for saving table definition.
type Publisher interface {
	Publish([]*dbmodel.Table)
	Errors() []error
}

func findPublisher(format string, config *Config, converter Converter, locale locale) (Publisher, error) {
	if format == "markdown" {
		return newMarkdownPublisher(config, converter, locale), nil
	}

	return nil, fmt.Errorf("Format '%s' is invalid format.", format)
}
