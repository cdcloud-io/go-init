# go-init

GO module/project initialization and styleguide for CDCLOUD

## RUN

```bash
curl -v https://raw.githubusercontent.com/cdcloud-io/go-init/refs/heads/develop/go_init_cli_remote.sh | bash
```

## TODO

### Makefile

- autoversioning GO code with Makefile

To set the version dynamically during the build process, you can use the -ldflags option in go build within your Makefile to inject version information. Here’s a step-by-step guide on how to set this up.

Add a Version Variable in Your Go Code: Define a variable for the version in your Go code that can be set at build time.

```go
// main.go
package main

import "fmt"

// Version is the version of the application, set during build time
var Version = "dev" // default to "dev" or any placeholder

func main() {
    fmt.Printf("Application Version: %s\n", Version)
}
```

```makefile
# Define the version; this can also be pulled from a git tag or passed in
VERSION := $(shell git describe --tags --always)

# Define the targets for your executables
TARGETS := executable1 executable2

all: $(TARGETS)

# Compile each target with the version set in ldflags
$(TARGETS):
    go build -o $@ -ldflags="-X 'main.Version=$(VERSION)'" ./cmd/$@

clean:
    rm -f $(TARGETS)

```
