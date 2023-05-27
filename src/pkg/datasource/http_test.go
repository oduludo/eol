package datasource

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var mockClient = MockCycleClient[CycleDetail, ListedCycleDetail]{}

func TestCycleClient_Get(t *testing.T) {
	// Perform a successful Get() call on the mocked client
	cycleDetail, err, _ := mockClient.Get("ruby", "3.2")

	if err != nil {
		t.Fatal(err)
	}

	expectedResult := CycleDetail{
		Eol:               "2046-03-31",
		Latest:            "3.2.2",
		LatestReleaseDate: "2023-03-30",
		ReleaseDate:       "2022-12-25",
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
	cycleList, err, _ := mockClient.All("ruby")

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, len(cycleList), 12)

	for _, cycleDetail := range cycleList {
		assert.NotNil(t, cycleDetail.Cycle)
	}
}
