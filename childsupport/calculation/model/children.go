package model

import (
	"errors"
)

const (
	OVER_FIFTEEN_INDEX   = 85
	UNDER_FOURTEEN_INDEX = 62
)

type Children []*Child

func (cr Children) howManyOverFifteen(month int) int {
	var count int
	for _, c := range cr {
		if c.IsGthFifteen(month) {
			count++
		}
	}
	return count
}

func (cr Children) canCalcChildSupport(rightHolderBasicIncome, obligorBasicIncome float64) bool {
	// 親の年収がないのであれば養育費は支給されない
	if (rightHolderBasicIncome+obligorBasicIncome) <= 0 || obligorBasicIncome <= 0 {
		return false
	}

	// Child構造体のコンストラクタに、２０さいを超えたら生成できない、というバリデーションが入っているので、Children構造体（スライス・配列）では要素数０＝子供ではない・子供がいないという状態
	if len(cr) == 0 {
		return false
	}

	return true
}

func (cr Children) calcChildSupportEachChild(childSupport, months int, totalFifteenIndex, totalFourteenIndex float64) {
	// 子供一人当たりの養育費金額概算 15の現在の養育費
	fifteenChildSupport := int(float64(childSupport*OVER_FIFTEEN_INDEX) / (totalFifteenIndex + totalFourteenIndex))
	// 子供一人当たりの養育費金額概算 14の現在の養育費
	fourteenChildSupport := int(float64(childSupport*UNDER_FOURTEEN_INDEX) / (totalFifteenIndex + totalFourteenIndex))

	for _, v := range cr {
		if v.IsGthFifteen(months) {
			v.PlsChildSupport(fifteenChildSupport)
			continue
		}
		v.PlsChildSupport(fourteenChildSupport)
	}
}

func (cr Children) CalcChildSupport(rightHolderBasicIncome, obligorBasicIncome float64) (int, int, int, error) {
	if !cr.canCalcChildSupport(rightHolderBasicIncome, obligorBasicIncome) {
		return 0, 0, 0, errors.New("養育費を計算できません")
	}

	lastIndex := len(cr) - 1
	lastChild := cr[lastIndex]
	totalPaymentMonth, err := lastChild.CalcMonthsUntilTwenty()
	if err != nil {
		return 0, 0, 0, err
	}

	var estimatedChildSupports []int
	var sumChildSupport int
	for i := 0; i < totalPaymentMonth; i++ {
		var children Children
		for _, v := range cr {
			if v.IsEndChildSupport(i) {
				continue
			}
			children = append(children, v)
		}

		// 養育費計算式は以下を参照してください
		// https://www.courts.go.jp/toukei_siryou/siryo/H30shihou_houkoku/index.html
		// https://www.mc-law.jp/rikon/22840/
		// https://conias.jp/%E3%80%90%E6%9C%80%E6%96%B0%E3%80%912019%E5%B9%B412%E6%9C%88%E6%94%B9%E8%A8%82%E3%81%AE%E9%A4%8A%E8%82%B2%E8%B2%BB%E3%81%AE%E6%A8%99%E6%BA%96%E7%9A%84%E7%AE%97%E5%AE%9A%E5%BC%8F%E3%81%AB%E3%82%88/
		overFifteenCount := children.howManyOverFifteen(i)
		underFourteenCount := len(children) - overFifteenCount
		totalFifteenIndex := float64(overFifteenCount * OVER_FIFTEEN_INDEX)
		totalFourteenIndex := float64(underFourteenCount * UNDER_FOURTEEN_INDEX)

		childrenLivingExpenses := obligorBasicIncome * ((totalFifteenIndex + totalFourteenIndex) / (100 + totalFifteenIndex + totalFourteenIndex))
		childSupport := int(childrenLivingExpenses * (obligorBasicIncome / (rightHolderBasicIncome + obligorBasicIncome) / 12))
		estimatedChildSupports = append(estimatedChildSupports, childSupport)
		sumChildSupport += childSupport

		children.calcChildSupportEachChild(childSupport, i, totalFifteenIndex, totalFourteenIndex)
	}

	// fmt.Println("平均単価", (sumChildSupport / totalPaymentMonth), "円", "合計給付金額", sumChildSupport, "円", "合計給付期間", totalPaymentMonth, "ヶ月")
	return (sumChildSupport / totalPaymentMonth), sumChildSupport, totalPaymentMonth, nil
}
