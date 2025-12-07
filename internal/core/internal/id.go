package internal

import (
	"strconv"
	"time"

	"github.com/goexl/id"
	"github.com/goexl/snowflake/internal/internal/param"
	"github.com/sony/sonyflake/v2"
)

var _ id.Value = (*Id)(nil)

type Id struct {
	value uint64
	flake *sonyflake.Sonyflake
}

func NewId(value uint64, flake *sonyflake.Sonyflake) *Id {
	return &Id{
		value: value,
		flake: flake,
	}
}

func Parse(from uint64) (id *Id) {
	id = new(Id)
	id.value = from
	id.parseFlake()

	return
}

func (i *Id) String() string {
	return strconv.FormatUint(i.value, 10)
}

func (i *Id) Time() time.Time {
	return i.flake.ToTime(int64(i.value))
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
	i.parseFlake()
}

func (i *Id) parseFlake() {
	settings := new(sonyflake.Settings)
	started := param.NewGenerator().Started
	if !started.IsZero() {
		settings.StartTime = started
	}
	if flake, err := sonyflake.New(*settings); nil != err {
		panic(err)
	} else {
		(*i).flake = flake
	}
}
