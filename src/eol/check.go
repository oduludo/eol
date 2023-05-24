package eol

import (
	"github.com/spf13/cobra"
	"io"
	"oduludo.io/eol/pkg/argutils"
	"oduludo.io/eol/pkg/datasource"
	"strconv"
)

var checkDesc = "Check the EOL status for a resource's version.\n" +
	"Version formatting differs per resource and follows the datasource's API convention."

func newCheckCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "check RESOURCE VERSION",
		Short: "Check the EOL status for a resource's version.",
		Long:  makeLongUsageDescription(checkDesc),
		Args:  argutils.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(out, args[0], args[1])
		},
	}

	return cmd
}

func run(out io.Writer, resource string, version string) error {
	client := datasource.CycleDetailClient[datasource.CycleDetail]{}
	cycleDetail, err := client.Get(resource, version)

	if err != nil {
		return err
	}

	eolResultStr := strconv.FormatBool(cycleDetail.HasPassedEol())

	if _, err := io.WriteString(out, eolResultStr); err != nil {
		return err
	}

	return nil
}
