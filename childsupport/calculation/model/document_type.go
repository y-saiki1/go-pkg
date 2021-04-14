package model

const (
	NotarizedDocumentType     = 1
	JudgmentDocumentType      = 2
	WrittenRulingType         = 3
	RecordsOfConciliationType = 4
	RecordOfSettlementType    = 5
	OtherType                 = 9
	NonDocType                = 6
)

type DocumentType struct {
	value int
}

func NewDocumentType(value int) *DocumentType {
	return &DocumentType{value: value}
}

func (d *DocumentType) CalcBillingBeginDate(firstPaymentDate, billingEndDate, now *Date) int {
	// 現在の書面のタイプから、有効期限を取得
	switch d.value {
	case NotarizedDocumentType:
		return 5 * 12
	case JudgmentDocumentType:
		return 10 * 12
	case WrittenRulingType:
		return 10 * 12
	case RecordsOfConciliationType:
		return 10 * 12
	case RecordOfSettlementType:
		return 10 * 12
	case OtherType:
		return 5 * 12
	default:
		return 5 * 12
	}
	// 有効期限が現在からいつまでなのか図る
	// firstPaymentよりも前の日付ならfirstpayment 後ならbillingEndDate
}

// {
// 	"val": 1,
// 	"text": "公正証書",
// 	"description": "養育費の支払いについて相手方との間で協議し「公証役場」に行って作成した書面"
// },
// {
//     "val": 2,
//     "text": "判決正本",
//     "description": "養育費の支払いについて「訴訟」を起こし「裁判所」が養育費の支払いを相手方に命じた書面"
// },
// {
//     "val": 3,
//     "text": "審判書",
//     "description": "養育費の支払いについて「審判」を申立て「家庭裁判所」が養育費の支払いを相手方に命じた書面"
// },
// {
//     "val": 4,
//     "text": "調停調書",
//     "description": "養育費の支払いについて「調停」が成立した場合に「家庭裁判所」が作成した書面"
// },
// {
//     "val": 5,
//     "text": "和解調書",
//     "description": "養育費の支払いについて相手方と「和解」し「裁判所」が作成した書面"
// },
// {
//     "val": 99,
//     "text": "その他(上記以外の書面)",
//     "description": "裁判所や公証役場で作成されたものではなく、単に、当事者の間でした約束を書面にしたもの"
// },
// {
//     "val": 6,
//     "text": "なし",
//     "description": ""
// }
