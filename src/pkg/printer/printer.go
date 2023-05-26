package printer

import (
	"fmt"
	"io"
	"oduludo.io/eol/pkg/datasource"
)

func PrintVersionsList(out io.Writer, cycleList []datasource.ListedCycleDetail) error {
	for _, cycleDetail := range cycleList {
		if _, err := fmt.Fprintln(out, cycleDetail.Cycle); err != nil {
			return err
		}
	}

	return nil
}
