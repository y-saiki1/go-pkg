package model

import (
	"errors"
	"math"
)

const (
	OVER_TWENTY = 240
	// 離婚してから10ヶ月までは元旦那との子供だと許容する
	ALLOW_TEN_MONTHS_FROM_DIVORCE = -10
)

type Child struct {
	totalChildSupport int
	ageInMonths       int
}

func NewChild(months int) (*Child, error) {
	if OVER_TWENTY <= months {
		return nil, errors.New("20歳を超えているため、養育費を受け取れない子供です")
	}
	//離婚後10ヶ月まで元夫の子供だと許容する
	if months < ALLOW_TEN_MONTHS_FROM_DIVORCE {
		return nil, errors.New("指定日が子供の誕生日を下回っているため、月齢を計算できません")
	}

	// 月齢マイナス = 離婚後生まれたが元夫の子供という扱いの場合、年齢を０歳に設定する
	if months < 0 {
		months = 0
	}

	c := &Child{
		ageInMonths: months,
	}
	return c, nil
}

func (c *Child) TotalChildSupport() int {
	return c.totalChildSupport
}

func (c *Child) IsEndChildSupport(month int) bool {
	// fmt.Println("現在の年齢", (c.ageInMonths+month)/12, "養育費終わり", (c.ageInMonths+month) >= OVER_TWENTY)
	return OVER_TWENTY <= (c.ageInMonths + month)
}

func (c *Child) CalcMonthsUntilTwenty() (int, error) {
	if c.ageInMonths < 0 {
		return 0, errors.New("現在の月齢が設定されていません")
	}
	month := OVER_TWENTY - c.ageInMonths
	if month < 0 {
		return 0, nil
	}
	return OVER_TWENTY - c.ageInMonths, nil
}

func (c *Child) IsGthFifteen(months int) bool {
	return math.Floor(float64(c.ageInMonths+months)/12) >= 15
}

func (c *Child) PlsChildSupport(childSupport int) {
	c.totalChildSupport += childSupport
}

func (c *Child) AverageChildSupport() int {
	return c.totalChildSupport / (OVER_TWENTY - c.ageInMonths)
}
