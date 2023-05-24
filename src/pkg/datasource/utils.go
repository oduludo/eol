package datasource

import "fmt"

const apiBase = "https://endoflife.date/api"

func constructCycleDetailUrl(resource string, version string) string {
	return fmt.Sprintf("%s/%s/%s.json", apiBase, resource, version)
}
