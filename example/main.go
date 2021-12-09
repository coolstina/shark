package main

import (
	"fmt"

	"github.com/coolstina/shark"
	"github.com/spf13/cobra"
)

func main() {
	sharkcmd := shark.NewShark(
		shark.WithUse("hello"),
		shark.WithLong("Say hello"),
		shark.WithLong("Execute this command output say hello"),
		shark.WithRun(func(cmd *cobra.Command, args []string) {
			fmt.Printf("hello world")
		}),
	)

	if err := sharkcmd.Command().Execute(); err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}
