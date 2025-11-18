package handlers

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/alfinkly/hci-golang-back/database"
	"github.com/alfinkly/hci-golang-back/models"
	"github.com/gofiber/fiber/v3"
)

type PurchaseHandler struct{}

func NewPurchaseHandler() *PurchaseHandler {
	return &PurchaseHandler{}
}

// GetAll returns all purchases
func (h *PurchaseHandler) GetAll(c fiber.Ctx) error {
	query := `
		SELECT id, medicine_id, supplier_id, quantity, unit_price, total_price, 
		       purchase_date, created_at
		FROM purchases
		ORDER BY purchase_date DESC
	`

	var purchases []models.Purchase
	err := database.DB.Select(&purchases, query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch purchases",
		})
	}

	return c.JSON(purchases)
}

// GetByID returns a purchase by ID
func (h *PurchaseHandler) GetByID(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid purchase ID",
		})
	}

	query := `
		SELECT id, medicine_id, supplier_id, quantity, unit_price, total_price, 
		       purchase_date, created_at
		FROM purchases
		WHERE id = $1
	`

	var purchase models.Purchase
	err = database.DB.Get(&purchase, query, id)
	if err == sql.ErrNoRows {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Purchase not found",
		})
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch purchase",
		})
	}

	return c.JSON(purchase)
}

// Create creates a new purchase and updates medicine quantity
func (h *PurchaseHandler) Create(c fiber.Ctx) error {
	var req models.CreatePurchaseRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate required fields
	if req.MedicineID == 0 || req.SupplierID == 0 || req.Quantity <= 0 || req.UnitPrice <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Medicine ID, supplier ID, quantity, and unit price are required",
		})
	}

	totalPrice := float64(req.Quantity) * req.UnitPrice

	// Start transaction
	tx, err := database.DB.Begin()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to start transaction",
		})
	}
	defer tx.Rollback()

	// Insert purchase
	query := `
		INSERT INTO purchases (medicine_id, supplier_id, quantity, unit_price, total_price, 
		                      purchase_date, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, medicine_id, supplier_id, quantity, unit_price, total_price, 
		          purchase_date, created_at
	`

	var purchase models.Purchase
	err = tx.QueryRow(
		query,
		req.MedicineID,
		req.SupplierID,
		req.Quantity,
		req.UnitPrice,
		totalPrice,
		time.Now(),
		time.Now(),
	).Scan(
		&purchase.ID,
		&purchase.MedicineID,
		&purchase.SupplierID,
		&purchase.Quantity,
		&purchase.UnitPrice,
		&purchase.TotalPrice,
		&purchase.PurchaseDate,
		&purchase.CreatedAt,
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create purchase: " + err.Error(),
		})
	}

	// Update medicine quantity
	updateQuery := `
		UPDATE medicines 
		SET quantity = quantity + $1, updated_at = $2
		WHERE id = $3
	`
	_, err = tx.Exec(updateQuery, req.Quantity, time.Now(), req.MedicineID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update medicine quantity",
		})
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to commit transaction",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(purchase)
}

// Delete deletes a purchase
func (h *PurchaseHandler) Delete(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid purchase ID",
		})
	}

	query := `DELETE FROM purchases WHERE id = $1`
	result, err := database.DB.Exec(query, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete purchase",
		})
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Purchase not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Purchase deleted successfully",
	})
}

// SaleHandler handles sale operations
type SaleHandler struct{}

func NewSaleHandler() *SaleHandler {
	return &SaleHandler{}
}

// GetAll returns all sales
func (h *SaleHandler) GetAll(c fiber.Ctx) error {
	query := `
		SELECT id, medicine_id, user_id, quantity, unit_price, total_price, 
		       sale_date, created_at
		FROM sales
		ORDER BY sale_date DESC
	`

	var sales []models.Sale
	err := database.DB.Select(&sales, query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch sales",
		})
	}

	return c.JSON(sales)
}

// GetByID returns a sale by ID
func (h *SaleHandler) GetByID(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid sale ID",
		})
	}

	query := `
		SELECT id, medicine_id, user_id, quantity, unit_price, total_price, 
		       sale_date, created_at
		FROM sales
		WHERE id = $1
	`

	var sale models.Sale
	err = database.DB.Get(&sale, query, id)
	if err == sql.ErrNoRows {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Sale not found",
		})
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch sale",
		})
	}

	return c.JSON(sale)
}

// Create creates a new sale and updates medicine quantity
func (h *SaleHandler) Create(c fiber.Ctx) error {
	var req models.CreateSaleRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate required fields
	if req.MedicineID == 0 || req.Quantity <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Medicine ID and quantity are required",
		})
	}

	// Get user ID from context
	userID, ok := c.Locals("user_id").(int)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User ID not found in context",
		})
	}

	// Start transaction
	tx, err := database.DB.Begin()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to start transaction",
		})
	}
	defer tx.Rollback()

	// Get medicine price and check quantity
	var price float64
	var availableQuantity int
	medicineQuery := `SELECT price, quantity FROM medicines WHERE id = $1`
	err = tx.QueryRow(medicineQuery, req.MedicineID).Scan(&price, &availableQuantity)
	if err == sql.ErrNoRows {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Medicine not found",
		})
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch medicine",
		})
	}

	// Check if enough quantity is available
	if availableQuantity < req.Quantity {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Insufficient quantity available",
		})
	}

	totalPrice := float64(req.Quantity) * price

	// Insert sale
	query := `
		INSERT INTO sales (medicine_id, user_id, quantity, unit_price, total_price, 
		                  sale_date, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, medicine_id, user_id, quantity, unit_price, total_price, 
		          sale_date, created_at
	`

	var sale models.Sale
	err = tx.QueryRow(
		query,
		req.MedicineID,
		userID,
		req.Quantity,
		price,
		totalPrice,
		time.Now(),
		time.Now(),
	).Scan(
		&sale.ID,
		&sale.MedicineID,
		&sale.UserID,
		&sale.Quantity,
		&sale.UnitPrice,
		&sale.TotalPrice,
		&sale.SaleDate,
		&sale.CreatedAt,
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create sale: " + err.Error(),
		})
	}

	// Update medicine quantity
	updateQuery := `
		UPDATE medicines 
		SET quantity = quantity - $1, updated_at = $2
		WHERE id = $3
	`
	_, err = tx.Exec(updateQuery, req.Quantity, time.Now(), req.MedicineID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update medicine quantity",
		})
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to commit transaction",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(sale)
}

// Delete deletes a sale
func (h *SaleHandler) Delete(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid sale ID",
		})
	}

	query := `DELETE FROM sales WHERE id = $1`
	result, err := database.DB.Exec(query, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete sale",
		})
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Sale not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Sale deleted successfully",
	})
}
