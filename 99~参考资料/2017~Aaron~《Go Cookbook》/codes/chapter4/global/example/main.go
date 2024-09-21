package main

import "github.com/agtorre/go-cookbook/chapter4/global"

func main() {
	if err := global.UseLog(); err != nil {
		panic(err)
	}
}
