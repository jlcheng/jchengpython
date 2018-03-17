package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	var in io.Reader
	if len(os.Args) > 1 {
		fname := os.Args[1]
		var err error
		if in, err = os.Open(fname); err != nil {
			fmt.Fprintf(os.Stderr, "cannot open file %v: %v\n", fname, err)
		}
	} else {
		in = os.Stdin
	}
	dec := json.NewDecoder(in)
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	var buf interface{}
	if err := dec.Decode(&buf); err != nil {
		fmt.Fprintf(os.Stderr, "cannot decode input:%v\n", err)
		os.Exit(1)
	} else {
		enc.Encode(buf)
	}
}
