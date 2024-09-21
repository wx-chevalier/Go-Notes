package main

import (
	"fmt"
	"net/http"

	"github.com/agtorre/go-cookbook/chapter13/vendoring/handlers"
	"github.com/agtorre/go-cookbook/chapter13/vendoring/models"
	"github.com/sirupsen/logrus"
)

func main() {
	c := handlers.NewController(models.NewDB())

	logrus.SetFormatter(&logrus.JSONFormatter{})

	http.HandleFunc("/get", c.GetHandler)
	http.HandleFunc("/set", c.SetHandler)
	fmt.Println("server started at localhost:8080")
	panic(http.ListenAndServe("localhost:8080", nil))
}
