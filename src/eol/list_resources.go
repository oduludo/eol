package eol

import (
	"github.com/spf13/cobra"
	"io"
	"oduludo.io/eol/pkg/argutils"
	"oduludo.io/eol/pkg/datasource"
	"oduludo.io/eol/pkg/printer"
)

var listResourcesDesc = "List the resources available in the datasource."

func newListResourcesCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-resources",
		Short: "List the resources available in the datasource.",
		Long:  makeLongUsageDescription(listResourcesDesc),
		Args:  argutils.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runListResources(out)
		},
	}

	return cmd
}

func runListResources(out io.Writer) error {
	client := datasource.NewCycleClient()
	resources, err := client.Resources()

	if err != nil {
		return err
	}

	if err := printer.PrintResourcesList(out, resources); err != nil {
		return err
	}

	return nil
}
