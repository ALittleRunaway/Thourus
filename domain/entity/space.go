package entity

type Space struct {
	Id      int     `json:"id"`
	Uid     string  `json:"uid"`
	Name    string  `json:"name"`
	Company Company `json:"company"`
}
