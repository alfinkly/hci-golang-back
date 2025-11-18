package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/alfinkly/hci-golang-back/config"
	"github.com/alfinkly/hci-golang-back/database"
	"github.com/alfinkly/hci-golang-back/handlers"
	"github.com/alfinkly/hci-golang-back/middleware"
	"github.com/gofiber/fiber/v3"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Connect to database
	if err := database.Connect(cfg); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	// Initialize database schema
	if err := database.InitSchema(); err != nil {
		log.Fatalf("Failed to initialize database schema: %v", err)
	}

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: "Pharmacy Backend API",
	})

	// Global middleware
	app.Use(middleware.CORSMiddleware())
	app.Use(middleware.LoggingMiddleware())

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(cfg)
	medicineHandler := handlers.NewMedicineHandler()
	supplierHandler := handlers.NewSupplierHandler()
	purchaseHandler := handlers.NewPurchaseHandler()
	saleHandler := handlers.NewSaleHandler()

	// Public routes
	api := app.Group("/api")

	// Auth routes (public)
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	// Protected routes - all require JWT authentication
	protected := api.Group("/", middleware.JWTMiddleware(cfg))

	// User profile
	protected.Get("/profile", authHandler.GetProfile)

	// Medicine routes
	medicines := protected.Group("/medicines")
	medicines.Get("/", medicineHandler.GetAll)
	medicines.Get("/:id", medicineHandler.GetByID)
	medicines.Post("/", medicineHandler.Create)
	medicines.Put("/:id", medicineHandler.Update)
	medicines.Delete("/:id", medicineHandler.Delete)

	// Supplier routes
	suppliers := protected.Group("/suppliers")
	suppliers.Get("/", supplierHandler.GetAll)
	suppliers.Get("/:id", supplierHandler.GetByID)
	suppliers.Post("/", supplierHandler.Create)
	suppliers.Put("/:id", supplierHandler.Update)
	suppliers.Delete("/:id", supplierHandler.Delete)

	// Purchase routes
	purchases := protected.Group("/purchases")
	purchases.Get("/", purchaseHandler.GetAll)
	purchases.Get("/:id", purchaseHandler.GetByID)
	purchases.Post("/", purchaseHandler.Create)
	purchases.Delete("/:id", purchaseHandler.Delete)

	// Sale routes
	sales := protected.Group("/sales")
	sales.Get("/", saleHandler.GetAll)
	sales.Get("/:id", saleHandler.GetByID)
	sales.Post("/", saleHandler.Create)
	sales.Delete("/:id", saleHandler.Delete)

	// Health check endpoint
	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Pharmacy API is running",
		})
	})

	// Handle graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-quit
		log.Println("Shutting down server...")
		if err := app.Shutdown(); err != nil {
			log.Fatalf("Server forced to shutdown: %v", err)
		}
	}()

	// Start server
	log.Printf("Starting server on port %s", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
