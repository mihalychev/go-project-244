package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3" // imports as package "cli"

	"code"
)

func main() {
	cmd := &cli.Command{
		Name:  "gendiff",
		Usage: "Compares two configuration files and shows a difference.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"f"},
				Usage:   "output format (default: \"stylish\")",
			},
		},
		Action: func(_ context.Context, cmd *cli.Command) error {
			if cmd.Args().Len() != 2 {
				return cli.ShowAppHelp(cmd)
			}

			filepath1 := cmd.Args().Get(0)
			filepath2 := cmd.Args().Get(1)
			diff, err := code.Gendiff(cmd.String("format"), filepath1, filepath2)
			if err != nil {
				return err
			}

			fmt.Println(diff)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
