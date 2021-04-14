package model

type Parent struct {
	workType WorkType
}

func NewParent(workType WorkType) *Parent {
	return &Parent{
		workType: workType,
	}
}

func (p *Parent) CalcBasicIncome() float64 {
	return p.workType.calcBasicIncome()
}
