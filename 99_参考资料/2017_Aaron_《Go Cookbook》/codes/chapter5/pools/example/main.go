package main

import "github.com/agtorre/go-cookbook/chapter5/pools"

func main() {
	if err := pools.ExecWithTimeout(); err != nil {
		panic(err)
	}
}
