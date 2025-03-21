package services

import (
	db "GoTwitter/db/repository"
	"GoTwitter/dto"
	"GoTwitter/models"
	"context"
)

type TweetService interface {
	CreateTweet(ctx context.Context, tweet *dto.CreateTweetDTO) (*models.Tweet, error)
	GetAllTweets(ctx context.Context) ([]*models.Tweet, error)
	GetTweetById(ctx context.Context, id int64) (*models.Tweet, error)
	DeleteTweet(ctx context.Context, id int64) (bool, error)
}

type tweetService struct {
	tweetRepository db.TweetsRepository
}

func NewTweetService(tweetRepository db.TweetsRepository) TweetService {
	return &tweetService{tweetRepository}
}

func (s *tweetService) CreateTweet(ctx context.Context, tweet *dto.CreateTweetDTO) (*models.Tweet, error) {
	return s.tweetRepository.Create(ctx, tweet)
}

func (s *tweetService) GetAllTweets(ctx context.Context) ([]*models.Tweet, error) {
	return s.tweetRepository.GetAll(ctx)
}

func (s *tweetService) GetTweetById(ctx context.Context, id int64) (*models.Tweet, error) {
	return s.tweetRepository.GetByID(ctx, id)
}

func (s *tweetService) DeleteTweet(ctx context.Context, id int64) (bool, error) {
	return s.tweetRepository.Delete(ctx, id)
}
