package argutils

import (
	"fmt"
	"github.com/spf13/cobra"
)

// pluralize comes from the Helm project.
func pluralize(word string, n int) string {
	if n == 1 {
		return word
	}
	return word + "s"
}

// ExactArgs comes from the Helm project.
func ExactArgs(n int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) != n {
			return fmt.Errorf(
				"%q requires %d %s\n\nUsage:  %s",
				cmd.CommandPath(),
				n,
				pluralize("argument", n),
				cmd.UseLine(),
			)
		}
		return nil
	}
}
