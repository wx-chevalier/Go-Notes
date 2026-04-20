package handlers

import "github.com/agtorre/go-cookbook/chapter13/vendoring/models"

type Controller struct {
	db models.DB
}

func NewController(db models.DB) *Controller {
	return &Controller{db: db}
}

type resp struct {
	Status string `json:"status"`
	Value  int64  `json:"value"`
}
