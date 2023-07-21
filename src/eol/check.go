package eol

import (
	"errors"
	"github.com/spf13/cobra"
	"io"
	"oduludo.io/eol/cfg"
	"oduludo.io/eol/eol/utils"
	"oduludo.io/eol/pkg/argutils"
	"oduludo.io/eol/pkg/datasource"
	"oduludo.io/eol/pkg/printer"
	"os"
	"strconv"
	"strings"
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

			source, err := cmd.Flags().GetString("source")

			if err != nil {
				return err
			}

			xsource, err := cmd.Flags().GetString("xsource")

			if err != nil {
				return err
			}

			// Only allow one flag of '--source' and '--xsource'
			if source != "" && xsource != "" {
				return errors.New(cfg.SourceXsourceXorMsg)
			}

			// Unpack the actual source URLs to use
			var sources []string
			useRootSource := true

			if source != "" {
				sources = strings.Split(source, ",")
			} else if xsource != "" {
				sources = strings.Split(xsource, ",")
				useRootSource = false
			} else {
				// Run only on the root source
				return runCheckOnRootSource(out, args[0], args[1], exitWithCode)
			}

			// Unpack keys and verify if they match the number of URLs
			key, err := cmd.Flags().GetString("key")

			if err != nil {
				return err
			}

			var keys []string

			if key == "" {
				// Create an array of empty keys with same length as number of sources
				keys = make([]string, len(sources))
			} else {
				rawKeys := strings.Split(key, ",")

				for _, rawKey := range rawKeys {
					keys = append(keys, strings.TrimSpace(rawKey))
				}
			}

			// keys is allowed to be empty, meaning none of the sources need decryption
			// in such a situation, the length mismatch between sources and keys is not an issue
			if len(keys) != 0 && len(sources) != len(keys) {
				return errors.New(cfg.InvalidKeysNumMsg)
			}

			// Run the appropriate command implementation
			if useRootSource {
				return runCheckWithRootSourceAndCustomSources(out, args[0], args[1], exitWithCode, sources, keys)
			} else {
				err, _ := runCheckWithCustomSourcesOnly(out, args[0], args[1], exitWithCode, sources, keys)
				return err
			}
		},
	}

	cmd.PersistentFlags().Bool("e", false, "set to true to give a non-zero exit code on EOL")
	cmd.PersistentFlags().String("source", "", "combine the root source with one or more URLs pointing to custom datasource resources, delimiting URLs using ','")
	cmd.PersistentFlags().String("xsource", "", "exclusively use one or more URLs pointing to custom datasource resources, delimiting URLs using ',', not using any of the root data")
	cmd.PersistentFlags().String("key", "", "one or more keys to decrypt custom sources, delimiting keys using ','")

	return cmd
}

func runCheckOnRootSource(out io.Writer, resource, version string, exitWithCode bool) error {
	client := datasource.NewCycleClient()
	cycleDetail, err, notFound := client.Get(resource, version)

	if notFound {
		// Requested version not found, show available versions instead
		result, err, _ := client.All(resource)

		if err != nil {
			return err
		}

		if err := printer.PrintVersionsList(out, resource, result); err != nil {
			return err
		}
	} else if err != nil {
		return err
	} else {
		// Continue with the command's logic
		hasPassedEol := cycleDetail.HasPassedEol()
		eolResultStr := strconv.FormatBool(hasPassedEol)

		if _, err := io.WriteString(out, eolResultStr); err != nil {
			return err
		}

		// Explicitly exit with code if user has indicated to give a non-zero status code in case EOL has been reached.
		if exitWithCode && hasPassedEol {
			os.Exit(eolExitCode)
		}
	}

	return nil
}

// Run a check on both the root source and the provided custom sources.
// If the custom datasource holds the resource+version, use that to determine EOL
// If the custom datasource does not hold the resource+version (both if the resource itself is present, but not the version and if the resource isn't present at all),
// continue to check on the root source.
func runCheckWithRootSourceAndCustomSources(out io.Writer, resource, version string, exitWithCode bool, customSources, keys []string) error {
	// Check on custom sources, returning an error if present
	// If error is nil, continue to checking on the root
	err, hasCompleted := runCheckWithCustomSourcesOnly(out, resource, version, exitWithCode, customSources, keys)

	if err != nil {
		return err
	}

	if hasCompleted {
		return nil
	}

	// Check on the root source if the custom sources yield no match
	return runCheckOnRootSource(out, resource, version, exitWithCode)
}

func runCheckWithCustomSourcesOnly(out io.Writer, resource, version string, exitWithCode bool, customSources, keys []string) (error, bool) {
	client := datasource.NewCycleClient()

	sourceAndKeyPairs, ok := utils.Zip[string, string](customSources, keys)

	if !ok {
		return errors.New(cfg.ZipLenMismatchMsg), false
	}

	for _, sourceAndKeyPair := range sourceAndKeyPairs {
		source := sourceAndKeyPair.A.(string)
		key := sourceAndKeyPair.B.(string)

		cycleDetail, err, notFound := client.GetCustom(source, resource, version, key)

		if err != nil && !notFound {
			return err, false
		}

		// Source did not contain the resource+version. Move on to the next source
		if notFound {
			continue
		}

		// Found the requested resource+version; check the EOL and report
		hasPassedEol := cycleDetail.HasPassedEol()
		eolResultStr := strconv.FormatBool(hasPassedEol)

		if _, err := io.WriteString(out, eolResultStr); err != nil {
			return err, true
		}

		// Explicitly exit with code if user has indicated to give a non-zero status code in case EOL has been reached.
		if exitWithCode && hasPassedEol {
			os.Exit(eolExitCode)
		} else {
			// Essentially a substitute for a 0 exit code, to not break the testing cycle
			return nil, true
		}
	}

	// No errors, but no resource+version matches either
	return nil, false
}
