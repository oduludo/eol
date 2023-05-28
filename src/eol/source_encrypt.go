package eol

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"oduludo.io/eol/pkg/crypt"
	"os"
)

var encryptDesc = "..."

func newEncryptCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "encrypt FILE",
		Short: "Encrypt a file",
		Long:  makeLongUsageDescription(encryptDesc),
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			filePath := args[0]

			key, err := cmd.Flags().GetString("key")

			if err != nil {
				return err
			}

			to, err := cmd.Flags().GetString("to")

			if err != nil {
				return err
			}

			return runSourceEncrypt(out, filePath, key, to)
		},
	}

	cmd.PersistentFlags().String("key", "", "optionally configure a key to use for encryption")
	cmd.PersistentFlags().String("to", "", "location to write encrypted data to")

	return cmd
}

func runSourceEncrypt(out io.Writer, file string, key string, to string) error {
	withNewKey := false
	data, err := os.ReadFile(file)

	if err != nil {
		return err
	}

	if key == "" {
		key = crypt.GenerateKey()
		withNewKey = true
	}

	encrypted, err := crypt.Encrypt(string(data), key)

	if err != nil {
		return err
	}

	var outputLocation string

	if to == "" {
		outputLocation = file
	} else {
		outputLocation = to
	}

	if err := os.WriteFile(outputLocation, []byte(encrypted), 0644); err != nil {
		return err
	}

	// Print the key if it had to be newly generated
	if withNewKey {
		fmt.Fprintf(out, "Generated new encryption key:\n%s\nStore it well!\n", key)
	}

	return nil
}
