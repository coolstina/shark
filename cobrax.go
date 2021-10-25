// Copyright 2021 helloshaohua <wu.shaohua@foxmail.com>;
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cobrax

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type Command struct {
	Cobra *cobra.Command
}

func (c *Command) SetFlags(set func(flags *pflag.FlagSet)) {
	flags := c.Cobra.Flags()
	set(flags)
}

// NewCommand Command is the cobra command structure wrapper.
func NewCommand(ops ...Option) *Command {
	options := options{}

	if len(ops) == 0 {
		panic(errors.New("cobrax: options value invalid"))
	}

	for _, o := range ops {
		o.apply(&options)
	}

	command := &cobra.Command{
		Use:   options.use,
		Short: options.short,
		Long:  options.long,
		Run:   options.run,
	}

	return &Command{Cobra: command}
}
