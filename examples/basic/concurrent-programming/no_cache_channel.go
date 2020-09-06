package main

var done = make(chan bool)
var msg string

func aGoroutine() {
	msg = "你好, 世界"
	done <- true
}

func main_1() {
	go aGoroutine()

	<-done

	println(msg)
}
