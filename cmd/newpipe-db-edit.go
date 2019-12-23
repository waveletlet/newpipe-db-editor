package main

import (
	"log"
	"os"

	"gopkg.in/urfave/cli.v2"
	//npdb "gitlab.com/waveletlet/newpipe-db-editor"
)

func main() {

	// Set up cli
	cmd := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "dump",
				Usage: "Dumps database or playlist to format (csv)",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "format",
						Aliases: []string{"f"},
						Usage:   "Dumps to `FORMAT` (csv)",
						// does it make sense to dump as playlist format for vlc, etc?
					},
				},
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name: "list",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
		},
	}
	cmd.Name = "NewPipe DB Editor"
	cmd.Usage = "Edit sqlite database exported from NewPipe"
	cmd.UseShortOptionHandling = true

	err := cmd.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
