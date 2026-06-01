package main

import (
	"log"

	"github.com/Debasish-87/featurepilot/internal/cache"
	"github.com/Debasish-87/featurepilot/internal/config"
	"github.com/Debasish-87/featurepilot/internal/database"

	environmentHandler "github.com/Debasish-87/featurepilot/internal/environment/handler"
	environmentRepository "github.com/Debasish-87/featurepilot/internal/environment/repository"
	environmentService "github.com/Debasish-87/featurepilot/internal/environment/service"

	healthHandler "github.com/Debasish-87/featurepilot/internal/health/handler"

	organizationHandler "github.com/Debasish-87/featurepilot/internal/organization/handler"
	organizationRepository "github.com/Debasish-87/featurepilot/internal/organization/repository"
	organizationService "github.com/Debasish-87/featurepilot/internal/organization/service"

	projectHandler "github.com/Debasish-87/featurepilot/internal/project/handler"
	projectRepository "github.com/Debasish-87/featurepilot/internal/project/repository"
	projectService "github.com/Debasish-87/featurepilot/internal/project/service"

	featureHandler "github.com/Debasish-87/featurepilot/internal/feature/handler"
	featureRepository "github.com/Debasish-87/featurepilot/internal/feature/repository"
	featureService "github.com/Debasish-87/featurepilot/internal/feature/service"

	evaluationCache "github.com/Debasish-87/featurepilot/internal/evaluation/cache"
	evaluationHandler "github.com/Debasish-87/featurepilot/internal/evaluation/handler"
	evaluationRepository "github.com/Debasish-87/featurepilot/internal/evaluation/repository"
	evaluationService "github.com/Debasish-87/featurepilot/internal/evaluation/service"

	"github.com/Debasish-87/featurepilot/pkg/logger"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	zapLogger, err := logger.New()
	if err != nil {
		log.Fatal(err)
	}
	defer zapLogger.Sync()

	db, err := database.NewPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Organization

	orgRepo := organizationRepository.New(db)

	orgService := organizationService.New(
		orgRepo,
	)

	orgHandler := organizationHandler.New(
		orgService,
	)

	// Project

	projectRepo := projectRepository.New(db)

	projectSvc := projectService.New(
		projectRepo,
	)

	projectHdl := projectHandler.New(
		projectSvc,
	)

	// Environment

	environmentRepo := environmentRepository.New(db)

	environmentSvc := environmentService.New(
		environmentRepo,
	)

	environmentHdl := environmentHandler.New(
		environmentSvc,
	)

	// Redis

	rdb := cache.NewRedis(cfg)

	if err := cache.Ping(rdb); err != nil {
		log.Fatal(err)
	}

	evalCache := evaluationCache.New(
		rdb,
	)

	// Features

	featureRepo := featureRepository.New(db)

	featureSvc := featureService.New(
		featureRepo,
		environmentRepo,
		evalCache,
	)

	featureHdl := featureHandler.New(
		featureSvc,
	)

	// Evaluation

	evaluationRepo := evaluationRepository.New(
		db,
	)

	evaluationSvc := evaluationService.New(
		evaluationRepo,
		evalCache,
	)

	evaluationHdl := evaluationHandler.New(
		evaluationSvc,
	)
	// Router

	router := gin.New()

	if err := router.SetTrustedProxies(nil); err != nil {
		log.Fatal(err)
	}

	router.Use(gin.Recovery())

	health := healthHandler.NewHealthHandler()

	router.GET("/health", health.Health)

	// Organization Routes

	router.GET(
		"/api/v1/organizations/:id",
		orgHandler.GetByID,
	)

	router.GET(
		"/api/v1/organizations",
		orgHandler.List,
	)

	router.POST(
		"/api/v1/organizations",
		orgHandler.Create,
	)

	router.DELETE(
		"/api/v1/organizations/:id",
		orgHandler.Delete,
	)

	// Project Routes

	router.GET(
		"/api/v1/projects/:id",
		projectHdl.GetByID,
	)

	router.GET(
		"/api/v1/projects",
		projectHdl.List,
	)

	router.POST(
		"/api/v1/projects",
		projectHdl.Create,
	)

	router.DELETE(
		"/api/v1/projects/:id",
		projectHdl.Delete,
	)

	// Environment Routes

	router.GET(
		"/api/v1/environments/:id",
		environmentHdl.GetByID,
	)

	router.GET(
		"/api/v1/environments",
		environmentHdl.List,
	)

	router.POST(
		"/api/v1/environments",
		environmentHdl.Create,
	)

	router.DELETE(
		"/api/v1/environments/:id",
		environmentHdl.Delete,
	)

	// Features Routes

	router.GET(
		"/api/v1/features/:id",
		featureHdl.GetByID,
	)

	router.GET(
		"/api/v1/features",
		featureHdl.List,
	)

	router.POST(
		"/api/v1/features",
		featureHdl.Create,
	)

	router.DELETE(
		"/api/v1/features/:id",
		featureHdl.Delete,
	)

	router.PATCH(
		"/api/v1/features/:id/enable",
		featureHdl.Enable,
	)

	router.PATCH(
		"/api/v1/features/:id/disable",
		featureHdl.Disable,
	)

	// Evaluates Routes

	router.POST(
		"/api/v1/evaluate",
		evaluationHdl.Evaluate,
	)

	zapLogger.Info("featurepilot started")

	if err := router.Run(":" + cfg.HTTPPort); err != nil {
		log.Fatal(err)
	}
}
