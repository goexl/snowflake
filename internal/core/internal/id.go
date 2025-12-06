package internal

import (
	"strconv"
	"time"

	"github.com/goexl/id"
	"github.com/goexl/snowflake/internal/internal/param"
	"github.com/kkrypt0nn/spaceflake"
)

var _ id.Value = (*Id)(nil)

type Id struct {
	value uint64
	time  time.Time
}

func NewId(value uint64, time time.Time) *Id {
	return &Id{
		value: value,
		time:  time,
	}
}

func Parse(from uint64) (id *Id) {
	id = new(Id)
	id.value = from
	id.parseTime()

	return
}

func (i *Id) String() string {
	return strconv.FormatUint(i.value, 10)
}

func (i *Id) Time() time.Time {
	return i.time
}

func (i *Id) Get() uint64 {
	return i.value - param.NewGenerator().Skip
}

func (i *Id) MarshalJSON() ([]byte, error) {
	return i.ToDB()
}

func (i *Id) UnmarshalJSON(from []byte) error {
	return i.FromDB(from)
}

func (i *Id) FromDB(from []byte) (err error) {
	if parsed, pue := strconv.ParseUint(string(from), 10, 64); nil != pue {
		err = pue
	} else {
		i.from(parsed)
	}

	return
}

func (i *Id) ToDB() ([]byte, error) {
	return []byte(strconv.FormatUint(i.value, 10)), nil
}

func (i *Id) from(value uint64) {
	(*i).value = value
	i.parseTime()
}

func (i *Id) parseTime() {
	base := uint64(0)
	epoch := param.NewGenerator().Base
	if !epoch.IsZero() {
		base = uint64(epoch.UnixMilli())
	}
	(*i).time = time.UnixMilli(int64(spaceflake.ParseTime(i.value, base)))
}
