package eol

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
)

var datasourceCredits = "EOL uses https://endoflife.date/ as a data source."
var globalUsage = "CLI application to check whether a tool's version has reached EOL."

func makeLongUsageDescription(desc string) string {
	return fmt.Sprintf("%s\n\n%s", desc, datasourceCredits)
}

func NewRootCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "eol",
		Short:        "CLI application to check whether a tool's version has reached EOL.",
		Long:         makeLongUsageDescription(globalUsage),
		SilenceUsage: true,
	}

	cmd.AddCommand(
		newCheckCmd(out),
	)

	return cmd
}
