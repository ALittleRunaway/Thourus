package entity

type Project struct {
	Id      int     `json:"id"`
	Uid     string  `json:"uid"`
	Name    string  `json:"name"`
	Space   Space   `json:"space"`
	Company Company `json:"company"`
}
