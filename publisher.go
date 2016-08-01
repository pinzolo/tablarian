package main

import "github.com/pinzolo/dbmodel"

type Publisher interface {
	Publish([]*dbmodel.Table)
	Errors() []error
}

func findPublisher(format string, config *Config, converter Converter) Publisher {
	if format == "markdown" {
		return newMarkdownPublisher(config, converter)
	}

	return newMarkdownPublisher(config, converter)
}
