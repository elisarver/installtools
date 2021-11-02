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
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint@v1.42.0"
)
```

Note the `//go:build tools`/`// +build tools` directive(s), which makes the compiler skip this particular file. The error of including the specific version at this point doesn't matter; only the textual representation does for the install tool to use it.

## How do I install?

First, `go install github.com/elisarver/installtools/cmd/gotoolinstaller` to get the tool and then run `gotoolinstaller | sh` to install the packages.

You can try it in the example directory, which installs a specific version of golangci-lint that you can verify with --version, post-install.