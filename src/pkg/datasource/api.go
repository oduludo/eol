package datasource

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func GetCycleDetail(resource string, version string) (CycleDetail, error) {
	cycleDetail := CycleDetail{}

	resp, err := http.Get(constructCycleDetailUrl(resource, version))
	if err != nil {
		return cycleDetail, err
	}

	if resp.StatusCode == 404 {
		return cycleDetail, errors.New("failed to find resource with specified version")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return cycleDetail, err
	}

	if err := json.Unmarshal(body, &cycleDetail); err != nil {
		return cycleDetail, err
	}

	return cycleDetail, nil
}
