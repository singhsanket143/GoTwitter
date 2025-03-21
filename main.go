package main

import (
	dbConfig "GoTwitter/config/db"
	config "GoTwitter/config/env"
	db "GoTwitter/db/repository"
	"GoTwitter/handlers"
	"GoTwitter/router"
	"GoTwitter/services"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func main() {

	// Load environment variables
	config.LoadEnv()

	app := fx.New(

		fx.WithLogger(func() fxevent.Logger {
			return &fxevent.ConsoleLogger{W: os.Stdout}
		}),

		fx.Provide(
			dbConfig.SetupNewDbConn,

			fx.Annotate(
				db.NewTweetsStore,
				fx.As(new(db.TweetsRepository)),
			),

			fx.Annotate(
				db.NewTagsStore,
				fx.As(new(db.TagsRepository)),
			),

			fx.Annotate(
				db.NewTweetTagsStore,
				fx.As(new(db.TweetTagsRepository)),
			),

			db.NewUsersStore,

			services.NewTweetService,
			handlers.NewTweetHandler,
			router.NewTweetRouter,
			router.Mount,
		),

		fx.Invoke(func(dot fx.DotGraph) {
			file, err := os.Create("fx-graph.dot") // brew install graphviz && dot -Tpng fx-graph.dot -o fx-graph.png

			if err != nil {
				log.Println("Failed to create graph file:", err)
				return
			}
			defer file.Close()

			_, err = fmt.Fprintf(file, "%s", dot)
			if err != nil {
				log.Println("Failed to write DOT graph:", err)
			} else {
				log.Println("DOT graph written to fx-graph.dot")
			}
		}),

		// Start HTTP server
		fx.Invoke(func(lc fx.Lifecycle, r *chi.Mux) {
			server := &http.Server{
				Addr:    ":8080",
				Handler: r,
			}

			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					log.Println("Server starting on :8080")
					go server.ListenAndServe()
					return nil
				},
				OnStop: func(ctx context.Context) error {
					log.Println("Server shutting down...")
					return server.Shutdown(ctx)
				},
			})
		}),
	)

	app.Run()
}
