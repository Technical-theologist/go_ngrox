package models

type FilterModel struct {
	StartingPrice int  `json:"lowerprice"`
	EndingPrice   int  `json:"upperprice"`
	InStock       bool `json:"instock"`
}
