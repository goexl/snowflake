package param

import (
	"time"
)

type Generator struct {
	Start    time.Time
	Node     uint64
	Worker   uint64
	Sequence uint64
}

func NewGenerator() *Generator {
	return &Generator{
		Node:   1,
		Worker: 1,
	}
}
