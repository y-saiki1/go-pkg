package output

type CalculationEachChildOutput struct {
	Sum                int `json:"sum"`
	Average            int `json:"average"`
	TotalPaymentMonths int `json:"total_payment_months"`
}

type CalculationOutput struct {
	Sum                        int                           `json:"sum"`
	Average                    int                           `json:"average"`
	TotalPaymentMonths         int                           `json:"total_payment_months"`
	CalculationEachChildOutput []*CalculationEachChildOutput `json:"calculation_each_child"`
}
