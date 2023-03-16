package handler

import (
	"github.com/gin-gonic/gin"
	"iiiproject/pkg/service"
)

type Handler struct {
	servicec *service.Service
}

func NewHandler(servicec *service.Service) *Handler {
	return &Handler{servicec: servicec}
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
			lists.POST("/", h.CreateList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)
		}

		items := lists.Group(":id/items")
		{
			items.POST("/", h.CreateItem)
			items.GET("/", h.getAllItems)
			items.GET("/:item_id", h.getItemById)
			items.PUT("/:item_id", h.updateItem)
			items.DELETE("/item_id ", h.deleteItem)
		}

	}

	return router

}
