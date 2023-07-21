package datasource

import (
	"github.com/stretchr/testify/assert"
	"oduludo.io/eol/cfg"
	"testing"
)

var mockClient = MockCycleClient{}

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

// Assert the GetCustom method works, using a docker compose static file service
// Test using both an empty key and the key placeholder
func TestCycleClient_GetCustomForDecryptedData(t *testing.T) {
	for _, key := range []string{"", cfg.DecryptionKeyPlaceholder} {
		// Perform a successful GetCustom() call on the mocked client
		listedCycleDetail, err, _ := mockClient.GetCustom(
			"http://static/example_datasource_readonly.json",
			"ruby",
			"3.2",
			key,
		)

		if err != nil {
			t.Fatal(err)
		}

		expectedResult := ListedCycleDetail{
			CycleDetail: CycleDetail{
				Eol: "2026-03-31",
			},
			Cycle: "3.2",
		}

		assert.Equal(t, listedCycleDetail.Eol, expectedResult.Eol)
		assert.Equal(t, listedCycleDetail.Cycle, expectedResult.Cycle)
	}
}

func TestCycleClient_GetCustomForEncryptedData(t *testing.T) {
	const key = "rijltlTWRbYHtVaS"

	// Perform a successful GetCustom() call on the mocked client
	// The call is made on encrypted data, so the getCustom() implementation should decrypt the data using the provided key.
	listedCycleDetail, err, _ := mockClient.GetCustom(
		"http://static/example_datasource_readonly_encrypted.json",
		"ruby",
		"3.2",
		key,
	)

	if err != nil {
		t.Fatal(err)
	}

	expectedResult := ListedCycleDetail{
		CycleDetail: CycleDetail{
			Eol: "2026-03-31",
		},
		Cycle: "3.2",
	}

	assert.Equal(t, listedCycleDetail.Eol, expectedResult.Eol)
	assert.Equal(t, listedCycleDetail.Cycle, expectedResult.Cycle)
}
