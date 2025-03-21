package services

import (
	db "GoTwitter/db/repository"
	"GoTwitter/models"
	"context"
)

type TweetService interface {
	CreateTweet(ctx context.Context, tweet *models.Tweet) error
}

type tweetService struct {
	tweetRepository db.TweetsRepository
}

func NewTweetService(tweetRepository db.TweetsRepository) TweetService {
	return &tweetService{tweetRepository}
}

func (s *tweetService) CreateTweet(ctx context.Context, tweet *models.Tweet) error {
	return s.tweetRepository.Create(ctx, tweet)
}
