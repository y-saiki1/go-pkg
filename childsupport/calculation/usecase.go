package calculation

import (
	"time"

	"Aquarius-neo/pkg/childsupport/calculation/model"
	"Aquarius-neo/pkg/childsupport/calculation/output"
)

func Calculate(
	birthDates []*time.Time,
	dateForCalculationAgeOfChild *time.Time,
	rightHolderAnnualIncome int,
	obligorAnnualIncome int,
	rightHolderWorkType interface{},
	obligorWorkType interface{},
) (*output.CalculationOutput, error) {
	workType, err := model.NewWorkType(rightHolderWorkType, rightHolderAnnualIncome)
	if err != nil {
		return nil, err
	}
	rightHolder := model.NewParent(workType)

	workType, err = model.NewWorkType(obligorWorkType, obligorAnnualIncome)
	if err != nil {
		return nil, err
	}
	obligor := model.NewParent(workType)

	dateForCalc := model.NewDate(dateForCalculationAgeOfChild)
	var children model.Children
	for _, v := range birthDates {
		months := dateForCalc.SubMonths(model.NewDate(v))
		c, err := model.NewChild(months)
		if err != nil {
			continue
		}
		children = append(children, c)
	}

	avarageChildSupport, totalChildSupport, totalPaymentMonth, err := children.CalcChildSupport(rightHolder.CalcBasicIncome(), obligor.CalcBasicIncome())
	if err != nil {
		return nil, err
	}

	var list []*output.CalculationEachChildOutput
	for _, v := range children {
		totalPaymentMonth, err := v.CalcMonthsUntilTwenty()
		if err != nil {
			continue
		}
		list = append(list, &output.CalculationEachChildOutput{
			Sum:                v.TotalChildSupport(),
			Average:            v.AverageChildSupport(),
			TotalPaymentMonths: totalPaymentMonth,
		})
	}

	return &output.CalculationOutput{
		Sum:                        totalChildSupport,
		Average:                    avarageChildSupport,
		TotalPaymentMonths:         totalPaymentMonth,
		CalculationEachChildOutput: list,
	}, nil
}
