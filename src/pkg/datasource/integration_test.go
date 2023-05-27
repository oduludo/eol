package datasource

import (
	"github.com/stretchr/testify/assert"
	"oduludo.io/eol/cfg"
	"os"
	"testing"
)

// Integration tests, aimed at:
// 1) confirming the integration with the datasource is still good and;
// 2) ensuring test coverage for the non-mock (real) clients
//
// Be mindful we don't control the data https://endoflife.date/ gives, so this test
// may break without any changes to the EOL codebase. If this integration test breaks
// it is likely the whole application can no longer fetch its data for usage.

func TestIntegrationCycleClient_Get(t *testing.T) {
	if err := os.Setenv(cfg.IsIntegrationTestEnvKey, "true"); err != nil {
		t.Fatal(err)
	}

	// Perform a Get() on the real client
	cycleDetail, err, notFound := NewCycleClient().Get("ruby", "2.7")

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

	if err := os.Unsetenv(cfg.IsIntegrationTestEnvKey); err != nil {
		t.Fatal(err)
	}
}

func TestIntegrationCycleClient_All(t *testing.T) {
	if err := os.Setenv(cfg.IsIntegrationTestEnvKey, "true"); err != nil {
		t.Fatal(err)
	}

	// Perform an All() on the real client
	cycleList, err, notFound := NewCycleClient().All("ruby")

	assert.False(t, notFound, "resource 'ruby' was not found in datasource")

	if err != nil {
		t.Fatal(err)
	}

	// At the time of writing this test, there are 12 entries.
	// This will likely increase over time, so don't fix the number, but do a GTE check.
	assert.GreaterOrEqual(t, len(cycleList), 12)

	if err := os.Unsetenv(cfg.IsIntegrationTestEnvKey); err != nil {
		t.Fatal(err)
	}
}

func TestIntegrationCycleClient_Resources(t *testing.T) {
	if err := os.Setenv(cfg.IsIntegrationTestEnvKey, "true"); err != nil {
		t.Fatal(err)
	}

	// Perform a Resources() on the real client
	resources, err := NewCycleClient().Resources()

	if err != nil {
		t.Fatal(err)
	}

	// Check the retrieved list is not empty
	assert.Greater(t, len(resources), 0)

	if err := os.Unsetenv(cfg.IsIntegrationTestEnvKey); err != nil {
		t.Fatal(err)
	}
}
