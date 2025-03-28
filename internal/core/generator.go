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
	settings := new(spaceflake.GeneratorSettings)
	if !config.Start.IsZero() {
		settings.BaseEpoch = uint64(config.Start.Unix())
	}
	if 0 != config.Node {
		settings.NodeID = config.Node
	}
	if 0 != config.Worker {
		settings.WorkerID = config.Worker
	}
	if 0 != config.Sequence {
		settings.Sequence = config.Sequence
	}

	return &Generator{
		settings: settings,
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
