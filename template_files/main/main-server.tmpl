package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
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

NOTE: 
This version is designed for **long-running services**.
*/

// run initializes and starts the main application logic.
// It takes a context (ctx), an io.Writer (w), and an argument list (args).
// This function wraps the application's execution with signal handling for graceful shutdown.
//
// Parameters:
//   ctx  - the base context for the application lifecycle
//   w    - a writer (typically os.Stdout) for normal application output
//   args - a slice of strings representing command-line arguments passed to the program
//
// Returns:
//   error - if an error occurs during execution of the application, it will be returned
func run(ctx context.Context, w io.Writer, args []string) error {
	// Listen for Ctrl+C / SIGTERM.
	// signal.NotifyContext returns a new context that is canceled when an OS interrupt signal is received.
	// It allows the program to clean up and shut down gracefully when interrupted.
	ctx, signalCancel := signal.NotifyContext(ctx, os.Interrupt)
	defer signalCancel() // Ensure resources tied to signal notification are released when run() exits.

	// Start your long-running service (no timeout).
	// This will block until either the service stops or an interrupt signal is received.
	// It uses the provided context to handle cancellation and graceful shutdown.
	if err := app.RunServer(ctx, w, args); err != nil {
		// If an error occurs in the RunServer function, propagate it to the caller.
		return err
	}

	// Return nil to indicate successful execution without errors.
	return nil
}

// main is the entry point of the Go application.
// It sets up the base context, invokes the run() function, and handles any errors returned.
//
// Behavior:
//   - Creates a root context using context.Background()
//   - Calls run() to start the application logic
//   - If run() returns an error, it writes the error to stderr and exits the program with a non-zero status code
func main() {
	// Create a root context for the application.
	// This context is typically never canceled directly, but acts as the parent for derived contexts.
	ctx := context.Background()

	// Call the run() function with the root context, standard output, and command-line arguments.
	// If an error is returned, write it to standard error and exit with status code 1.
	if err := run(ctx, os.Stdout, os.Args); err != nil {
		// Write the error message to os.Stderr.
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)

		// Exit the application with a non-zero exit code to indicate failure.
		os.Exit(1)
	}
}
