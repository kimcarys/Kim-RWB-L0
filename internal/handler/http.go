package handler

import (
	"log"
	"net/http"
	"order-service/internal/cache"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	cache *cache.Cache
}

func NewHandler(cache *cache.Cache) *Handler {
	return &Handler{
		cache: cache,
	}
}

func (h *Handler) GetOrderByID(c *gin.Context) {
	orderID := c.Param("id")
	if orderID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Order ID is required"})
		return
	}

	order, exists := h.cache.Get(orderID)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *Handler) GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Order Service",
	})
}

func StartServer(cache *cache.Cache) {
	handler := NewHandler(cache)
	
	router := gin.Default()
	
	// Настраиваем шаблоны (позже добавим HTML)
	router.LoadHTMLGlob("templates/*")
	
	// Статические файлы
	router.Static("/static", "./static")
	
	// Маршруты
	router.GET("/", handler.GetIndex)
	router.GET("/order/:id", handler.GetOrderByID)
	
	log.Println("Starting HTTP server on :8080")
	router.Run(":8080")
}