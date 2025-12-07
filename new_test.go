package snowflake_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/goexl/snowflake"
	"github.com/stretchr/testify/assert"
)

func TestNext(t *testing.T) {
	fmt.Println(time.Now().UTC().Hour())
	id, err := snowflake.New().Epoch(time.Date(2025, time.December, 7, 12, 31, 23, 12, time.UTC)).Build().Next()
	assert.Nil(t, err)
	assert.NotNil(t, id)
	fmt.Println(id.Get(), ": ", id.Time())
}

func BenchmarkNext(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = snowflake.New().Build().Next()
	}
}
