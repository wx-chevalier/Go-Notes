package main

import (
	"github.com/agtorre/go-cookbook/chapter10/orchestrate"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	session, err := mgo.Dial("mongodb")
	if err != nil {
		panic(err)
	}
	if err := orchestrate.ConnectAndQuery(session); err != nil {
		panic(err)
	}
}
