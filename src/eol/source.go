package eol

import (
	"github.com/spf13/cobra"
	"io"
)

var sourceDesc = "..."

func newSourceCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "source encrypt|decrypt|verify",
		Short: "encrypt|decrypt|verify",
		Long:  makeLongUsageDescription(sourceDesc),
		Args:  cobra.NoArgs,
	}

	cmd.AddCommand(
		newEncryptCmd(out),
		newDecryptCmd(out),
		newVerifyCmd(out),
	)

	//cmd.PersistentFlags().Bool("e", false, "set to true to give a non-zero exit code on EOL")

	return cmd
}
