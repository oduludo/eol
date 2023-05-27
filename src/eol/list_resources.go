package eol

import (
	"github.com/spf13/cobra"
	"io"
	"oduludo.io/eol/pkg/argutils"
	"oduludo.io/eol/pkg/datasource"
	"oduludo.io/eol/pkg/printer"
	"strings"
)

var listResourcesDesc = "List the resources available in the datasource."

func newListResourcesCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-resources",
		Short: "List the resources available in the datasource.",
		Long:  makeLongUsageDescription(listResourcesDesc),
		Args:  argutils.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			contains, err := cmd.Flags().GetString("contains")

			if err != nil {
				return err
			}

			return runListResources(out, contains)
		},
	}

	cmd.PersistentFlags().String("contains", "", "filtering on resources that contain the provided substring")

	return cmd
}

func runListResources(out io.Writer, contains string) error {
	client := datasource.NewCycleClient()
	allResources, err := client.Resources()

	if err != nil {
		return err
	}

	var resources []string

	// Apply filtering if `contains` has been set
	if contains != "" {
		for _, resource := range allResources {
			if strings.Contains(resource, contains) {
				resources = append(resources, resource)
			}
		}
	} else {
		resources = allResources
	}

	if err := printer.PrintResourcesList(out, resources); err != nil {
		return err
	}

	return nil
}
