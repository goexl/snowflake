package snowflake_test

import (
	"testing"

	"github.com/goexl/snowflake"
)

func BenchmarkNext(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = snowflake.New().Build().Next()
	}
}
