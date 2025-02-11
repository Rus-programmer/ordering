package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"ordering/logging"
	"ordering/middleware"
	"ordering/services"
	"ordering/util"
	"ordering/validators"
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
		err = v.RegisterValidation("orderStatus", validators.ValidOrderStatus)
		err = v.RegisterValidation("notEmptyArray", validators.NotEmptyArrayValidator)
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
	server.router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	server.router.POST("/customers", server.createCustomer)
	server.router.POST("/login", server.login)
	server.router.POST("/renew_access", server.renewAccessToken)

	authRoutes := server.router.Group("/").Use(server.middleware.Auth())
	authRoutes.GET("/products/:id", server.getProduct)
	authRoutes.GET("/products", server.listProducts)
	authRoutes.PUT("/products/:id", server.updateProduct)
	authRoutes.DELETE("/products/:id", server.deleteProduct)
	authRoutes.POST("/products", server.createProduct)

	authRoutes.GET("/orders/:id", server.getOrder)
	authRoutes.GET("/orders", server.listOrders)
	authRoutes.DELETE("/orders/:id", server.deleteOrder)
	authRoutes.POST("/orders", server.createOrder)
	authRoutes.PUT("/orders/:id", server.updateOrder)

	authRoutes.GET("/metrics", server.getMetrics)
}
