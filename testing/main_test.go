package main

import (
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
