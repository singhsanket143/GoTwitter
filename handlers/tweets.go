package handlers

import (
	"GoTwitter/models"
	"GoTwitter/services"
	"GoTwitter/utils"
	"net/http"
)

type TweetHandler struct {
	tweetService services.TweetService
}

func NewTweetHandler(tweetService services.TweetService) *TweetHandler {
	return &TweetHandler{tweetService}
}

func (h *TweetHandler) CreateTweetHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var tweet models.Tweet
	if err := utils.ReadJson(r, &tweet); err != nil {
		utils.WriteJsonError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.tweetService.CreateTweet(ctx, &tweet); err != nil {
		utils.WriteJsonError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJson(w, http.StatusCreated, tweet)
}
