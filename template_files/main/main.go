package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"

	// "<module>/internal/app"
)

/*
main, in GO currently does not return.  Therefore we dont have an application exit code.
The run function is like the main function, except that it takes in
operating system fundamentals as arguments, and returns an error.  like in c int main()
---------------------------------------------------------------------------------------------------------
    VALUE     |    TYPE                |    DESCRIPTION
---------------------------------------------------------------------------------------------------------
   os.Args	  | []string	             |  The arguments passed executing your program /or parsing flags.
   os.Stdin	  | io.Reader	             |  For reading input
   os.Stdout	| io.Writer	             |  For writing output
   os.Stderr	| io.Writer	             |  For writing error logs
   os.Getenv	| func(string)           |  string	For reading environment variables
   os.Getwd	  | func() (string, error) | 	Get the working directory
---------------------------------------------------------------------------------------------------------
*/

func run(ctx context.Context, w io.Writer, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	if err := app.RunApp(ctx, w, args); err != nil {
		return err
	}

	return nil
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
