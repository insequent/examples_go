package main

import (
	"fmt"
	"github.com/google/uuid"
	"time"
	"unsafe"
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

	fmt.Printf("a %T: %v, %d\n", a, a, unsafe.Sizeof(a))
	fmt.Printf("b %T: %v, %d\n", b, b, unsafe.Sizeof(b))
	fmt.Printf("c %T: %v, %d\n", c, c, unsafe.Sizeof(c))
	fmt.Printf("d %T: %v, %d\n", d, d, unsafe.Sizeof(d))
	fmt.Printf("e %T: %v, %d\n", e, e, unsafe.Sizeof(e))
	fmt.Printf("f %T: %v, %d\n", f, f, unsafe.Sizeof(f))
	fmt.Printf("g %T: %v, %d\n", g, g, unsafe.Sizeof(g))
	fmt.Printf("h %T: %v, %d\n", h, h, unsafe.Sizeof(h))
	fmt.Printf("i %T: %v, %d\n", i, i, unsafe.Sizeof(i))
	fmt.Printf("j %T: %v, %d\n", j, j, unsafe.Sizeof(j))
	fmt.Printf("k %T: %v, %d\n", k, k, unsafe.Sizeof(k))
}
