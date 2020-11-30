package main

import (
	"Order_Map"
	"fmt"

	"math/rand"

	"reflect"
	"time"
)

func main() {

	t := time.Now()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	testmap := Order_Map.NewIntMap(10000)
	t1 := t.Second()
	for i := 0; i < 10000; i++ {
		testmap.Insert(r.Intn(10000), r.Intn(10000))
	}
	t = time.Now()
	t2 := t.Second()
	fmt.Println("insert  time  span", t2-t1)
	testmap.Erase(88)
	for i := 0; i < testmap.Size(); i++ {
		k, v, _ := testmap.GetByOrderIndex(i)
		tmp_v := reflect.ValueOf(v)
		fmt.Println("k:", k, "---", "value:", tmp_v)
	}

	t = time.Now()
	t3 := t.Second()
	fmt.Println("range  time  span:", t3-t2)

}
