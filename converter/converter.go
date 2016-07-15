package converter

import (
	"fmt"

	"github.com/pinzolo/dbmodel"
)

// Converter convert column data to string slice.
type Converter interface {
	Convert(c *dbmodel.Column) []string
}

// FindConverter find Converter implementation for given driver.
func FindConverter(driver string) (Converter, error) {
	if driver == "postgres" {
		return postgres{}, nil
	}
	return nil, fmt.Errorf("Driver '%v' is unknown driver.", driver)
}

type postgres struct{}

func (p postgres) Convert(c *dbmodel.Column) []string {
	null := "NO"
	if c.IsNullable() {
		null = ""
	}
	return []string{
		c.Name(),
		c.DataType(),
		c.Size().String(),
		null,
		c.DefaultValue(),
		c.Comment(),
	}
}
