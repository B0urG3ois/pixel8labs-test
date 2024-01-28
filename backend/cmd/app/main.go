package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	_githubHandler "pixel8labs-test/backend/api/handler/github"
	_middleware "pixel8labs-test/backend/api/middleware/helper"
	"pixel8labs-test/backend/configs/app"
	_dbHandler "pixel8labs-test/backend/configs/db"
	_githubConfig "pixel8labs-test/backend/configs/github"
	"pixel8labs-test/backend/internal/entity"
	_githubRepo "pixel8labs-test/backend/internal/repo/github"
	_githubService "pixel8labs-test/backend/internal/service/github"
	_githubUsecase "pixel8labs-test/backend/internal/usecase/github"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
	"golang.org/x/oauth2"
)

func main() {
	var (
		ctx = context.Background()
		err error
	)

	// Load env file for dev env.
	/*err = godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
		return
	}*/

	// Initialize Config.
	config := app.InitAppConfig()
	if err != nil {
		log.Fatalf("Failed to init config: %v", err)
		return
	}

	// Initialize OAuth Config.
	oAuthConfig := _githubConfig.InitOAuthConfig(config)

	// Initialize Database.
	db, err := sql.Open("sqlite3", config.DatabaseConfig.Path)
	if err != nil {
		log.Fatalf("Failed to init database: %v", err)
		return
	}
	defer db.Close()

	// Init tables.
	if err := _dbHandler.New(db); err != nil {
		log.Fatalf("failed init tables %v", err)
		return
	}

	startApp(ctx, config, oAuthConfig, db)
}

func startApp(
	ctx context.Context,
	config entity.Config,
	oAuthConfig oauth2.Config,
	db *sql.DB,
) {
	// Init Services.
	githubSvc := _githubService.New(oAuthConfig, config)

	// Init Repo.
	githubRepo := _githubRepo.New(githubSvc, db)

	// Init Usecase.
	githubUC := _githubUsecase.New(githubRepo)

	// Init Handler.
	githubHandler := _githubHandler.New(githubUC)

	r := newRoutes(githubHandler)

	// Run apps.
	log.Printf("Application started at port :%s", config.ApplicationConfig.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", config.ApplicationConfig.Port), r) //nolint:all
}

func newRoutes(githubHandler *_githubHandler.Handler) *chi.Mux {
	r := chi.NewRouter()

	// Create a new CORS middleware with default options.
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "https://pixel8labs-assignment.vercel.app"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	})

	r.Use(corsMiddleware.Handler)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/api", func(r chi.Router) {
		r.Post("/login", _middleware.GenericMiddleware(githubHandler.GetAuthCode))
		r.Post("/logout", _middleware.GenericMiddleware(githubHandler.Logout))
		r.Get("/callback", _middleware.GenericMiddleware(githubHandler.Callback))

		r.Route("/v1", func(r chi.Router) {
			r.Get("/user", _middleware.AuthMiddleware(_middleware.GenericMiddleware(githubHandler.GetUserInfo)))
			r.Get("/repositories", _middleware.AuthMiddleware(_middleware.GenericMiddleware(githubHandler.GetRepositories)))
		})
	})

	return r
}
