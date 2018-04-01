package cli

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

// ParseApp parse JSON file: flag: --parse [fileName]
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
		fmt.Println(fileName)
		return nil
	}
	app.Run(os.Args)

}
