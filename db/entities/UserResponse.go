package entities

type UserRespnse struct {
	Id      int    `json:"id"`
	FName   string `json:"fname"`
	City    string `json:"city"`
	Phone   int64  `json:"phone"`
	Height  int    `json:"height"`
	Married bool   `json:"Married"`
}
