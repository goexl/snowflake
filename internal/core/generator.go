package core

import (
	"time"

	"github.com/goexl/id"
	"github.com/goexl/snowflake/internal/core/internal"
	"github.com/goexl/snowflake/internal/internal/param"
	"github.com/kkrypt0nn/spaceflake"
)

var _ id.Generator = (*Generator)(nil)

type Generator struct {
	settings *spaceflake.GeneratorSettings
}

func NewGenerator(config *param.Generator) *Generator {
	return &Generator{
		settings: config.Settings(),
	}
}

func (s *Generator) Next() (value id.Value, err error) {
	if next, ge := spaceflake.Generate(*s.settings); nil != ge {
		err = ge
	} else {
		value = internal.NewId(next.ID(), time.UnixMilli(int64(next.Time())))
	}

	return
}
