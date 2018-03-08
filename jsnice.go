package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	buf := make(map[string]interface{})
	if err := dec.Decode(&buf); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		enc.Encode(buf)
	}
}
