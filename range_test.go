package jchengscripts

import (
	"testing"
)

// the jchengscript benchmarks demonstrates that overhead from use the range keyword when you need
// to update values inside a slice
//
//    $ go test -bench=.
//    goos: ...
//    goarch: ...
//    BenchmarkRange-4          500000              3327 ns/op
//    BenchmarkNoRange-4        500000              2856 ns/op
//    PASS


var _data [10240]int
var initialized = false

func init() {
	if !initialized {
		for idx, _ := range(_data) {
			_data[idx] = idx
		}
		initialized = true

	}
}


func BenchmarkRange(b *testing.B) {
	var i = 0
	for i < b.N {
		for idx, _ := range(_data) {
			_ = idx
		}
		i++
	}
}

func BenchmarkNoRange(b *testing.B) {
	var i = 0
	for i < b.N {
		j := 0
		end := len(_data)
		for j < end {
			_ = _data[j]
			j++
		}
		i++
	}
}