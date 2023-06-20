# debug

![Go](https://badgen.net/badge/Go/%3E=1.16/orange)
[![Go Version](https://badgen.net/github/release/go-packagist/debug/stable)](https://github.com/go-packagist/debug/releases)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-packagist/debug)](https://pkg.go.dev/github.com/go-packagist/debug)
[![codecov](https://codecov.io/gh/go-packagist/debug/branch/master/graph/badge.svg?token=5TWGQ9DIRU)](https://codecov.io/gh/go-packagist/debug)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-packagist/debug)](https://goreportcard.com/report/github.com/go-packagist/debug)
[![tests](https://github.com/go-packagist/debug/actions/workflows/go.yml/badge.svg)](https://github.com/go-packagist/debug/actions/workflows/go.yml)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

## Installation

```bash
go get github.com/go-packagist/debug
```

## Usage

```go
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
```

## License

- The MIT License (MIT). Please see [License File](LICENSE) for more information.
- Thanks to [go-spew](https://github.com/davecgh/go-spew) for the inspiration.
