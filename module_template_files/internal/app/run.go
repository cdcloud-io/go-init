package run

// ---------------- run: run.go ----------------------
// # Package Documentation
// - Package Name: run
// - Description:  Application entry point logic for CLI or service startup.
//
// ## Package File(s)
//
// ### Current File
// - run.go: (current_file)
//
// #### Function(s)
// - App: Application runtime entrypoint, with injected OS-level dependencies.
//
// **Developer Note(s)**
// - All real work happens here. Useful for tests by injecting mocked args/env/stdin/stdout.
// -----------------------------------------------------------------------------

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"time"
)

// ================ PUBLIC FUNCTION(s) =========================================

// App runs the CLI application with full dependency injection
func App(
	ctx context.Context,
	args []string,
	getenv func(string) string,
	stdin io.Reader,
	stdout io.Writer,
	stderr io.Writer,
) error {

	// wrap the context with interrupt signal awareness
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	// create a FlagSet for parsing CLI flags (for testability)
	fs := flag.NewFlagSet(args[0], flag.ContinueOnError)
	var (
		outFile = fs.String("out", "default.txt", "output file path")
		format  = fs.String("fmt", "text", "output format")
	)

	// direct flag errors to stderr instead of stdout
	fs.SetOutput(stderr)

	// parse the remaining args (skipping args[0], the program name)
	if err := fs.Parse(args[1:]); err != nil {
		return fmt.Errorf("parsing flags: %w", err)
	}

	// simulate doing work, respecting the context
	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintf(stdout, "🟩 Output file: %s | Format: %s\n", *outFile, *format)
	case <-ctx.Done():
		return errors.New("🟥 interrupted by user")
	}

	return nil
}
