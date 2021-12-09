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

package shark

import (
	"errors"
	"github.com/spf13/pflag"
	"sync"

	"github.com/spf13/cobra"
)

type Command = *cobra.Command
type FlagSet = *pflag.FlagSet
type FlagSetFunc = func(flags FlagSet)

type Shark struct {
	command Command
	access  sync.Mutex
}

func (super *Shark) Command() Command {
	return super.command
}

func (super *Shark) SetFlags(set FlagSetFunc) {
	super.access.Lock()
	defer super.access.Unlock()
	set(super.Command().Flags())
}

// NewShark Shark is the cobra command structure wrapper.
func NewShark(ops ...Option) *Shark {
	options := options{}

	if len(ops) == 0 {
		panic(errors.New("shark: options value invalid"))
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

	return &Shark{command: command}
}
