package datasource

import (
	"testing"
)

var client = MockCycleDetailClient[CycleDetail]{}

func TestCycleDetailClient_Get(t *testing.T) {
	// Perform a Get on the mocked client
	cycleDetail, err := client.Get("smth_fake", "4.2")

	if err != nil {
		t.Fatal(err)
	}

	expectedResult := CycleDetail{
		Eol:               "2048-03-15",
		Latest:            "2023.0.20230503.0",
		LatestReleaseDate: "2023-05-04",
		ReleaseDate:       "2023-03-01",
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
