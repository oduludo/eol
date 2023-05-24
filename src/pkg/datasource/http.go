package datasource

import (
	"errors"
	"io"
	"net/http"
)

// CycleDetailClient gets CycleDetail objects and implements the ClientInterface from oduludo.io/pkg/datasource.
type CycleDetailClient[T CycleDetail] struct{}

func (c CycleDetailClient[T]) Get(args ...string) (T, error) {
	resp, err := http.Get(constructCycleDetailUrl(args[0], args[1]))

	if err != nil {
		return T{}, err
	}

	if resp.StatusCode == 404 {
		return T{}, errors.New("failed to find resource with specified version")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return T{}, err
	}

	res := T{}

	if newObjectFromBytes(&res, body) != nil {
		return T{}, err
	}

	return res, nil
}

// MockCycleDetailClient is used to mock the client/API during testing.
type MockCycleDetailClient[T CycleDetail] struct{}

func (c MockCycleDetailClient[T]) Get(args ...string) (T, error) {
	data := loadMockData("cycle_detail.json")

	res := T{}

	if err := newObjectFromBytes(&res, data); err != nil {
		return T{}, err
	}

	return res, nil
}
