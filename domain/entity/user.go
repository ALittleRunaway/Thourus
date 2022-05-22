package entity

type User struct {
	Id       int       `json:"id"`
	Uid      string    `json:"uid"`
	Name     string    `json:"name"`
	Surname  string    `json:"surname"`
	Email    string    `json:"email"`
	Role     Role      `json:"role"`
	Company  Company   `json:"company"`
	Projects []Project `json:"projects"`
}
