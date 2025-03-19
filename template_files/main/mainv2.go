package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"time"

	// "<module>/internal/app"
)

/*
main, in GO currently does not return. Therefore, we don't have an application exit code.
The run function is like the main function, except that it takes in
operating system fundamentals as arguments and returns an error, like `int main()` in C.

---------------------------------------------------------------------------------------------------------
    VALUE       |    TYPE                |    DESCRIPTION
---------------------------------------------------------------------------------------------------------
   os.Args      | []string               |  The arguments passed when executing your program.
   os.Stdin     | io.Reader              |  For reading input.
   os.Stdout    | io.Writer              |  For writing output.
   os.Stderr    | io.Writer              |  For writing error logs.
   os.Getenv    | func(string) string    |  For reading environment variables.
   os.Getwd     | func() (string, error) |  Get the working directory.
---------------------------------------------------------------------------------------------------------
*/

func run(ctx context.Context, w io.Writer, args []string) error {
	// Create a timeout context first.
	ctx, timeoutCancel := context.WithTimeout(ctx, 5*time.Minute)
	defer timeoutCancel()

	// Wrap it with signal handling (e.g., Ctrl+C).
	ctx, signalCancel := signal.NotifyContext(ctx, os.Interrupt)
	defer signalCancel()

	// Run the application logic.
	if err := app.RunCli(ctx, w, args); err != nil {
		// Provide more context if it was a timeout.
		if ctx.Err() == context.DeadlineExceeded {
			return fmt.Errorf("operation timed out after 5 minutes: %w", err)
		}
		return err
	}

	return nil
}

func main() {
	// Root context, typically background for CLI apps.
	ctx := context.Background()

	// Run the main logic and handle any errors.
	if err := run(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}
