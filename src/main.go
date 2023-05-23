package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const ApiBase = "https://endoflife.date/api"

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

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: eol <resource> <version>\nE.g.: eol ruby 2.7")
		os.Exit(1)
	}

	resource := os.Args[1]
	version := os.Args[2]

	if strings.Count(version, ".") > 1 {
		fmt.Println("Please set a version according to <major>.<minor>")
		os.Exit(1)
	}

	url := fmt.Sprintf("%s/%s/%s.json", ApiBase, resource, version)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode == 404 {
		log.Fatalln(errors.New("failed to find resource with specified version"))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	cycleDetail := CycleDetail{}
	if err := json.Unmarshal(body, &cycleDetail); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(strconv.FormatBool(cycleDetail.HasPassedEol()))
}
