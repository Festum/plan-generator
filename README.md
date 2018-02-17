# Assignment: Plan Generator (1A-2B-3X)

You are expected to complete the following task within 72 hours of receiving the
assignment.

In order to inform borrowers about the final repayment schedule, we need to have pre-
calculated repayment plans throughout the life time of a loan.

To be able to calculate a repayment plan specific input parameters are necessary:
- duration (number of instalments in months)
- nominal interest rate
- total loan amount ("total principal amount")
- Date of Disbursement/Payout

These four parameters need to be input parameters.

The goal is to calculate a repayment plan for an annuity loan. Therefor the amount that the
borrower has to payback every month, consisting principal and interest repayments, does
not change (the last instalment might be an exception).

The annuity amount has to be derived from three of the input parameters (duration,
nominal interest rate, total loan amount) before starting the plan calculation.
(use http://financeformulas.net/Annuity_Payment_Formula.html as reference)

## Usage

Direct pull this project and run will initial a web server for this API:
```
git clone plan-generator
cd plan-generator/
go run main.go
```

or include this lib to your code, example.go:
```
import "github.com/Festum/plan-generator/src/repayment"

// duration: number of instalments in months
// rate: nominal interest rate
// iop: total loan amount
// start: Date of Disbursement/Payout
pvPlans := PVPlan(24, 5, 5000, "01.01.2018")

// JSON version:
pvPlans2 := repayment.PVPlanJSON([]byte(`{"loanAmount":"5000", "nominalRate": "5", "duration": "24", "startDate": "2018-01-01T00:00:01Z"}`))
```

## Issues

Some values are incorrect in `Example Loan Details after annuity calculation`, such as:
- `4638 €` should be `4801.47 €`
- `01.01.2020` should be `01.12.2019`
- `"remainingOutstandingPrincipal": "0",` should be `"remainingOutstandingPrincipal": "0.00"`

Maybe I'm wrong please correct me :)
