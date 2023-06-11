package config

import (
	config "lore_project/configs"
	"lore_project/internal/models"
	"lore_project/internal/routes"
	"lore_project/internal/services"
	"lore_project/pkg/database"
	logger "lore_project/pkg/logging"

	"github.com/gin-gonic/gin"
)

// Server represents the HTTP server.
type Server struct {
	router      *gin.Engine
	db          *database.DB
	userService *services.UserService
}

// NewServer creates a new instance of the HTTP server.
func NewServer() *Server {
	// Initialize the Gin router
	router := gin.Default()

	cfg := config.LoadConfig()

	// Initialize the database connection
	db, err := database.NewDB(*cfg)

	if err != nil {
		logger.GetLogger().Errorf("Failed to connect to the database: %v", err)
	}
	// Create migrate table
	db.AutoMigrate(&models.User{})

	// Create a new instance of the user service
	userService := services.NewUserService(db)

	// Create a new instance of the server
	server := &Server{
		router:      router,
		db:          db,
		userService: userService,
	}
	//Initial Router
	server.InitializeRoutes()

	return server

}

// Start starts the HTTP server on the specified address.
func (s *Server) Start(address string) error {
	// Start the server
	err := s.router.Run(address)
	if err != nil {
		return err
	}

	return nil
}

// Run starts the HTTP server on the specified address.
// It logs any error that occurs during server startup.
func Run(address string) {

	server := NewServer()

	logger.GetLogger().Infof("Starting HTTP server on %s", address)

	err := server.Start(address)
	if err != nil {
		logger.GetLogger().Errorf("Failed to start HTTP server: %v", err)
	}
}

func (r *Server) InitializeRoutes() {
	routes.RegisterRoutes(r.router.Group("/api"), r.userService)
}
