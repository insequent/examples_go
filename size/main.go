package main

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/google/uuid"
)

func main() {
	a := uuid.New()
	b := "words"
	var c int
	var d int8
	var e int32
	var f int64
	var g uint
	var h uint8
	var i uint32
	var j uint64
	k := time.Now()

	fmt.Printf("a %T: %v, Size: %d bytes\n", a, a, unsafe.Sizeof(a))
	fmt.Printf("b %T: %v, Size: %d bytes\n", b, b, unsafe.Sizeof(b))
	fmt.Printf("c %T: %v, Size: %d bytes\n", c, c, unsafe.Sizeof(c))
	fmt.Printf("d %T: %v, Size: %d bytes\n", d, d, unsafe.Sizeof(d))
	fmt.Printf("e %T: %v, Size: %d bytes\n", e, e, unsafe.Sizeof(e))
	fmt.Printf("f %T: %v, Size: %d bytes\n", f, f, unsafe.Sizeof(f))
	fmt.Printf("g %T: %v, Size: %d bytes\n", g, g, unsafe.Sizeof(g))
	fmt.Printf("h %T: %v, Size: %d bytes\n", h, h, unsafe.Sizeof(h))
	fmt.Printf("i %T: %v, Size: %d bytes\n", i, i, unsafe.Sizeof(i))
	fmt.Printf("j %T: %v, Size: %d bytes\n", j, j, unsafe.Sizeof(j))
	fmt.Printf("k %T: %v, Size: %d bytes\n", k, k, unsafe.Sizeof(k))
}
