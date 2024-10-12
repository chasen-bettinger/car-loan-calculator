package main

import (
	"fmt"

	gofinancial "github.com/razorpay/go-financial"
	"github.com/razorpay/go-financial/enums/paymentperiod"
	"github.com/shopspring/decimal"
)

type LoanDetails struct {
	Label string
	Rate  float64
	Years int
	Pv    decimal.Decimal
}

func getLoanDetails(loanDetails LoanDetails) {

	fmt.Printf("\n\nPurchasing the %v:", loanDetails.Label)
	fmt.Printf("\n\nloanDetails: %+v\n", loanDetails)

	rate := decimal.NewFromFloat(loanDetails.Rate / 12)
	years := loanDetails.Years
	months := 12 * years
	nper := int64(months)
	pv := loanDetails.Pv
	fv := decimal.NewFromInt(0)
	when := paymentperiod.ENDING
	pmt := gofinancial.Pmt(rate, nper, pv, fv, when)
	fmt.Printf("payment: $%.2f", pmt.InexactFloat64())

	monthsCounter := months

	interestSum := decimal.NewFromInt(0)
	pmtSum := decimal.NewFromInt(0)
	pv = pv.Abs()
	for monthsCounter > 0 {
		interest := pv.Mul(rate).Abs()
		principal := pmt.Sub(interest)

		pv = pv.Sub(principal)
		pmtSum = pmtSum.Add(pmt)
		interestSum = interestSum.Add(interest)

		monthsCounter--
	}

	fmt.Println()
	fmt.Printf("interestSum: $%.2f\n", interestSum.InexactFloat64())
	fmt.Printf("pmtSum: $%.2f", pmtSum.InexactFloat64())
}

func main() {

	ld := []LoanDetails{
		{
			Label: "volvo",
			Rate:  0.07,
			Years: 5,
			Pv:    decimal.NewFromInt(-40_000),
		},
		{
			Label: "tesla FSD",
			Rate:  0,
			Years: 6,
			Pv:    decimal.NewFromInt(-51_000),
		},
		{
			Label: "tesla no FSD",
			Rate:  0.0529,
			Years: 6,
			Pv:    decimal.NewFromInt(-43_690),
		},

		{
			Label: "volvo w/ down payment",
			Rate:  0.07,
			Years: 5,
			Pv:    decimal.NewFromInt(-30_000),
		},
		{
			Label: "tesla FSD w/ down payment ",
			Rate:  0,
			Years: 6,
			Pv:    decimal.NewFromInt(-45_042),
		},
		{
			Label: "tesla no FSD w/ down payment ",
			Rate:  0.0529,
			Years: 6,
			Pv:    decimal.NewFromInt(-38_242),
		},
	}

	for _, loan := range ld {
		getLoanDetails(loan)
	}
}
