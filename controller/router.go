package controller

import (
	"github.com/gin-gonic/gin"
	"http-type-one/controller/middleware"
	"http-type-one/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// InitRouter just init our routes via Gin Gonic
func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	//CORS middleware
	router.Use(middleware.CORSMiddleware())

	// Router group for Accounts
	accounts := router.Group("/accounts")
	{
		//Get all
		accounts.GET("/", h.GetAllAccounts)

		// Get by ID
		accounts.GET("/:id", h.GetAccount)

		//Post to create new
		accounts.POST("/", h.CreateAccounts)

		//Put to update by ID
		accounts.PUT("/", h.UpdateAccount)

		// Delete by ID
		accounts.DELETE("/:id", h.DeleteAccount)
	}

	// Router group for Account Integration
	accountIntegration := router.Group("/account-integration")
	{
		//Get all
		accountIntegration.GET("/", h.GetAllAccountIntegration)

		// Get by ID
		accountIntegration.GET("/:id", h.GetAccountIntegration)

		//Post to create new
		accountIntegration.POST("/", h.CreateAccountIntegration)

		//Put to update by ID
		accountIntegration.PUT("/:id", h.UpdateAccountIntegration)

		// Delete by ID
		accountIntegration.DELETE("/:id", h.DeleteAccountIntegration)
	}

	return router
}
