package cli

import (
	"encoding/json"
	"fmt"
	"os"
)

import "github.com/urfave/cli"

type config struct {
	Os    string
	Distr string
}

// ParseApp parse JSON file: flag: --parse [fileName]
// parses struct:
// os    string
// distr string
func ParseApp() {
	app := cli.NewApp()
	app.Name = "Parse JSON file"
	app.Usage = "JSON file parser"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "parse, p",
			Usage: "parse JSON file",
			Value: "bad file",
		},
	}
	app.Action = func(c *cli.Context) error {
		fileName := c.GlobalString("parse")
		file, err := os.Open(fileName)
		defer file.Close()
		if err != nil {
			fmt.Println("bad file")
			return err
		}
		decoder := json.NewDecoder(file)
		conf := config{}

		err = decoder.Decode(&conf)
		if err != nil {
			fmt.Println("failed to decode")
			return err
		}
		fmt.Println(conf)
		return err
	}
	app.Run(os.Args)

}
