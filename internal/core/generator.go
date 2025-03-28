package core

import (
	"github.com/goexl/id"
	"github.com/goexl/snowflake/internal/core/internal"
	"github.com/goexl/snowflake/internal/internal/param"
	"github.com/kkrypt0nn/spaceflake"
)

var _ id.Generator = (*Generator)(nil)

type Generator struct {
	settings spaceflake.GeneratorSettings
}

func NewGenerator(config *param.Generator) *Generator {
	return &Generator{
		settings: spaceflake.GeneratorSettings{
			BaseEpoch: uint64(config.Start.Unix()),
			NodeID:    config.Node,
			Sequence:  config.Sequence,
			WorkerID:  config.Worker,
		},
	}
}

func (s *Generator) Next() (id id.Id, err error) {
	if next, ge := spaceflake.Generate(s.settings); nil != ge {
		err = ge
	} else {
		id = internal.NewId(next)
	}

	return
}
