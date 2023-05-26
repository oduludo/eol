package datasource

import (
	"encoding/json"
	"log"
	"strings"
	"time"
)

type CommonUnderlyingType interface{}

type CycleDetail struct {
	CommonUnderlyingType
	Eol               string `json:"eol"`
	Latest            string `json:"latest"`
	LatestReleaseDate string `json:"latestReleaseDate"`
	ReleaseDate       string `json:"releaseDate"`
	Lts               bool   `json:"lts"`
}

func (cd CycleDetail) EolTime() time.Time {
	eolTime, err := time.Parse(time.DateOnly, strings.Replace(cd.Eol, "/", "-", 2))

	if err != nil {
		log.Fatalln(err)
	}

	return eolTime
}

// Extract hasPassedEolWithTime's logic from HasPassedEol for better testability.
// EOL is defined as the time's date being before the current date. Equal dates are not flagged as EOL.
func hasPassedEolWithTime(cd *CycleDetail, now time.Time) bool {
	return cd.EolTime().Before(now)
}

func (cd CycleDetail) HasPassedEol() bool {
	return hasPassedEolWithTime(&cd, time.Now())
}

// Parse bytes into a CycleDetail object using newCycleDetailFromBytes.
func newCycleDetailFromBytes(data []byte) (CycleDetail, error) {
	cycleDetail := CycleDetail{}

	if err := json.Unmarshal(data, &cycleDetail); err != nil {
		return cycleDetail, err
	}

	return cycleDetail, nil
}

func newObjectFromBytes[T CommonUnderlyingType](ref *T, data []byte) error {
	return json.Unmarshal(data, &ref)
}

type ListedCycleDetail struct {
	CycleDetail
	Cycle string `json:"cycle"`
}
