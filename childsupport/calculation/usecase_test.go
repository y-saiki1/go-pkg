package calculation

import (
	"testing"
	"time"
)

func Test養育費計算(t *testing.T) {
	type args struct {
		birthDates              []*time.Time
		divorcedAt              *time.Time
		rightHolderAnnualIncome int
		obligorAnnualIncome     int
		rightHolderWorkType     interface{}
		obligorWorkType         interface{}
	}
	time1, _ := time.Parse("2006-01-02", "2001-11-11")
	time2, _ := time.Parse("2006-01-02", "2001-11-11")
	time3, _ := time.Parse("2006-01-02", "2018-11-11")

	tests := []struct {
		name string
		args args
	}{
		{
			name: "計算成功",
			args: args{
				birthDates:              []*time.Time{&time1, &time2},
				divorcedAt:              &time3,
				rightHolderAnnualIncome: 300,
				obligorAnnualIncome:     400,
				rightHolderWorkType:     1,
				obligorWorkType:         2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Calculate(tt.args.birthDates, tt.args.divorcedAt, tt.args.rightHolderAnnualIncome, tt.args.obligorAnnualIncome, tt.args.rightHolderWorkType, tt.args.obligorWorkType)
			if err != nil {
				t.Errorf("エラー内容: %v", err)
			}
		})
	}

	time1, _ = time.Parse("2006-01-02", "2001-11-11")
	time2, _ = time.Parse("2006-01-02", "2001-11-11")
	time3, _ = time.Parse("2006-01-02", "2018-11-11")
	time4, _ := time.Parse("1994-01-02", "2019-11-11")
	time5, _ := time.Parse("1995-01-02", "2019-11-11")

	tests = []struct {
		name string
		args args
	}{
		{
			name: "ありえないワークタイプ値を入れた場合",
			args: args{
				birthDates:              []*time.Time{&time1, &time2},
				divorcedAt:              &time3,
				rightHolderAnnualIncome: 300,
				obligorAnnualIncome:     400,
				rightHolderWorkType:     32.4,
				obligorWorkType:         true,
			},
		},
		{
			name: "義務者の年収を0円で入れた場合",
			args: args{
				birthDates:              []*time.Time{&time1, &time2},
				divorcedAt:              &time3,
				rightHolderAnnualIncome: 300,
				obligorAnnualIncome:     0,
				rightHolderWorkType:     "string",
				obligorWorkType:         "string",
			},
		},
		{
			name: "子供が全員２０歳を超えていた場合",
			args: args{
				birthDates:              []*time.Time{&time4, &time5},
				divorcedAt:              &time3,
				rightHolderAnnualIncome: 300,
				obligorAnnualIncome:     400,
				rightHolderWorkType:     "string",
				obligorWorkType:         "string",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Calculate(tt.args.birthDates, tt.args.divorcedAt, tt.args.rightHolderAnnualIncome, tt.args.obligorAnnualIncome, tt.args.rightHolderWorkType, tt.args.obligorWorkType)
			if err == nil {
				t.Errorf("エラーがnilで返却")
				t.Errorf("内容: %#v", tt.args)
			}
		})
	}
}
