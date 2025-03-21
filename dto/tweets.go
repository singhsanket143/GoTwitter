package dto

type CreateTweetDTO struct {
	Tweet  string `json:"tweet" validate:"required,min=1,max=280"`
	UserId int64  `json:"user_id" validate:"required"`
}
