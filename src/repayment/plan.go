package repayment

import (
	"encoding/json"
	"log"
	"math"
	"strconv"
	"time"
)

const DateFormat = "02.01.2006"
const DateLongFormat = "2006-01-02T15:04:05Z"

func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

func getAnnuity(duration float64, rate float64, principal float64) float64 {
	n := 12.0 // number of months per year
	m := (principal * rate / 100 / n) / (1 - math.Pow(1+rate/100/n, -duration))
	return Round(m, .5, 2)
}

func generateDates(t string, n int) chan string {
	ch := make(chan string)

	date, err := time.Parse(DateFormat, t)
	if err != nil {
		err = nil
		date, err = time.Parse(DateLongFormat, t)
		if err != nil {
			log.Fatal(err)
			return ch
		}
	}
	y, m, d := date.Date()
	nextMonth, nextYear := int(m), int(y)

	go func() {
		defer close(ch)
		for i := 0; i < n; i++ {
			date, err = time.Parse("2.1.2006", strconv.Itoa(d)+"."+
				strconv.Itoa(nextMonth)+"."+
				strconv.Itoa(nextYear))
			if err != nil {
				log.Fatal(err)
				continue
			}
			ch <- date.Format(DateLongFormat)

			nextMonth++
			if nextMonth > 12 {
				nextMonth = 1
				nextYear++
			}
		}
	}()
	return ch
}

type Loan struct {
	LoanAmount  float64 `json:"loanAmount,string,float64"`
	NominalRate float64 `json:"nominalRate,string,float64"`
	Duration    float64 `json:"duration,string,float64"`
	StartDate   string  `json:"startDate"`
}

type Plan struct { // Use string to prevent additional digits
	BorrowerPaymentAmount         string `json:"borrowerPaymentAmount"`
	Date                          string `json:"date"`
	InitialOutstandingPrincipal   string `json:"initialOutstandingPrincipal"`
	Interest                      string `json:"interest"`
	Principal                     string `json:"principal"`
	RemainingOutstandingPrincipal string `json:"remainingOutstandingPrincipal"`
}

func f2str(fv float64) string {
	return strconv.FormatFloat(fv, 'f', 2, 64)
}

func PVPlan(duration float64, rate float64, iop float64, start string) []Plan {
	pv := getAnnuity(duration, rate, iop)
	// Interest = (Nominal-Rate * Days in Month * Initial Outstanding Principal) / days in year
	dayInMonth, dayInYear := 30.0, 360.0
	rop, principle := 0.0, 0.0
	out := []Plan{}

	for date := range generateDates(start, int(duration)) {
		if date != start {
			iop -= principle
		}
		i := Round(((rate*dayInMonth*iop)/dayInYear)/100, 1, 2) //Round up by 1 based on example 20.00
		principle = pv - i
		rop = iop - principle
		if rop < 0 {
			pv += rop
			principle = pv - i
			rop = 0
		}
		out = append(out, Plan{f2str(pv), date, f2str(iop), f2str(i), f2str(principle), f2str(rop)})
	}

	return out
}

func PVPlanJSON(in []byte) []Plan {
	var loan Loan
	err := json.Unmarshal(in, &loan)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return PVPlan(
		loan.Duration,
		loan.NominalRate,
		loan.LoanAmount,
		loan.StartDate)
}
