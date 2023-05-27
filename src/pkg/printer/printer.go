package printer

import (
	"fmt"
	"io"
	"oduludo.io/eol/pkg/datasource"
)

const (
	cycleNameTitle   = "Cycle name"
	releaseNameTitle = "Release name"
)

func PrintVersionsList(out io.Writer, resource string, cycleList []datasource.ListedCycleDetail) error {
	if _, err := fmt.Fprintf(out, "Available versions for resource `%s` are:\n", resource); err != nil {
		return err
	}

	cycleNames := make([]string, 0)
	releaseNames := make([]string, 0)

	for _, cycleDetail := range cycleList {
		cycleNames = append(cycleNames, cycleDetail.Cycle)
		releaseNames = append(releaseNames, cycleDetail.ReleaseLabel)
	}

	cycleNameMaxLen, err := maxStringLength(append(cycleNames, cycleNameTitle))

	if err != nil {
		return err
	}

	releaseNameMaxLen, err := maxStringLength(append(releaseNames, releaseNameTitle))

	if err != nil {
		return err
	}

	var table string

	table += fmt.Sprintf("+%s+%s+\n", nChar(cycleNameMaxLen+PADDING, '-'), nChar(releaseNameMaxLen+PADDING, '-'))
	table += fmt.Sprintf("|%s|%s|\n", centerString(cycleNameTitle, cycleNameMaxLen), centerString(releaseNameTitle, releaseNameMaxLen))
	table += fmt.Sprintf("+%s+%s+\n", nChar(cycleNameMaxLen+PADDING, '-'), nChar(releaseNameMaxLen+PADDING, '-'))

	for i := 0; i < len(cycleNames); i++ {
		table += fmt.Sprintf("|%s|%s|\n", centerString(cycleNames[i], cycleNameMaxLen), centerString(releaseNames[i], releaseNameMaxLen))
	}

	table += fmt.Sprintf("+%s+%s+", nChar(cycleNameMaxLen+PADDING, '-'), nChar(releaseNameMaxLen+PADDING, '-'))

	fmt.Fprintln(out, table)

	return nil
}

func PrintResourcesList(out io.Writer, resources []string) error {
	if _, err := fmt.Fprintf(out, "Found %d resources:\n", len(resources)); err != nil {
		return err
	}

	for _, resource := range resources {
		fmt.Fprintf(out, "- %s\n", resource)
	}

	return nil
}
