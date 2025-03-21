package app

type Application struct {
	Config Config
}

type Config struct {
	Addr string
}

// func (app *Application) Run() error {

// 	// server := &http.Server{
// 	// 	Addr:         app.Config.Addr,
// 	// 	Handler:      r,
// 	// 	WriteTimeout: 30 * time.Second,
// 	// 	ReadTimeout:  10 * time.Second,
// 	// 	IdleTimeout:  time.Minute,
// 	// }

// 	// return server.ListenAndServe()

// }
