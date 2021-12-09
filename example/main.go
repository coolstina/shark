package main

import (
	"fmt"

	cobrax "github.com/coolstina/cobrasuper"
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

	if err := command.Command().Execute(); err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}
