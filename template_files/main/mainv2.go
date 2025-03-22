package main

// cdcloud-io proprietary main() entrypoint v2.0
// validated 3/21/25

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"time"
)

/*
main.go version 2
- Timeout handling: includes a 5-minute timeout context, preventing infinite hangs
- Better context handling: It properly layers contexts (timeout wrapped with signal handling)
- Better error messaging: It provides more detailed error context, especially for timeouts
- More comprehensive comments: It includes detailed documentation about the OS fundamentals


main() does not return an error code. Therefore, we do NOT
have an application exit code directly in main.

The run() function is like the main function but takes in
operating system input(s) as arguments and returns an error,
similar to `int main()` in C.

--------------------------------------------------------------------------------
    VALUE     |    TYPE                |    DESCRIPTION
--------------------------------------------------------------------------------
   os.Args	  | []string	             |  Arguments / flags passed on execution
   os.Stdin	  | io.Reader	             |  For reading input
   os.Stdout	| io.Writer	             |  For writing output
   os.Stderr	| io.Writer	             |  For writing error logs
   os.Getenv	| func(string)           |  For reading environment variables
   os.Getwd	  | func() (string, error) |  Gets the working directory
--------------------------------------------------------------------------------
note(s)
- signal.NotifyContext: listens for OS signals (e.g., SIGINT or SIGTERM) and
  cancels the context upon receiving them with graceful shutdown

--------------------------------------------------------------------------------
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

/*
// application entrypoint function

package application

func Run(ctx context.Context, w io.Writer, args []string) error {
	// create a context that we can cancel on Ctrl+C
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

  // code

}

*/
