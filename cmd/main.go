package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tamaaa13/fastcampus/internal/configs"
	"github.com/tamaaa13/fastcampus/internal/handlers/memberships"
	"github.com/tamaaa13/fastcampus/internal/handlers/posts"
	membershipRepo "github.com/tamaaa13/fastcampus/internal/repository/memberships"
	postsRepo "github.com/tamaaa13/fastcampus/internal/repository/posts"
	membershipSvc "github.com/tamaaa13/fastcampus/internal/service/memberships"
	postsSvc "github.com/tamaaa13/fastcampus/internal/service/posts"
	"github.com/tamaaa13/fastcampus/pkg/internalsql"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found")
	} else {
		log.Println("Environment Variables Loaded")
	}

	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder(
			[]string{"./internal/configs/"},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("Gagal inisiasi config", err)
	}

	cfg = configs.Get()
	log.Println("=== CONFIG LOADED ===")
	log.Println("config", cfg)
	log.Println("Service Port:", cfg.Service.Port)
	log.Println("JWT Secret:", cfg.Service.SecretJWT)
	log.Println("Database DSN:", cfg.Database.DataSourceName)

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Gagal inisiasi database", err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Memberships
	membershipRepo := membershipRepo.NewRepository(db)
	membershipSvc := membershipSvc.NewService(cfg, membershipRepo)
	membershipHandler := memberships.NewHandler(r, membershipSvc)
	membershipHandler.RegisterRoute()

	// Posts
	postRepo := postsRepo.NewRepository(db)
	postSvc := postsSvc.NewService(cfg, postRepo)
	postsHandler := posts.NewHandler(r, postSvc)
	postsHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
