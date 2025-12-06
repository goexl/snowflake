package param

import (
	"sync"
	"time"

	"github.com/kkrypt0nn/spaceflake"
)

var (
	generator *Generator
	once      sync.Once
)

type Generator struct {
	Skip     uint64
	Base     time.Time
	Node     uint64
	Worker   uint64
	Sequence uint64
}

func NewGenerator() *Generator {
	once.Do(func() { // 使用单例模式保证只有一份配置
		if nil == generator {
			generator = &Generator{
				Node:   1,
				Worker: 1,
			}
		}
	})

	return generator
}

func (g *Generator) Settings() (settings *spaceflake.GeneratorSettings) {
	settings = new(spaceflake.GeneratorSettings)
	if !g.Base.IsZero() {
		settings.BaseEpoch = uint64(g.Base.UnixMilli())
	}
	if 0 != g.Node {
		settings.NodeID = g.Node
	}
	if 0 != g.Worker {
		settings.WorkerID = g.Worker
	}
	if 0 != g.Sequence {
		settings.Sequence = g.Sequence
	}

	return
}
