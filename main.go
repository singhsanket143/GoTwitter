package main

import (
	dbConfig "GoTwitter/config/db"
	config "GoTwitter/config/env"
	db "GoTwitter/db/repository"
	"GoTwitter/handlers"
	"GoTwitter/router"
	"GoTwitter/services"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

func main() {

	// Load environment variables
	config.LoadEnv()

	app := fx.New(
		fx.Provide(
			dbConfig.SetupNewDbConn,

			fx.Annotate(
				db.NewTweetsStore,
				fx.As(new(db.TweetsRepository)), // 👈 Important line
			),

			db.NewUsersStore,

			services.NewTweetService,
			handlers.NewTweetHandler,
			router.NewTweetRouter,
			router.Mount,
		),
		fx.Invoke(func(r *chi.Mux) {
			http.ListenAndServe(":8080", r)
		}),
	)

	app.Run()
}
