package snowflake_test

import (
	"testing"

	"github.com/goexl/snowflake"
	"github.com/stretchr/testify/assert"
)

func TestNext(t *testing.T) {
	id, err := snowflake.New().Build().Next()
	assert.Nil(t, err)
	assert.NotNil(t, id)
}

func BenchmarkNext(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = snowflake.New().Build().Next()
	}
}
