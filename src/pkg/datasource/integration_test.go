package datasource

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Integration tests, aimed at:
// 1) confirming the integration with the datasource is still good and;
// 2) ensuring test coverage for the non-mock (real) clients
//
// Be mindful we don't control the data https://endoflife.date/ gives, so this test
// may break without any changes to the EOL codebase. If this integration test breaks
// it is likely the whole application can no longer fetch its data for usage.

var cycleDetailClient = CycleClient[CycleDetail, ListedCycleDetail]{}

func TestIntegrationCycleDetailClient_Get(t *testing.T) {
	// Perform a Get on the real client
	cycleDetail, err, notFound := cycleDetailClient.Get("ruby", "2.7")

	assert.False(t, notFound, "resource 'ruby@2.7' was not found in datasource")

	if err != nil {
		t.Fatal(err)
	}

	// Use real data that is unlikely to change
	expectedResult := CycleDetail{
		Eol:               "2023-03-31",
		Latest:            "2.7.8",
		LatestReleaseDate: "2023-03-30",
		ReleaseDate:       "2019-12-25",
		Lts:               false,
	}

	if cycleDetail.Eol != expectedResult.Eol ||
		cycleDetail.Latest != expectedResult.Latest ||
		cycleDetail.LatestReleaseDate != expectedResult.LatestReleaseDate ||
		cycleDetail.ReleaseDate != expectedResult.ReleaseDate ||
		cycleDetail.Lts != expectedResult.Lts {
		t.Fail()
	}
}
