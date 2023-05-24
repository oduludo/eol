package eol

import (
	"github.com/spf13/cobra"
	"io"
)

var globalUsage = "CLI application to check whether a tool's version has reached EOL."

func NewRootCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "eol",
		Short:        "CLI application to check whether a tool's version has reached EOL.",
		Long:         globalUsage,
		SilenceUsage: true,
	}

	cmd.AddCommand(
		newCheckCmd(out),
	)

	return cmd
}
