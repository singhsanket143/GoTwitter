package services

import (
	db "GoTwitter/db/repository"
	"GoTwitter/dto"
	"GoTwitter/models"
	"GoTwitter/utils"
	"context"
	"fmt"
	"sync"
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

// CreateTweet handles the creation of a new tweet along with its associated hashtags and tweet-tags.
// It performs the following steps:
// 1. Parses hashtags from the provided tweet content.
// 2. Creates the hashtags (tags) in the database using a goroutine.
// 3. Creates the tweet in the database using another goroutine.
// 4. Waits for both the hashtag creation and tweet creation to complete.
// 5. Associates the created hashtags with the tweet by creating tweet-tags in the database.
//
// Parameters:
// - ctx: The context for managing request-scoped values, deadlines, and cancellations.
// - tweet: A pointer to a CreateTweetDTO object containing the tweet content.
//
// Returns:
// - A pointer to the created Tweet model if successful.
// - An error if any step in the process fails.
func (s *tweetService) CreateTweet(ctx context.Context, tweet *dto.CreateTweetDTO) (*models.Tweet, error) {
	tweetContent := tweet.Tweet

	// parse hashtags from tweet content
	hashtags := utils.ParseHashtags(tweetContent)
	fmt.Println("Hashtags: ", hashtags)

	var (
		newtweet *models.Tweet
		tags     []*models.Tag
		tweetErr error
		tagsErr  error
	)

	// Use a WaitGroup to wait for both goroutines to complete
	var wg sync.WaitGroup
	wg.Add(2)

	// Create tags in a goroutine
	go func() {
		defer wg.Done()
		tags, tagsErr = s.tagRepository.BulkCreate(ctx, hashtags)
	}()

	// Create tweet in a goroutine
	go func() {
		defer wg.Done()
		newtweet, tweetErr = s.tweetRepository.Create(ctx, tweet)
	}()

	// Wait for both goroutines to complete
	wg.Wait()

	// Check for errors
	if tagsErr != nil {
		return nil, tagsErr
	}
	if tweetErr != nil {
		return nil, tweetErr
	}

	fmt.Printf("New Tweet: %+v\n", newtweet)
	for _, tag := range tags {
		fmt.Printf("Tag: %+v\n", tag)
	}

	// Create tweet_tags
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
