package datasource

import (
	"errors"
	"io"
	"net/http"
)

func GetCycleDetail(resource string, version string) (CycleDetail, error) {
	resp, err := http.Get(constructCycleDetailUrl(resource, version))
	if err != nil {
		return CycleDetail{}, err
	}

	if resp.StatusCode == 404 {
		return CycleDetail{}, errors.New("failed to find resource with specified version")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return CycleDetail{}, err
	}

	res, err := newCycleDetailFromBytes(body)

	if err != nil {
		return CycleDetail{}, err
	}

	return res, nil
}
