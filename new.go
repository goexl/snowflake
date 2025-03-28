package snowflake

import (
	"github.com/goexl/snowflake/internal/builder"
)

func New() *builder.Generator {
	return builder.NewGenerator()
}
