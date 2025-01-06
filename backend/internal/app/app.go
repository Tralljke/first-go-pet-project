package app

import (
	"first-go-pet-project/backend/internal/config"
	"first-go-pet-project/backend/internal/infrastructure"
	"first-go-pet-project/backend/internal/routes"
	"first-go-pet-project/backend/internal/usecases"
	"fmt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type App struct {
	Config     *config.Config
	DB         *gorm.DB
	Repository *infrastructure.MySQLStudentRepository
	Service    *usecases.StudentService
	Router     *http.ServeMux
}

func NewApp(configPath string) (*App, error) {

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}

	dsnMigrate := fmt.Sprintf("mysql://%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName)

	dsnGORM := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName)

	infrastructure.RunMigrations(dsnMigrate)

	db := infrastructure.NewMySQLConnection(dsnGORM)

	repo := infrastructure.NewMySQLStudentRepository(db)
	service := usecases.NewStudentService(repo)

	router := http.NewServeMux()
	routes.RegisterRoutes(router, service)

	return &App{
		Config:     cfg,
		DB:         db,
		Repository: repo,
		Service:    service,
		Router:     router,
	}, nil
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Разрешить все источники, можно заменить "*" на конкретный домен
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (a *App) Start() {
	fmt.Println("Server is running on http://localhost:8080")

	handlerWithCORS := corsMiddleware(a.Router)
	if err := http.ListenAndServe(":8080", handlerWithCORS); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
