package main

import (
	"context"

	"github.com/agtorre/go-cookbook/chapter6/oauthcli"
)

func main() {
	ctx := context.Background()
	conf := oauthcli.Setup()

	tok, err := oauthcli.GetToken(ctx, conf)
	if err != nil {
		panic(err)
	}
	client := conf.Client(ctx, tok)

	if err := oauthcli.GetUsers(client); err != nil {
		panic(err)
	}

}
