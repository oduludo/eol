package datasource

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var mockClient = MockCycleClient[CycleDetail, ListedCycleDetail]{}

func TestCycleClient_Get(t *testing.T) {
	// Perform a Get() call on the mocked client
	cycleDetail, err, _ := mockClient.Get("smth_fake", "4.2")

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

	assert.Equal(t, cycleDetail.Eol, expectedResult.Eol)
	assert.Equal(t, cycleDetail.Latest, expectedResult.Latest)
	assert.Equal(t, cycleDetail.LatestReleaseDate, expectedResult.LatestReleaseDate)
	assert.Equal(t, cycleDetail.ReleaseDate, expectedResult.ReleaseDate)
	assert.Equal(t, cycleDetail.Lts, expectedResult.Lts)
}

func TestCycleClient_All(t *testing.T) {
	// Perform an All() call on the mocked client
	cycleList, err, _ := mockClient.All("smth_fake")

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, len(cycleList), 12)

	for _, cycleDetail := range cycleList {
		assert.NotNil(t, cycleDetail.Cycle)
	}
}
