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


var _data = make([]int, 10240)
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
	for n := 0 ;n < b.N; n++ {
		for idx, _ := range(_data) {
			_ = idx
		}
	}
	if _data[2] != 2 {
		b.Error("_data was not modified")
	}
}

func BenchmarkNoRange(b *testing.B) {
	var end = len(_data)
	for n := 0; n < b.N; n++ {
		for idx := 0; idx < end; idx++ {
			_data[idx] = idx
		}
	}
	if _data[2] != 2 {
		b.Error("_data was not modified")
	}
}