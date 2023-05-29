package eol

import (
	"errors"
	"github.com/spf13/cobra"
	"io"
	"oduludo.io/eol/pkg/crypt"
	"os"
)

var decryptDesc = "..."

func newDecryptCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "decrypt FILE",
		Short: "Decrypt a file",
		Long:  makeLongUsageDescription(decryptDesc),
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			filePath := args[0]

			key, err := cmd.Flags().GetString("key")

			if err != nil {
				return err
			}

			if !crypt.ValidateKey(key) {
				return errors.New("key is invalid")
			}

			to, err := cmd.Flags().GetString("to")

			if err != nil {
				return err
			}

			return runSourceDecrypt(out, filePath, key, to)
		},
	}

	cmd.PersistentFlags().String("key", "", "configure a key to use for decryption")
	cmd.PersistentFlags().String("to", "", "location to write decrypted data to")

	return cmd
}

func runSourceDecrypt(out io.Writer, file string, key string, to string) error {
	encrypted, err := os.ReadFile(file)

	if err != nil {
		return err
	}

	// TODO: Check if key is not empty!
	decrypted, err := crypt.Decrypt(string(encrypted), key)

	if err != nil {
		return err
	}

	var outputLocation string

	if to == "" {
		outputLocation = file
	} else {
		outputLocation = to
	}

	if err := os.WriteFile(outputLocation, []byte(decrypted), 0644); err != nil {
		return err
	}

	return nil
}
