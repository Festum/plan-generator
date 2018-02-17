package repayment

import (
	"fmt"
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func TestNumbers(t *testing.T) {
	assertEqual(t, Round(33.6677, .5, 2), 33.67, "Round result is wrong")
	assertEqual(t, Round(33.6677, 1, 2), 33.66, "Round result is not Floor")
	assertEqual(t, f2str(33.67), "33.67", "Failed to convert float64 to string")
	assertEqual(t, getAnnuity(24, 5, 5000), 219.36, "Calculation for annuity is not correct")
}

func TestGenerateDates(t *testing.T) {
	dates, dates2 := []string{}, []string{}
	for date := range generateDates("13.12.2017", 2) {
		dates = append(dates, date)
	}
	assertEqual(t, dates[0], "2017-12-13T00:00:00Z", "")
	assertEqual(t, dates[1], "2018-01-13T00:00:00Z", "")
	for date := range generateDates("2018-01-01T00:00:00Z", 2) {
		dates2 = append(dates2, date)
	}
	assertEqual(t, dates2[1], "2018-02-01T00:00:00Z", "")
}

func TestPlan(t *testing.T) {
	plans := PVPlan(24, 5, 5000, "01.01.2018")
	assertEqual(t, plans[0], Plan{"219.36", "2018-01-01T00:00:00Z", "5000.00", "20.83", "198.53", "4801.47"}, "")
	assertEqual(t, plans[1], Plan{"219.36", "2018-02-01T00:00:00Z", "4801.47", "20.00", "199.36", "4602.11"}, "")
	assertEqual(t, plans[len(plans)-1], Plan{"219.15", "2019-12-01T00:00:00Z", "218.25", "0.90", "218.25", "0.00"}, "")
}
