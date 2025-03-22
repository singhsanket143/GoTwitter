package handlers

import (
	"GoTwitter/dto"
	"GoTwitter/services"
	"GoTwitter/utils"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type TweetHandler struct {
	tweetService services.TweetService
}

func NewTweetHandler(tweetService services.TweetService) *TweetHandler {
	return &TweetHandler{tweetService}
}

func (h *TweetHandler) CreateTweetHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var payload dto.CreateTweetDTO
	if err := utils.ReadJson(r, &payload); err != nil {
		utils.WriteJsonError(w, http.StatusBadRequest, "Something went wrong", err.Error())
		return
	}

	tweet, err := h.tweetService.CreateTweet(ctx, &payload)
	if err != nil {
		utils.WriteJsonError(w, err.Code, "Something went wrong", err.Error())
		return
	}

	utils.WriteJsonSuccess(w, http.StatusCreated, "Tweet created successfully", tweet)
}

func (h *TweetHandler) GetAllTweetsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tweets, err := h.tweetService.GetAllTweets(ctx)
	if err != nil {
		utils.WriteJsonError(w, http.StatusInternalServerError, "Something went wrong", err.Error())
		return
	}

	utils.WriteJsonSuccess(w, http.StatusOK, "Tweets fetched successfully", tweets)
}

func (h *TweetHandler) GetTweetByIdHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tweetId := chi.URLParam(r, "tweetId")

	id, err := strconv.ParseInt(tweetId, 10, 64)
	if err != nil {
		utils.WriteJsonError(w, http.StatusBadRequest, "Invalid tweet ID", err.Error())
		return
	}

	tweet, err := h.tweetService.GetTweetById(ctx, id)

	if err != nil {
		utils.WriteJsonError(w, http.StatusInternalServerError, "Something went wrong", err.Error())
		return
	}

	utils.WriteJsonSuccess(w, http.StatusOK, "Tweet fetched successfully", tweet)
}

func (h *TweetHandler) DeleteTweetHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tweetId := chi.URLParam(r, "tweetId")

	id, err := strconv.ParseInt(tweetId, 10, 64)
	if err != nil {
		utils.WriteJsonError(w, http.StatusBadRequest, "Invalid tweet ID", err.Error())
		return
	}

	deleted, err := h.tweetService.DeleteTweet(ctx, id)
	if err != nil {
		utils.WriteJsonError(w, http.StatusInternalServerError, "Something went wrong", err.Error())
		return
	}

	if !deleted {
		utils.WriteJsonError(w, http.StatusNotFound, "Tweet not found", nil)
		return
	}

	utils.WriteJsonSuccess(w, http.StatusOK, "Tweet deleted successfully", nil)

}
