package main

import (
	"fmt"

	"github.com/coolstina/cobrasuper"
	"github.com/spf13/cobra"
)

func main() {
	command := cobrasuper.NewCommand(
		cobrasuper.WithUse("hello"),
		cobrasuper.WithLong("Say hello"),
		cobrasuper.WithLong("Execute this command output say hello"),
		cobrasuper.WithRun(func(cmd *cobra.Command, args []string) {
			fmt.Printf("hello world")
		}),
	)

	if err := command.Command().Execute(); err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}
