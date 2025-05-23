package main

// https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/

// ---------------- main: main.go ----------------------
// # Package Documentation
// - Package Name: main
// - Description:  Entry point for the CLI application. Sets up a base context
//   and passes OS-level dependencies to the app runner.
//
// ## Package File(s)
//
// ### Current File
// - main.go: (current_file)
//
// #### Function(s)
// - main: Creates base context, calls run.App, and handles exit logic.
//
// **Developer Note(s)**
// - Keeps `main()` minimal for testability.
// - Delegates full control to internal/app/run.go.
// -----------------------------------------------------------------------------

import (
	"context"
	"fmt"
	"os"

	"{{module_name}}/internal/{{app}}"
)

func main() {
	// create a fresh context (empty backpack)
	ctx := context.Background()

	// delegate to run.App, passing all OS-level dependencies
	if err := run.App(ctx, os.Args, os.Getenv, os.Stdin, os.Stdout, os.Stderr); err != nil {
		// if an error is returned, print to stderr and exit with code 1
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
