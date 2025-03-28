package internal

import (
	"time"

	"github.com/goexl/id"
	"github.com/kkrypt0nn/spaceflake"
)

var _ id.Id = (*Id)(nil)

type Id struct {
	value *spaceflake.Spaceflake
}

func NewId(value *spaceflake.Spaceflake) *Id {
	return &Id{
		value: value,
	}
}

func (i *Id) String() string {
	return i.value.StringID()
}

func (i *Id) Time() time.Time {
	return time.UnixMilli(int64(i.value.Time()))
}

func (i *Id) Value() uint64 {
	return i.value.ID()
}
