package input

import "time"

type CalculationInput struct {
	birthDates                   []time.Time
	dateForCalculationAgeOfChild time.Time
	rightHolderAnnualIncome      int
	obligorAnnualIncome          int
	rightHolderWorkType          interface{}
	obligorWorkType              interface{}
}
