package main

import (
	"ecommerce-service/internal/db"
	"ecommerce-service/internal/handlers"
	"ecommerce-service/internal/repository"
	"ecommerce-service/internal/routes"
	service "ecommerce-service/internal/services"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	db.InitDB()

	// Wire Cart Module
	cartRepo := repository.NewCartRepository(db.DB)
	cartService := service.NewCartService(cartRepo)
	cartHandler := handlers.NewCartHandler(cartService)

	// Wire Wishlist Module
	wishlistRepo := repository.NewWishlistRepository(db.DB)
	wishlistService := service.NewWishlistService(wishlistRepo)
	wishlistHandler := handlers.NewWishlistHandler(wishlistService)

	router := gin.Default()

	routes.RegisterRoutes(router, cartHandler, wishlistHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
