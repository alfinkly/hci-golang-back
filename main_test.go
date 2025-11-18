package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/alfinkly/hci-golang-back/config"
	"github.com/alfinkly/hci-golang-back/database"
	"github.com/alfinkly/hci-golang-back/handlers"
	"github.com/alfinkly/hci-golang-back/middleware"
	"github.com/alfinkly/hci-golang-back/models"
	"github.com/gofiber/fiber/v3"
)

var testApp *fiber.App
var testToken string

func setupTestApp() {
	cfg := config.Load()
	
	// Connect to test database
	if err := database.Connect(cfg); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	
	// Initialize schema
	if err := database.InitSchema(); err != nil {
		log.Fatalf("Failed to initialize schema: %v", err)
	}

	testApp = fiber.New()
	testApp.Use(middleware.CORSMiddleware())

	authHandler := handlers.NewAuthHandler(cfg)
	medicineHandler := handlers.NewMedicineHandler()
	
	api := testApp.Group("/api")
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)
	
	protected := api.Group("/", middleware.JWTMiddleware(cfg))
	protected.Get("/profile", authHandler.GetProfile)
	
	medicines := protected.Group("/medicines")
	medicines.Get("/", medicineHandler.GetAll)
	medicines.Post("/", medicineHandler.Create)
}

func TestHealthEndpoint(t *testing.T) {
	if testApp == nil {
		setupTestApp()
	}
	
	testApp.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	req, _ := http.NewRequest("GET", "/health", nil)
	resp, err := testApp.Test(req)
	if err != nil {
		t.Fatalf("Failed to test /health: %v", err)
	}
	
	if resp.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

func TestRegisterEndpoint(t *testing.T) {
	if testApp == nil {
		setupTestApp()
	}

	registerData := models.RegisterRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
		Role:     "user",
	}
	
	jsonData, _ := json.Marshal(registerData)
	req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	
	resp, err := testApp.Test(req)
	if err != nil {
		t.Fatalf("Failed to test /api/auth/register: %v", err)
	}
	
	if resp.StatusCode != 201 && resp.StatusCode != 500 {
		body, _ := io.ReadAll(resp.Body)
		t.Logf("Response body: %s", string(body))
		t.Errorf("Expected status 201 or 500 (if user exists), got %d", resp.StatusCode)
	}
}

func TestLoginEndpoint(t *testing.T) {
	if testApp == nil {
		setupTestApp()
	}

	loginData := models.LoginRequest{
		Username: "testuser",
		Password: "password123",
	}
	
	jsonData, _ := json.Marshal(loginData)
	req, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	
	resp, err := testApp.Test(req)
	if err != nil {
		t.Fatalf("Failed to test /api/auth/login: %v", err)
	}
	
	if resp.StatusCode == 200 {
		var loginResp models.LoginResponse
		body, _ := io.ReadAll(resp.Body)
		json.Unmarshal(body, &loginResp)
		testToken = loginResp.Token
		t.Logf("Login successful, token received")
	}
}

func TestCreateMedicine(t *testing.T) {
	if testApp == nil {
		setupTestApp()
	}
	
	if testToken == "" {
		t.Skip("No token available, skipping test")
	}

	medicineData := models.CreateMedicineRequest{
		Name:                 "Test Medicine",
		Description:          "Test Description",
		Manufacturer:         "Test Manufacturer",
		Price:                100.50,
		Quantity:             50,
		ExpiryDate:           time.Now().AddDate(1, 0, 0),
		Category:             "Test Category",
		RequiresPrescription: false,
	}
	
	jsonData, _ := json.Marshal(medicineData)
	req, _ := http.NewRequest("POST", "/api/medicines", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+testToken)
	
	resp, err := testApp.Test(req)
	if err != nil {
		t.Fatalf("Failed to test /api/medicines POST: %v", err)
	}
	
	if resp.StatusCode != 201 {
		body, _ := io.ReadAll(resp.Body)
		t.Logf("Response body: %s", string(body))
		t.Errorf("Expected status 201, got %d", resp.StatusCode)
	}
}

func TestGetMedicines(t *testing.T) {
	if testApp == nil {
		setupTestApp()
	}
	
	if testToken == "" {
		t.Skip("No token available, skipping test")
	}

	req, _ := http.NewRequest("GET", "/api/medicines", nil)
	req.Header.Set("Authorization", "Bearer "+testToken)
	
	resp, err := testApp.Test(req)
	if err != nil {
		t.Fatalf("Failed to test /api/medicines GET: %v", err)
	}
	
	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		t.Logf("Response body: %s", string(body))
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}
