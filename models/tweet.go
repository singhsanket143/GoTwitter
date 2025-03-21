package models

type Tweet struct {
	Id        int64  `json:"id"`
	Tweet     string `json:"tweet"`
	UserId    int64  `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
