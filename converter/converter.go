package converter

import (
	"fmt"
	"strconv"

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
	var pkPosition string
	if c.PrimaryKeyPosition() > 0 {
		pkPosition = strconv.FormatInt(c.PrimaryKeyPosition(), 10)
	}
	return []string{
		pkPosition,
		c.Name(),
		c.DataType(),
		c.Size().String(),
		null,
		c.DefaultValue(),
		c.Comment(),
	}
}
