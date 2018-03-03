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


var _slice = make([]int, 10240)
var _array [10240]int

func BenchmarkRangeSlice(b *testing.B) {
	for n := 0 ;n < b.N; n++ {
		for idx, _ := range(_slice) {
			_slice[idx] = idx
		}
	}
	if _slice[2] != 2 {
		b.Error("_slice was not modified")
	}
}

func BenchmarkResetSlice(b *testing.B) {
	_slice = make([]int, 10240)
	if _slice[2] != 0 {
		b.Error("_slice was not reset")
	}
}

func BenchmarkNoRangeSlice(b *testing.B) {
	var end = len(_slice)
	for n := 0; n < b.N; n++ {
		for idx := 0; idx < end; idx++ {
			_slice[idx] = idx
		}
	}
	if _slice[2] != 2 {
		b.Error("_slice was not modified")
	}
}

func BenchmarkRangeArray(b *testing.B) {
	for n := 0 ;n < b.N; n++ {
		for idx, _ := range(_array) {
			_array[idx] = idx
		}
	}
	if _array[2] != 2 {
		b.Error("_array was not modified")
	}
}

func BenchmarkResetArray(b *testing.B) {
	_array = [10240]int{}
	if _array[2] != 0 {
		b.Error("_array was not reset")
	}
}


func BenchmarkNoRangeArray(b *testing.B) {
	var end = len(_array)
	for n := 0; n < b.N; n++ {
		for idx := 0; idx < end; idx++ {
			_array[idx] = idx
		}
	}
	if _array[2] != 2 {
		b.Error("_array was not modified")
	}
}