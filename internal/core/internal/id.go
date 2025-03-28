package internal

import (
	"strconv"
	"strings"
	"time"

	"github.com/goexl/id"
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

func (i *Id) String() string {
	return strconv.FormatUint(i.value, 10)
}

func (i *Id) Time() time.Time {
	return i.time
}

func (i *Id) Get() uint64 {
	return i.value
}

func (i *Id) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatUint(i.Get(), 10)), nil
}

func (i *Id) UnmarshalJSON(from []byte) (err error) {
	converted := strings.Trim(string(from), `"`) // 去除字符串的引号
	if value, pue := strconv.ParseUint(converted, 10, 64); nil != pue {
		err = pue
	} else {
		(*i).value = value
	}

	return
}

func (i *Id) FromDB(from []byte) (err error) {
	if parsed, pue := strconv.ParseUint(string(from), 10, 64); nil != pue {
		err = pue
	} else {
		(*i).value = parsed
	}

	return
}

func (i *Id) ToDB() ([]byte, error) {
	return []byte(strconv.FormatUint(i.value, 10)), nil
}
