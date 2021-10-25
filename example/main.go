package main

import (
	"fmt"

	"github.com/coolstina/cobrax"
	"github.com/spf13/cobra"
)

func main() {
	command := cobrax.NewCommand(
		cobrax.WithUse("hello"),
		cobrax.WithLong("Say hello"),
		cobrax.WithLong("Execute this command output say hello"),
		cobrax.WithRun(func(cmd *cobra.Command, args []string) {
			fmt.Printf("hello world")
		}),
	)

	if err := command.Cobra.Execute(); err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}
