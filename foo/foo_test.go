package foo

//go test
import (
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(1, 2)
	if Add(1, 2) != 3 {
		t.Error("1 + 2 = 3 but get", sum)
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}
