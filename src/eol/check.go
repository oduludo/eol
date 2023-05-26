package eol

import (
	"github.com/spf13/cobra"
	"io"
	"oduludo.io/eol/pkg/argutils"
	"oduludo.io/eol/pkg/datasource"
	"oduludo.io/eol/pkg/printer"
	"os"
	"strconv"
)

const eolExitCode = 69

var checkDesc = "Check the EOL status for a resource's version.\n" +
	"Version formatting differs per resource and follows the datasource's API convention.\n" +
	"Run with --e to ensure a non-zero exit code if EOL has been reached."

func newCheckCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "check RESOURCE VERSION",
		Short: "Check the EOL status for a resource's version.",
		Long:  makeLongUsageDescription(checkDesc),
		Args:  argutils.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			exitWithCode, err := cmd.Flags().GetBool("e")

			if err != nil {
				return err
			}

			return run(out, args[0], args[1], exitWithCode)
		},
	}

	cmd.PersistentFlags().Bool("e", false, "set to true to give a non-zero exit code on EOL")

	return cmd
}

func run(out io.Writer, resource string, version string, exitWithCode bool) error {
	client := datasource.CycleClient[datasource.CycleDetail, datasource.ListedCycleDetail]{}
	cycleDetail, err, notFound := client.Get(resource, version)

	if notFound {
		res, err, _ := client.All(resource)

		if err != nil {
			return err
		}

		if err := printer.PrintVersionsList(out, res); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	hasPassedEol := cycleDetail.HasPassedEol()
	eolResultStr := strconv.FormatBool(hasPassedEol)

	if _, err := io.WriteString(out, eolResultStr); err != nil {
		return err
	}

	// Explicitly exit with code 1 if user has indicated to give a non-zero status code in case EOL has been reached.
	if exitWithCode && hasPassedEol {
		os.Exit(eolExitCode)
	}

	return nil
}
