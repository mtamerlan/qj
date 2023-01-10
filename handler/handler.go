package handler

import (
	"todo/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{service: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			item := lists.Group(":id/item")
			{
				item.POST("/", h.createItem)
				item.GET("/", h.getAllItems)
				item.GET("/:item-id", h.getItemById)
				item.PUT("/item-id", h.updateItem)
				item.DELETE("/item-id", h.deleteItem)
			}
		}
	}
	return router
}
