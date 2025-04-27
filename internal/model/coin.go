package model

type Market struct {
	Market     string `json:"market"`
	KoreanName string `json:"korean_name"`
}

type Ticker struct {
	Market           string  `json:"market"`
	SignedChangeRate float64 `json:"signed_change_rate"`
}

type ChangeInfo struct {
	Symbol     string
	KoreanName string
	ChangeRate float64
}
