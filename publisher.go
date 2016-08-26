package main

import (
	"fmt"
	"io"

	"github.com/pinzolo/dbmodel"
)

// Publisher is interface for saving table definition.
type Publisher interface {
	Publish([]*dbmodel.Table)
	Errors() []error
}

func findPublisher(format string, config *Config, converter Converter, locale locale, logger io.Writer) (Publisher, error) {
	if format == "markdown" {
		return newMarkdownPublisher(config, converter, locale, logger), nil
	}

	return nil, fmt.Errorf("Format '%s' is invalid format.", format)
}
