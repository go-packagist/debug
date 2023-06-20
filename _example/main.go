package main

import (
	"github.com/go-packagist/debug"
	"os"
)

func main() {
	debug.Dump("hello world", []byte("hello world"), 123, 123.456, true, nil)
	debug.Sdump("hello world: %s", "Lily")
	debug.Fdump(os.Stdout, "hello world")

	// Output:
	// (string) (len=11) "hello world"
	// ([]uint8) (len=11 cap=11) {
	// 	00000000  68 65 6c 6c 6f 20 77 6f  72 6c 64                 |hello world|
	// }
	// (int) 123
	// (float64) 123.456
	// (bool) true
	// (interface {}) <nil>
	// (string) (len=11) "hello world"
}
