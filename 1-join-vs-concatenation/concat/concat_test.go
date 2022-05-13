package concat_test

import (
	"github.com/vinigofr/a-linguagem-go/1-join-vs-concatenation/concat"
	"testing"
)

func BenchmarkWithOperator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concat.WithOperador()
	}
}

func BenchmarkWithJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concat.WithJoin()
	}
}
