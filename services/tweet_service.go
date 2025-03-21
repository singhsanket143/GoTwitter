package services

import (
	db "GoTwitter/db/repository"
	"GoTwitter/dto"
	"GoTwitter/models"
	"GoTwitter/utils"
	"context"
	"fmt"
)

type TweetService interface {
	CreateTweet(ctx context.Context, tweet *dto.CreateTweetDTO) (*models.Tweet, error)
	GetAllTweets(ctx context.Context) ([]*models.Tweet, error)
	GetTweetById(ctx context.Context, id int64) (*models.Tweet, error)
	DeleteTweet(ctx context.Context, id int64) (bool, error)
}

type tweetService struct {
	tweetRepository     db.TweetsRepository
	tagRepository       db.TagsRepository
	tweetTagsRepository db.TweetTagsRepository
}

func NewTweetService(
	tweetRepository db.TweetsRepository,
	tagRepository db.TagsRepository,
	tweetTagsRepository db.TweetTagsRepository) TweetService {
	return &tweetService{tweetRepository, tagRepository, tweetTagsRepository}
}

func (s *tweetService) CreateTweet(ctx context.Context, tweet *dto.CreateTweetDTO) (*models.Tweet, error) {

	tweetContent := tweet.Tweet

	// parse hashtags from tweet content
	hashtags := utils.ParseHashtags(tweetContent)

	fmt.Println("Hashtags: ", hashtags)

	// create tags
	tags, err := s.tagRepository.BulkCreate(ctx, hashtags)

	if err != nil {
		return nil, err
	}

	for _, tag := range tags {
		fmt.Printf("Tag: %+v\n", tag)
	}

	newtweet, tweetErr := s.tweetRepository.Create(ctx, tweet)

	if tweetErr != nil {
		return nil, tweetErr
	}

	fmt.Printf("New Tweet: %+v\n", newtweet)

	// create tweet_tags
	tagIds := make([]int64, 0)

	for _, tag := range tags {
		tagIds = append(tagIds, tag.Id)
	}

	_, tweetTagsErr := s.tweetTagsRepository.BulkCreate(ctx, tagIds, newtweet.Id)

	if tweetTagsErr != nil {
		return nil, tweetTagsErr
	}

	return newtweet, nil
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
