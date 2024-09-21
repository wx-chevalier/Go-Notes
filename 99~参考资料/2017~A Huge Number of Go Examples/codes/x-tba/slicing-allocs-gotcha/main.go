
package main

// Before running this program:
//
//     RUN: `go run nums/main.go`
//     It will create a 12 MB long nums.txt file.

import (
	"fmt"
	"io/ioutil"
	"runtime"

	s "github.com/inancgumus/prettyslice"
)

const (
	loops = 1000
	file  = "nums.txt"
)

var buf []byte

func main() {
	{
		b, _ := ioutil.ReadFile(file)
		buf = b[:1]

		s.Show("sliced buf", buf)
	}
	report()

	{
		var nilBuf []byte

		b, _ := ioutil.ReadFile(file)
		buf = append(nilBuf, b[:1]...)

		s.Show("copied buf", buf)
	}
	report()
}

func report() {
	const KB = 1024

	var m runtime.MemStats
	r := func() {
		runtime.ReadMemStats(&m)
		fmt.Printf("%v KB\n", m.Alloc/KB)
	}

	fmt.Print(" > Before Garbage Collection: ")
	r()

	runtime.GC()
	fmt.Print(" > After Garbage Colletion  : ")
	r()
}

func init() {
	s.Width = 50
	s.PrintBacking = true
	s.MaxElements = 10
}
