package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "dpfile",
		Usage: "A cli duplicates a file with any name and and sequential number",
		Authors: []*cli.Author{
			{
				Name:  "Keisuke Umegaki",
				Email: "keisuke.umegaki.630@gmail.com",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "src",
				Aliases:  []string{"s"},
				Usage:    "path to the file you want to duplicate",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "dst",
				Aliases: []string{"d"},
				Usage:   "path to the destination which the file will be duplicated",
				Value:   "./",
			},
			&cli.StringFlag{
				Name:    "filename",
				Aliases: []string{"f"},
				Value:   "",
			},
			&cli.StringFlag{
				Name:    "offset",
				Aliases: []string{"o"},
				Value:   "0",
				Usage:   "-1 < offset",
			},
			&cli.StringFlag{
				Name:    "limit",
				Aliases: []string{"l"},
				Value:   "1",
				Usage:   "0 < limit < 10001",
			},
		},
		Action: func(c *cli.Context) error {
			offset, err := newOffset(c.String("offset"))
			if err != nil {
				return err
			}
			limit, err := newLimit(c.String("limit"))
			if err != nil {
				return err
			}
			dp, err := newDuplicator(c.String("src"), c.String("dst"), c.String("filename"), offset, limit)
			if err != nil {
				return err
			}
			return dp.duplicate()
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
