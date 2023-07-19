package eol

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xeipuuv/gojsonschema"
	"io"
	"oduludo.io/eol/cfg"
	"oduludo.io/eol/pkg/crypt"
	"os"
)

var verifyDesc = "..."

func newVerifyCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verify FILE",
		Short: "Verify a file",
		Long:  makeLongUsageDescription(verifyDesc),
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			filePath := args[0]

			key, err := cmd.Flags().GetString("key")

			if err != nil {
				return err
			}

			return runSourceVerify(out, filePath, key)
		},
	}

	cmd.PersistentFlags().String("key", "", "optionally configure a key to use for encryption")

	return cmd
}

func runSourceVerify(out io.Writer, file string, key string) error {
	data, err := os.ReadFile(file)

	if err != nil {
		return err
	}

	datasource := string(data)

	// Decrypt the datasource if a key was provided
	if key != "" {
		datasource, err = crypt.Decrypt(datasource, key)

		if err != nil {
			return err
		}
	}

	// Load schema and data and validate the datasource
	schemaLoader := gojsonschema.NewStringLoader(cfg.DatasourceSchema)
	documentLoader := gojsonschema.NewStringLoader(datasource)
	res, err := gojsonschema.Validate(schemaLoader, documentLoader)

	if err != nil {
		return err
	}

	if !res.Valid() {
		return errors.New(cfg.DatasourceInvalidMsg)
	}

	fmt.Fprintf(out, "%s\n", cfg.DatasourceValidMsg)

	return nil
}
