package core

import (
	"github.com/goexl/id"
	"github.com/goexl/snowflake/internal/core/internal"
	"github.com/goexl/snowflake/internal/internal/param"
	"github.com/sony/sonyflake/v2"
)

var _ id.Generator = (*Generator)(nil)

type Generator struct {
	settings *sonyflake.Settings
	flake    *sonyflake.Sonyflake
}

func NewGenerator(config *param.Generator) *Generator {
	return &Generator{
		settings: config.Settings(),
	}
}

func (g *Generator) Next() (value id.Value, err error) {
	if ie := g.init(); nil != ie {
		err = ie
	} else if next, ge := g.flake.NextID(); nil != ge {
		err = ge
	} else {
		value = internal.NewId(uint64(next), g.flake)
	}

	return
}

func (g *Generator) Parse(from uint64) id.Value {
	return internal.Parse(from)
}

func (g *Generator) init() (err error) {
	if nil == g.flake {
		g.flake, err = sonyflake.New(*g.settings)
	}

	return
}
