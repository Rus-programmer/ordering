package api

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"ordering/logging"
	"ordering/middleware"
	"ordering/services"
	"ordering/util"
	"ordering/validators"

	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our ordering service.
type Server struct {
	config     util.Config
	router     *gin.Engine
	middleware middleware.Middleware
	service    services.Service
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, middleware middleware.Middleware, service services.Service) (*Server, error) {
	router := gin.New()

	if config.Environment != "test" {
		router.Use(logging.GinLogger())
		router.Use(middleware.LogDB())
	}

	server := &Server{
		config:     config,
		router:     router,
		service:    service,
		middleware: middleware,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("role", validators.ValidRole)
		if err != nil {
			return nil, err
		}
	}

	server.setupRouter()
	return server, nil
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) setupRouter() {
	server.router.POST("/users", server.createCustomer)
	server.router.POST("/login", server.login)
	server.router.POST("/tokens/renew_access", server.renewAccessToken)

	authRoutes := server.router.Group("/").Use(server.middleware.Auth())
	authRoutes.GET("/products/:id", server.getProduct)
	authRoutes.GET("/products", server.listProducts)
	authRoutes.PUT("/products/:id", server.updateProduct)
	authRoutes.DELETE("/products/:id", server.deleteProduct)
}
