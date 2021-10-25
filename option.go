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

import "github.com/spf13/cobra"

type Option interface {
	apply(*options)
}

type RunFunc func(cmd *cobra.Command, args []string)

type options struct {
	use   string  // Command
	short string  // Short information.
	long  string  // Long information.
	run   RunFunc // Run command handle
}

type optionFunc func(*options)

func (o optionFunc) apply(ops *options) {
	o(ops)
}

// WithUse Use specific parameter for use value override original value.
func WithUse(use string) Option {
	return optionFunc(func(ops *options) {
		ops.use = use
	})
}

// WithShort Use specific parameter for short value override original value.
func WithShort(short string) Option {
	return optionFunc(func(ops *options) {
		ops.short = short
	})
}

// WithLong Use specific parameter for long value override original value.
func WithLong(long string) Option {
	return optionFunc(func(ops *options) {
		ops.long = long
	})
}

// WithRun Use specific parameter for run value override original value.
func WithRun(run RunFunc) Option {
	return optionFunc(func(ops *options) {
		ops.run = run
	})
}
