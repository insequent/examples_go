package main

import (
	"fmt"
	"math"
	"testing"
)

func TestRandomString(t *testing.T) {
	for i := 0; i < 10; i++ {
		str := randomString(i)

		if len(str) != i {
			t.Fatalf("Unexpected length returned by randomString(): %d", len(str))
		}
	}
}

func TestRandomBytes(t *testing.T) {
	for i := 0; i < 10; i++ {
		b := randomBytes(i)

		if len(b) != i {
			t.Fatalf("Unexpected length returned by randomBytes(): %d", len(b))
		}
	}
}

func BenchmarkRandomString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		randomString(n)
	}
}

func BenchmarkRandomBytes(b *testing.B) {
	for n := 0; n < b.N; n++ {
		randomBytes(n)
	}
}

func i2size(i int) int {
	return int(math.Pow10(i))
}

func BenchmarkAllocateByteArray(b *testing.B) {
	for i := 1; i <= 10; i++ {
		b.Run(fmt.Sprintf("test_array_size_%d", i2size(i)), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = make([]byte, i2size(i))
			}
		})
	}
}
