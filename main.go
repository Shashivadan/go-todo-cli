package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:  "create",
				Usage: "create a csv file to add todos",
				Action: func(ctx context.Context, c *cli.Command) error {
					fileName := fmt.Sprintf("todos/%s.csv", c.Args().Get(0))

					if c.Args().Get(0) == "" {
						return errors.New("file name missing")
					}
					_, err := os.Create(fileName)
					if err != nil {
						fmt.Println("there is an error ", err)
						return err
					}
					fmt.Println("create a file :", fileName)
					return nil
				},
			},
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
