package datasource

import (
	"encoding/json"
	"errors"
	"flag"
	"io"
	"net/http"
	"oduludo.io/eol/cfg"
	pkghttp "oduludo.io/eol/pkg/http"
	"os"
)

// CycleClient gets CycleDetail objects and implements the ClientInterface from oduludo.io/pkg/datasource.
type CycleClient[T CycleDetail, L ListedCycleDetail] struct{}

func (c CycleClient[T, L]) Resources() ([]string, error) {
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

func (c CycleClient[T, L]) Get(args ...string) (T, error, bool) {
	resp, err := http.Get(constructCycleDetailUrl(args[0], args[1]))
	res := T{}

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

func (c CycleClient[T, L]) All(args ...string) ([]L, error, bool) {
	resp, err := http.Get(constructCycleListUrl(args[0]))
	res := make([]L, 0)

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

// MockCycleClient is used to mock the client/API during testing.
type MockCycleClient[T CycleDetail, L ListedCycleDetail] struct{}

func (c MockCycleClient[T, L]) Resources() ([]string, error) {
	data := loadMockData(mockAll)

	res := make([]string, 0)

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c MockCycleClient[T, L]) Get(args ...string) (T, error, bool) {
	resource := args[0]
	version := args[1]

	if resource != "ruby" {
		return T{}, errors.New("unsupported resource under testing"), false
	}

	var file string

	if version == "3.2" {
		file = mockResourceBeforeEol
	} else if version == "2.7" {
		file = mockResourcePassedEol
	} else {
		return T{}, errors.New("failed to find resource"), true
	}

	data := loadMockData(file)

	res := T{}

	if err := newObjectFromBytes(&res, data); err != nil {
		return res, err, false
	}

	return res, nil, false
}

func (c MockCycleClient[T, L]) All(args ...string) ([]L, error, bool) {
	var file string

	if args[0] == "ruby" {
		file = mockResourceAll
	} else {
		return make([]L, 0), nil, true
	}

	data := loadMockData(file)

	res := make([]L, 0)

	if err := newObjectFromBytes(&res, data); err != nil {
		return res, err, false
	}

	return res, nil, false
}

func NewCycleClient() pkghttp.ClientInterface[CycleDetail, ListedCycleDetail] {
	isIntegrationTestStr := os.Getenv(cfg.IsIntegrationTestEnvKey)
	isIntegrationTest := isIntegrationTestStr == "true"

	if flag.Lookup("test.v") != nil && !isIntegrationTest {
		return MockCycleClient[CycleDetail, ListedCycleDetail]{}
	}

	return CycleClient[CycleDetail, ListedCycleDetail]{}
}
