package datasource

import (
	"encoding/json"
	"errors"
	"flag"
	"io"
	"net/http"
	"oduludo.io/eol/cfg"
	"oduludo.io/eol/pkg/crypt"
	"os"
)

type customDataset = map[string][]ListedCycleDetail

// getCustom fetches a ListedCycleDetail object for the given resource and version from a custom datasource.
//
// NOTE: This implementation works for both the real client and the mock client, as it allows us to mock the returned data by
// pointing to the desired resource on localhost:8000 under testing.
//
// The root datasource offers URL endpoints (or 404) for this, but custom datasources (which are 'flat' JSON data) do not.
// In this case, simply search through the complete JSON dataset and construct a ListedCycleDetail object on the fly.
func getCustom(source, resource, version, decryptionKey string) (ListedCycleDetail, error, bool) {
	// Fetch the full JSON dataset
	resp, err := http.Get(source)
	res := ListedCycleDetail{}

	if err != nil {
		return res, err, false
	}

	if resp.StatusCode == 404 {
		return res, errors.New("failed to find custom source"), true
	}

	// Read the full JSON dataset into a struct
	dataset := customDataset{}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, err, false
	}

	// Decrypt the body if necessary, and if the key is not empty or the placeholder
	if decryptionKey != "" && decryptionKey != cfg.DecryptionKeyPlaceholder {
		isEncrypted, err := crypt.StringIsEncrypted(string(body))

		if err != nil {
			return res, err, false
		}

		if isEncrypted {
			decrypted, err := crypt.Decrypt(string(body), decryptionKey)

			if err != nil {
				return res, err, false
			}

			body = []byte(decrypted)
		}
	}

	// Unmarshal data
	if err := newObjectFromBytes(&dataset, body); err != nil {
		return res, err, false
	}

	resourceData, ok := dataset[resource]

	if !ok {
		return res, errors.New("failed to find resource data in custom source"), true
	}

	for _, listedCycleDetail := range resourceData {
		if listedCycleDetail.Cycle == version {
			return listedCycleDetail, nil, false
		}
	}

	// None of the cycle versions matched the criteria
	return res, nil, true
}

// CycleClient gets CycleDetail objects and implements the ClientInterface from oduludo.io/pkg/datasource.
type CycleClient struct{}

func (c CycleClient) Resources() ([]string, error) {
	resp, err := http.Get(constructResourcesUrl())

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := make([]string, 0)

	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// Get a CycleDetail object for the given resource and version.
func (c CycleClient) Get(args ...string) (CycleDetail, error, bool) {
	resp, err := http.Get(constructCycleDetailUrl(args[0], args[1]))
	res := CycleDetail{}

	if err != nil {
		return res, err, false
	}

	if resp.StatusCode == 404 {
		return res, errors.New("failed to find resource with specified version"), true
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, err, false
	}

	if err := newObjectFromBytes(&res, body); err != nil {
		return res, err, false
	}

	return res, nil, false
}

func (c CycleClient) All(args ...string) ([]ListedCycleDetail, error, bool) {
	resp, err := http.Get(constructCycleListUrl(args[0]))
	res := make([]ListedCycleDetail, 0)

	if err != nil {
		return res, err, false
	}

	if resp.StatusCode == 404 {
		return res, errors.New("failed to find resource"), true
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, err, false
	}

	if err := newObjectFromBytes(&res, body); err != nil {
		return res, err, false
	}

	return res, nil, false
}

func (c CycleClient) GetCustom(source, resource, version, decryptionKey string) (ListedCycleDetail, error, bool) {
	return getCustom(source, resource, version, decryptionKey)
}

// MockCycleClient is used to mock the client/API during testing.
type MockCycleClient struct{}

func (c MockCycleClient) Resources() ([]string, error) {
	data := loadMockData(mockAll)

	res := make([]string, 0)

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c MockCycleClient) Get(args ...string) (CycleDetail, error, bool) {
	resource := args[0]
	version := args[1]

	if resource != "ruby" {
		return CycleDetail{}, errors.New("unsupported resource under testing"), false
	}

	var file string

	if version == "3.2" {
		file = mockResourceBeforeEol
	} else if version == "2.7" {
		file = mockResourcePassedEol
	} else {
		return CycleDetail{}, errors.New("failed to find resource"), true
	}

	data := loadMockData(file)

	res := CycleDetail{}

	if err := newObjectFromBytes(&res, data); err != nil {
		return res, err, false
	}

	return res, nil, false
}

func (c MockCycleClient) All(args ...string) ([]ListedCycleDetail, error, bool) {
	var file string

	if args[0] == "ruby" {
		file = mockResourceAll
	} else {
		return make([]ListedCycleDetail, 0), nil, true
	}

	data := loadMockData(file)

	res := make([]ListedCycleDetail, 0)

	if err := newObjectFromBytes(&res, data); err != nil {
		return res, err, false
	}

	return res, nil, false
}

func (c MockCycleClient) GetCustom(source, resource, version, decryptionKey string) (ListedCycleDetail, error, bool) {
	return getCustom(source, resource, version, decryptionKey)
}

func NewCycleClient() ClientInterface {
	isIntegrationTestStr := os.Getenv(cfg.IsIntegrationTestEnvKey)
	isIntegrationTest := isIntegrationTestStr == "true"

	if flag.Lookup("test.v") != nil && !isIntegrationTest {
		return MockCycleClient{}
	}

	return CycleClient{}
}
