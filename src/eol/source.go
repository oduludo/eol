package eol

import (
	"github.com/spf13/cobra"
	"io"
)

var sourceDesc = "..."

func newSourceCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "source encrypt|decrypt",
		Short: "encrypt|decrypt",
		Long:  makeLongUsageDescription(sourceDesc),
		Args:  cobra.NoArgs,
	}

	cmd.AddCommand(
		newEncryptCmd(out),
		newDecryptCmd(out),
	)

	//cmd.PersistentFlags().Bool("e", false, "set to true to give a non-zero exit code on EOL")

	return cmd
}
