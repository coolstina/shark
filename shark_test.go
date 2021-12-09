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
	"fmt"
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/suite"
)

func TestCommandSuite(t *testing.T) {
	suite.Run(t, new(CommandSuite))
}

type CommandSuite struct {
	suite.Suite
	Command *Shark
}

func (suite *CommandSuite) BeforeTest(suiteName, testName string) {
	suite.Command = NewShark(
		WithUse("hello"),
		WithLong("Say hello"),
		WithLong("Execute this command output say hello"),
		WithRun(func(cmd Command, args []string) {
			fmt.Printf("hello world")
		}),
	)
}

func (suite *CommandSuite) TestCommand_SetFlags() {
	suite.Command.SetFlags(func(flags *pflag.FlagSet) {
		flags.StringP("hello", "l", "hello world", "Execute this command output say hello message")
	})

	fmt.Printf("%+v\n", suite.Command.command)
}
