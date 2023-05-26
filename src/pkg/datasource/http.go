package datasource

import (
	"errors"
	"io"
	"net/http"
)

// CycleClient gets CycleDetail objects and implements the ClientInterface from oduludo.io/pkg/datasource.
type CycleClient[T CycleDetail, L ListedCycleDetail] struct{}

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

// MockCycleClient is used to mock the client/API during testing.
type MockCycleClient[T CycleDetail, L ListedCycleDetail] struct{}

func (c MockCycleClient[T, L]) Get(_ ...string) (T, error, bool) {
	data := loadMockData("cycle_detail.json")

	res := T{}

	if err := newObjectFromBytes(&res, data); err != nil {
		return res, err, false
	}

	return res, nil, false
}

func (c MockCycleClient[T, L]) All(_ ...string) ([]L, error, bool) {
	data := loadMockData("cycle_list.json")

	res := make([]L, 0)

	if err := newObjectFromBytes(&res, data); err != nil {
		return res, err, false
	}

	return res, nil, false
}
