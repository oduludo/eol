package datasource

import (
	"fmt"
	"log"
	"oduludo.io/eol/cfg"
	"os"
	"path"
)

const apiBase = "https://endoflife.date/api"

func constructCycleDetailUrl(resource string, version string) string {
	return fmt.Sprintf("%s/%s/%s.json", apiBase, resource, version)
}

func constructCycleListUrl(resource string) string {
	return fmt.Sprintf("%s/%s.json", apiBase, resource)
}

func constructResourcesUrl() string {
	return fmt.Sprintf("%s/all.json", apiBase)
}

func loadMockData(file string) []byte {
	data, err := os.ReadFile(path.Join(cfg.TestAssets, file))

	if err != nil {
		log.Fatal(err)
	}

	return data
}
