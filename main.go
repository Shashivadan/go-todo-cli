package main

import (
	"context"
	"fmt"
	"hello/util"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:  "create",
				Usage: "help to create a todo file",
				Action: func(ctx context.Context, c *cli.Command) error {
					fileName := fmt.Sprintf("todos/%s.csv", c.Args().Get(0))
					createfile, err := util.CreateCsvFile(fileName)
					if err != nil {
						return err
					}
					fmt.Println(createfile)
					return nil
				},
			},
			{
				Name:  "read",
				Usage: "list of all todos",
				Action: func(ctx context.Context, c *cli.Command) error {
					fileName := "todos/shashi.csv"
					data, err := util.ReadCsvFile(fileName)
					if err != nil {
						return err
					}
					for _, val := range data {
						fmt.Println(val)
					}
					return nil
				},
			},
			{
				Name:  "add",
				Usage: "add a todo item <task>",
				Action: func(ctx context.Context, c *cli.Command) error {
					todo := c.Args().First()

					data, err := util.AddTodo(todo, "todos/shashi.csv")
					if err != nil {
						return err
					}
					if data != nil {
						fmt.Println("task added")
					}
					return nil
				},
			},
			{
				Name:  "delete",
				Usage: "delete a specfie to do",
				Action: func(ctx context.Context, c *cli.Command) error {
					todoId := c.Args().First()
					data, err := util.DeleteTodo(todoId, "todos/shashi.csv")
					if err != nil {
						return err
					}
					fmt.Println(data)
					return nil
				},
			},
			{
				Name:  "doen",
				Usage: "make task complete",
				Action: func(ctx context.Context, c *cli.Command) error {
					todoId := c.Args().First()
					data, err := util.DoenTodo(todoId, "todos/shashi.csv")
					if err != nil {
						return err
					}
					fmt.Println(data)
					return nil
				},
			},
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
