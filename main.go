package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var debug bool

	app := &cli.App{
		Name:                 "meltedbrain",
		Usage:                "Brainfuck interpreter",
		Version:              "v0.0.1",
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "debug",
				Aliases:     []string{"d"},
				Value:       false,
				Usage:       "prints the compiled instructions",
				Destination: &debug,
			},
		},
		Action: func(ctx *cli.Context) error {
			file := ctx.Args().First()
			code, err := ioutil.ReadFile(file)
			if err != nil {
				return err
			}

			c := newCompiler(string(code))
			instructions := c.compile()
			if debug {
				fileStats, err := os.Stat(file)
				if err != nil {
					return err
				}

				fmt.Println("file name: ", fileStats.Name())
				fmt.Println("file size: ", fileStats.Size())
				for i := 0; i < len(instructions); i++ {
					fmt.Printf("{ it: %v, arg: %v }\n", string(instructions[i].it), instructions[i].arg)
				}

				return nil
			}

			m := newMachine(instructions, os.Stdin, os.Stdout)
			if err := m.execute(); err != nil {
				return err
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
