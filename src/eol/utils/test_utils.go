package utils

import (
	"fmt"
	"os"
	"strings"
)

func FileContainsJson(file string) (error, bool) {
	data, err := os.ReadFile(file)

	if err != nil {
		return err, false
	}

	bareText := strings.TrimSpace(string(data))
	firstChar := fmt.Sprintf("%c", bareText[0])
	lastChar := fmt.Sprintf("%c", bareText[len(bareText)-1])

	if firstChar != "{" {
		return nil, false
	}

	if lastChar != "}" {
		return nil, false
	}

	return nil, true
}

func FileIsEncrypted(file string) (error, bool) {
	encryptedData, err := os.ReadFile(file)

	if err != nil {
		return err, false
	}

	encryptedText := strings.TrimSpace(string(encryptedData))
	firstChar := fmt.Sprintf("%c", encryptedText[0])
	lastChar := fmt.Sprintf("%c", encryptedText[len(encryptedText)-1])

	// The encrypted data should not contain any new lines and should not begin or end with '{' and '}' respectively
	if strings.Contains(string(encryptedData), "\n") {
		return nil, false
	}

	if firstChar == "{" {
		return nil, false
	}

	if lastChar == "}" {
		return nil, false
	}

	return nil, true
}
