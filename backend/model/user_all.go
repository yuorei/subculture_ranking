package model

type UserAll struct{
	user []UserAll
}
type User struct {
	UserId int `json:"userId"`
	UserName string `json:"userName"`
	ImageBase64 string `json:"imageBase64"`
}
