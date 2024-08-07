package main

import "github.com/NSXBet/protoc-gen-go-cache-manager/cmd"

func main() {
	// generate code here
	if err := cmd.NewGenerator().Generate(); err != nil {
		panic(err)
	}
}
