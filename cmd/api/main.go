package main

import (
	"log"
	"net/http"
	"os"

	"github.com/alekzandr-s/dynamic-form-engine/internal/config"
	"github.com/alekzandr-s/dynamic-form-engine/internal/database"
	"github.com/alekzandr-s/dynamic-form-engine/internal/handler"
	"github.com/alekzandr-s/dynamic-form-engine/internal/repository/postgres"
	"github.com/alekzandr-s/dynamic-form-engine/internal/service"
	"github.com/alekzandr-s/dynamic-form-engine/internal/validator"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	cfg := config.Load()

	pool, err := database.NewPostgresPool(cfg.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	definitionRepo := postgres.NewFormDefinitionRepository(pool)

	versionRepo := postgres.NewFormTemplateVersionRepository(pool)

	submissionRepo := postgres.NewFormSubmissionRepository(pool)

	validator := validator.NewJSONSchemaValidator()

	definitionService := service.NewDefinitionService(definitionRepo)

	versionService := service.NewVersionService(versionRepo)

	engineService := service.NewTemplateEngineService(versionRepo, submissionRepo, validator)

	definitionHandler := handler.NewDefinitionHandler(definitionService)

	versionHandler := handler.NewVersionHandler(versionService)

	submissionHandler := handler.NewSubmissionHandler(engineService)

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowCredentials: false,
	}))

	r.Post("/definitions", definitionHandler.Create)

	r.Post("/definitions/{id}/versions", versionHandler.Create)

	r.Post("/definitions/{id}/publish", versionHandler.Publish)

	r.Get("/forms/{id}", submissionHandler.GetPublished)

	r.Post("/forms/{id}/submissions", submissionHandler.Submit)

	r.Get("/definitions", definitionHandler.List)

	r.Get("/health", handler.Health)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on :%s", port)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		panic(err)
	}
}
