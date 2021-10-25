
![Logo](assets/banner/cobrax.jpg)

    
## Installation

Install cobrax with go module.

```bash
 go get github.com/coolstina/cobrax
```
    
## Features

- Use option design model definition command.
- Use callback function addition flags.
- Fast build client application.

## Usage/Examples

### Step1: Build client application

```go
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
```

### Step2: Install client application

```shell script
go install ./example
```

### Step3: Run client application on terminal

View help information:

```shell script
$ example -h
Execute this command output say hello

Usage:
  hello [flags]

Flags:
  -h, --help   help for hello
``` 

Execute command: 

```shell script
$ example hello
hello world
```

  
## Authors

- [@helloshaohua](https://www.github.com/helloshaohua)
  
## Feedback

If you have any feedback, please reach out to us at fake@fake.com

  
## License

[Apache](http://www.apache.org/licenses/LICENSE-2.0)