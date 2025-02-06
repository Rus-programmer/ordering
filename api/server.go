package api

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "ordering/db/sqlc"
	"ordering/logging"
	"ordering/middleware"
	"ordering/token"
	util "ordering/util"
	"ordering/validators"

	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config     util.Config
	router     *gin.Engine
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	router := gin.New()

	if config.Environment != "test" {
		router.Use(logging.GinLogger())
		router.Use(middleware.LogDB(store))
	}

	server := &Server{
		config:     config,
		router:     router,
		store:      store,
		tokenMaker: tokenMaker,
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

	authRoutes := server.router.Group("/").Use(middleware.Auth(server.tokenMaker))
	authRoutes.GET("/products/:id", server.getProduct)
	authRoutes.GET("/products", server.listProducts)
}
