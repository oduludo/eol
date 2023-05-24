package datasource

import (
	"fmt"
	"testing"
	"time"
)

func TestCycleDetailFromBytes(t *testing.T) {
	data := loadMockData("cycle_detail.json")
	_, err := newCycleDetailFromBytes(data)

	if err != nil {
		t.Fatal(err)
	}
}

func TestNewObjectFromBytes(t *testing.T) {
	data := loadMockData("cycle_detail.json")
	obj := CycleDetail{}

	if err := newObjectFromBytes(&obj, data); err != nil {
		t.Fatal(err)
	}
}

func TestCycleDetail_EolTime(t *testing.T) {
	data := loadMockData("cycle_detail.json")
	cycleData, err := newCycleDetailFromBytes(data)

	if err != nil {
		t.Fatal(err)
	}

	// Should be `2048-03-15 00:00:00`
	eolTime := cycleData.EolTime()

	if eolTime.Year() != 2048 ||
		eolTime.Month() != 3 ||
		eolTime.Day() != 15 ||
		eolTime.Hour() != 0 ||
		eolTime.Minute() != 0 ||
		eolTime.Second() != 0 {
		t.Fail()
	}
}

// Test hasPassedEolWithTime returns the correct value for a resource version which has reached EOL.
// The CycleDetail's EOL value stays the same. The `now` Time value is shifted to create a mimicked time difference.
func testCompareEolTimeWith(t *testing.T, cd *CycleDetail, yearOffset int, expected bool) {
	// We are testing against an EOL of `2048-03-15`
	result := hasPassedEolWithTime(
		cd,
		time.Date(2048+yearOffset, 3, 15, 0, 0, 0, 0, time.UTC),
	)

	if result != expected {
		t.Fail()
	}
}

type EolCompare struct {
	Name       string
	YearOffset int
	Expected   bool
}

func TestCompareEolTimeWith(t *testing.T) {
	data := loadMockData("cycle_detail.json")
	cycleData, err := newCycleDetailFromBytes(data)

	if err != nil {
		t.Fatal(err)
	}

	table := []EolCompare{
		{
			// Mimics EOL one year in the past
			Name:       "eol_in_past",
			YearOffset: 1,
			Expected:   true,
		},
		{
			// Mimics EOL being today, which should not flag the EOL for this resource+version
			Name:       "eol_now",
			YearOffset: 0,
			Expected:   false,
		},
		{
			// Mimics EOL one year ahead from now
			Name:       "eol_in_future",
			YearOffset: -1,
			Expected:   false,
		},
	}

	for _, testValue := range table {
		t.Run(fmt.Sprintf("%s_%s", t.Name(), testValue.Name), func(t *testing.T) {
			testCompareEolTimeWith(t, &cycleData, testValue.YearOffset, testValue.Expected)
		})
	}
}
