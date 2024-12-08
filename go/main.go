package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/jwillisnetdev/advent_of_code2024/go/pkg/input"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:    "run",
				Aliases: []string{"r"},
				Usage:   "Run a specific day, or run all days.",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if cmd.Args().Len() == 0 {
						fmt.Println("Running all days...")
						return runAllDays()
					}
					if cmd.Args().Len() == 1 {
						dayNum, err := strconv.Atoi(cmd.Args().Get(0))
						if err != nil {
							return err
						}
						return runDay(dayNum)
					}
					return nil
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		input.Error(err.Error())
	}
}

func runDay(day int) error {
	fp := fmt.Sprintf("2024/Day_%02d/main.go", day)

	fname := filepath.Base(fp)
	fargs := []string{"run", fname, "input.txt"}
	cmd := exec.Command("go", fargs...)
	cmd.Dir = filepath.Dir(fp)
	stdout, stderr := cmd.CombinedOutput()
	if len(stdout) > 0 {
		fmt.Printf("%v\n", string(stdout))
	}
	if stderr != nil {
		return errors.New(stderr.Error())
	}

	return nil
}

func runAllDays() error {
	files, _ := filepath.Glob("2024/Day_*/main.go")
	for _, fp := range files {
		fname := filepath.Base(fp)
		fargs := []string{"run", fname, "input.txt"}
		cmd := exec.Command("go", fargs...)
		cmd.Dir = filepath.Dir(fp)
		stdout, stderr := cmd.CombinedOutput()
		if len(stdout) > 0 {
			fmt.Printf("%v\n", string(stdout))
		}
		if stderr != nil {
			return errors.New(stderr.Error())
		}
	}

	return nil
}
