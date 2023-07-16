package main

import "github.com/iBaiYang/go-toolbox/cmd"
import "log"

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
