package jchengscripts

import (
	"testing"
	"fmt"
)

type Foo struct {
	Id int
}
var container = make([]Foo, 5)

func init() {
	for idx, _ := range container {
		container[idx].Id = idx
	}
	fmt.Println(container[1])
}

func TestRange(t *testing.T) {

}
