package router

import (
	"GoTwitter/handlers"

	"github.com/go-chi/chi/v5"
)

type TweetRouter struct {
	TweetHandler *handlers.TweetHandler
}

func NewTweetRouter(tweetHandler *handlers.TweetHandler) Routes {
	return &TweetRouter{TweetHandler: tweetHandler}
}

func (tweetRouter *TweetRouter) Register(r chi.Router) {
	r.Route("/tweets", func(r chi.Router) {
		r.Post("/", tweetRouter.TweetHandler.CreateTweetHandler)
	})
}
