package model

import "time"

type Profile struct {
	UserId      int `json:"userId"`
	UserName    int	`json:"uuid"`
	Age         int`json:"age"`
	ImageBase64 string`json:"imageBase64"`
	CreatedAt   time.Time `json:"createdAt"`
}
