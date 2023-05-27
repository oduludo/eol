package eol

import (
	"github.com/spf13/cobra"
	"io"
	"oduludo.io/eol/pkg/argutils"
	"oduludo.io/eol/pkg/datasource"
	"oduludo.io/eol/pkg/printer"
)

var listVersionsDesc = "List the available versions for the specified resource."

func newListVersionsCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-versions RESOURCE",
		Short: "List the available versions for the specified resource.",
		Long:  makeLongUsageDescription(listVersionsDesc),
		Args:  argutils.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runListVersions(out, args[0])
		},
	}

	return cmd
}

func runListVersions(out io.Writer, resource string) error {
	client := datasource.NewCycleClient()
	result, err, _ := client.All(resource)

	if err != nil {
		return err
	}

	if err := printer.PrintVersionsList(out, resource, result); err != nil {
		return err
	}

	return nil
}
