package main

import (
	"fmt"

	"github.com/bandarji/aoc/adventofcode"
	"github.com/spf13/cobra"
)

func cliCmd() *cobra.Command {
	params := runParams{}
	rootCmd := &cobra.Command{
		Use:   "make YEAR=2015 DAY=1 run",
		Short: "Advent of Code solutions in Go",
		Long:  "Advent of Code puzzles solved with Go. Use `make` to see build and execution options.",
		Run:   func(cmd *cobra.Command, args []string) { solve(params) },
	}
	processParams(rootCmd, &params)
	return rootCmd
}

func processParams(cmd *cobra.Command, params *runParams) {
	cmd.Flags().IntVarP(&params.year, "year", "y", 0, "Year of the puzzle")
	cmd.Flags().IntVarP(&params.day, "day", "d", 0, "Day of the puzzle")
}

func solve(params runParams) {
	if data, err := adventofcode.NewAOCDay(params.year, params.day); err == nil {
		adventofcode.RunDay(data, params.year, params.day)
	} else {
		fmt.Println(err)
	}
}
