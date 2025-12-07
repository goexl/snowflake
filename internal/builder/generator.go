package builder

import (
	"time"

	"github.com/goexl/snowflake/internal/core"
	"github.com/goexl/snowflake/internal/internal/param"
)

type Generator struct {
	config *param.Generator
}

func NewGenerator() *Generator {
	return &Generator{
		config: param.NewGenerator(),
	}
}

func (g *Generator) Started(time time.Time) (generator *Generator) {
	g.config.Started = time
	generator = g

	return
}

func (g *Generator) Machine(machine uint32) (generator *Generator) {
	g.config.Machine = machine
	generator = g

	return
}

func (g *Generator) Build() *core.Generator {
	return core.NewGenerator(g.config)
}
