package datasource

import (
	"fmt"
	"log"
	"os"
)

const apiBase = "https://endoflife.date/api"

func constructCycleDetailUrl(resource string, version string) string {
	return fmt.Sprintf("%s/%s/%s.json", apiBase, resource, version)
}

func constructCycleListUrl(resource string) string {
	return fmt.Sprintf("%s/%s.json", apiBase, resource)
}

func loadMockData(file string) []byte {
	data, err := os.ReadFile(fmt.Sprintf("./test_assets/%s", file))

	if err != nil {
		log.Fatal(err)
	}

	return data
}
