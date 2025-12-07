package param

import (
	"sync"
	"time"

	"github.com/sony/sonyflake/v2"
)

var (
	generator *Generator
	once      sync.Once
)

type Generator struct {
	Started  time.Time
	Node     uint16
	Worker   uint16
	Machines int
}

func NewGenerator() *Generator {
	once.Do(func() { // 使用单例模式保证只有一份配置
		if nil == generator {
			generator = &Generator{
				Started:  time.Date(2025, time.December, 7, 12, 35, 30, 30, time.UTC),
				Node:     1,
				Worker:   1,
				Machines: 16,
			}
		}
	})

	return generator
}

func (g *Generator) Settings() (settings *sonyflake.Settings) {
	settings = new(sonyflake.Settings)
	if !g.Started.IsZero() {
		settings.StartTime = g.Started
	}
	settings.BitsMachineID = g.Machines
	settings.MachineID = func() (int, error) {
		return int(g.Node) * int(g.Worker), nil
	}
	settings.CheckMachineID = func(id int) bool {
		return id <= 2<<g.Machines
	}

	return
}
