package main

import (
	"fmt"
	"log"
	"os"

	client "dev.wx/client"
	model "dev.wx/model"
	"github.com/urfave/cli/v2"
)

var (
	bpClient, _ = client.NewBpClient("")
)

func main() {
	app := &cli.App{
		Name:  "go-http-proxy-cli",
		Usage: "Proxy wrapper with Golang.",
		Commands: []*cli.Command{
			{
				Name:    "greet",
				Aliases: []string{"g"},
				Usage:   "Say greetings",
				Action: func(c *cli.Context) error {
					fmt.Println("Say greetingss")
					return nil
				},
			},
			{
				Name:    "bpc",
				Aliases: []string{"b"},
				Usage:   "Send msg to bpc",
				Action: func(c *cli.Context) error {
					_, err := bpClient.CreateTask(
						model.NewFile("1", "INI", "remote", "http://172.16.7.208:8000/ini/SLA_600.ini", "084b3157cae7198b1ce3c5871beb8760"),
						model.NewFile("2", "BPP", "remote", "http://172.16.7.208:8000/bpp/%E8%80%81%E6%AC%BE3300-%E6%AD%A3%E9%9B%85600-0902.bpp", "0f77314f93832ae3c7d088dbe28c00b9"),
						[]*model.File{
							model.NewFile("3", "CLI", "remote", "http://172.16.7.208:8000/cli/eStageMergedPart_s.cli", "c41c5b6a11aa56c2357c022ae151275a"),
						})

					return err
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
