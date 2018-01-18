package main

import (
	"os"
	"os/exec"
	"fmt"
)

func main() {
	if len(os.Args) == 1 {
		call_echo("{")
		call_echo("{}")
		call_echo("@{}")
	} else if len(os.Args) > 2 && os.Args[1] == "echo" {
		fmt.Println(os.Args[2])
	}

}

func call_echo(arg string) {
	executable := os.Args[0]
	cmd := exec.Command(executable, "echo", arg)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(out))
}