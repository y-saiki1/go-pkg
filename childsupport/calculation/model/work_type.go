package model

import (
	"errors"
)

const (
	CONVERT_MILLION = 10000
)

const (
	EMPLOYEE = iota + 1
	PART_TIME_JOB
	TEMPORARY_EMPLOYEE
	SELF_EMPLOYEE
	CEO
	UNEMPLOYED
)

type WorkType interface {
	String() string
	calcBasicIncome() float64
}

func NewWorkType(workType interface{}, annualIncome int) (WorkType, error) {
	if str, ok := workType.(string); ok {
		return NewWorkTypeFromString(str, annualIncome), nil
	}
	if integer, ok := workType.(uint); ok {
		return NewWorkTypeFromInt(integer, annualIncome), nil
	}
	return nil, errors.New("workTypeに変換できませんでした")
}

func NewWorkTypeFromInt(workTypeAsInt uint, annualIncome int) WorkType {
	switch workTypeAsInt {
	case EMPLOYEE:
		return &EmployeeType{value: annualIncome}
	case PART_TIME_JOB:
		return &EmployeeType{value: annualIncome}
	case TEMPORARY_EMPLOYEE:
		return &EmployeeType{value: annualIncome}
	case SELF_EMPLOYEE:
		return &SelfEmployeeType{value: annualIncome}
	case CEO:
		return &EmployeeType{value: annualIncome}
	case UNEMPLOYED:
		return &EmployeeType{value: annualIncome}
	default:
		return &EmployeeType{value: annualIncome}
	}
}

func NewWorkTypeFromString(workTypeAsString string, annualIncome int) WorkType {
	switch workTypeAsString {
	case "自営業":
		return &SelfEmployeeType{value: annualIncome}
	case "給与所得者":
		return &EmployeeType{value: annualIncome}
	default:
		return &EmployeeType{value: annualIncome}
	}
}

type EmployeeType struct {
	value int
}

func (w *EmployeeType) String() string {
	return "給与所得者"
}

func (e *EmployeeType) calcBasicIncome() float64 {
	switch {
	case e.value <= 100:
		return float64((e.value * CONVERT_MILLION)) * 0.42
	case e.value <= 125:
		return float64((e.value * CONVERT_MILLION)) * 0.41
	case e.value <= 150:
		return float64((e.value * CONVERT_MILLION)) * 0.40
	case e.value <= 250:
		return float64((e.value * CONVERT_MILLION)) * 0.39
	case e.value <= 500:
		return float64((e.value * CONVERT_MILLION)) * 0.38
	case e.value <= 700:
		return float64((e.value * CONVERT_MILLION)) * 0.37
	case e.value <= 850:
		return float64((e.value * CONVERT_MILLION)) * 0.36
	case e.value <= 1350:
		return float64((e.value * CONVERT_MILLION)) * 0.35
	case e.value <= 2000:
		return float64((e.value * CONVERT_MILLION)) * 0.34
	default:
		return float64((e.value * CONVERT_MILLION)) * 0.34
	}
}

type SelfEmployeeType struct {
	value int
}

func (s *SelfEmployeeType) String() string {
	return "自営業"
}

func (s *SelfEmployeeType) calcBasicIncome() float64 {
	switch {
	case s.value <= 421:
		return float64((s.value * CONVERT_MILLION)) * 0.52
	case s.value <= 526:
		return float64((s.value * CONVERT_MILLION)) * 0.51
	case s.value <= 870:
		return float64((s.value * CONVERT_MILLION)) * 0.50
	case s.value <= 975:
		return float64((s.value * CONVERT_MILLION)) * 0.49
	case s.value <= 1144:
		return float64((s.value * CONVERT_MILLION)) * 0.48
	case s.value <= 1409:
		return float64((s.value * CONVERT_MILLION)) * 0.47
	default:
		return float64((s.value * CONVERT_MILLION)) * 0.47
	}
}
