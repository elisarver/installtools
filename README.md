# installtools
Code to install tools in the tools pattern

## What is the "tools" pattern?

Go has an informal pattern of naming repository targets for tools to install in a project, used in combination with go.mod to handle dependency version.

It looks like this inside the file:

```go
//go:build tools
// +build tools

package tools

import (
	_ "github.com/go-swagger/go-swagger/cmd/swagger"
	_ "github.com/maxbrunsfeld/counterfeiter/v6"
)
```

Note the `//go:build tools`/`// +build tools` directive(s), which makes the compiler skip this particular file.

## How do I do the install?

You could use a sed script to do some go modding, but you happen to have a programming language at your disposal, so let's make a main and run it!
