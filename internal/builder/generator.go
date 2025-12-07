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

func (g *Generator) Started(time time.Time) (snowflake *Generator) {
	g.config.Started = time
	snowflake = g

	return
}

func (g *Generator) Skip(skip uint64) (snowflake *Generator) {
	g.config.Skip = skip
	snowflake = g

	return
}

func (g *Generator) Node(node uint16) (snowflake *Generator) {
	g.config.Node = node
	snowflake = g

	return
}

func (g *Generator) Worker(worker uint16) (snowflake *Generator) {
	g.config.Worker = worker
	snowflake = g

	return
}

func (g *Generator) Build() *core.Generator {
	return core.NewGenerator(g.config)
}
