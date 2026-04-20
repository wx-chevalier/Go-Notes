package main

import "runtime"
import "fmt"

func main() {
	runtime.GOMAXPROCS(2)
	go func() {
		// 永远不会输出
		fmt.Println("hello world")
	}()
	go func() {
		for {

		}
	}()
	select {}
}
