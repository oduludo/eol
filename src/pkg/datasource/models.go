package datasource

import (
	"log"
	"strings"
	"time"
)

type CycleDetail struct {
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

func (cd CycleDetail) HasPassedEol() bool {
	return cd.EolTime().Before(time.Now())
}
