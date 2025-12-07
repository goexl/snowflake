package snowflake_test

import (
	"fmt"
	"testing"

	"github.com/goexl/snowflake"
	"github.com/stretchr/testify/assert"
)

func TestNext(t *testing.T) {
	id, err := snowflake.New().Build().Next()
	assert.Nil(t, err)
	assert.NotNil(t, id)
	fmt.Println(id.Time())
}

func BenchmarkNext(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = snowflake.New().Build().Next()
	}
}
