package handlers

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/alfinkly/hci-golang-back/database"
	"github.com/alfinkly/hci-golang-back/models"
	"github.com/gofiber/fiber/v3"
)

type MedicineHandler struct{}

func NewMedicineHandler() *MedicineHandler {
	return &MedicineHandler{}
}

// GetAll returns all medicines
func (h *MedicineHandler) GetAll(c fiber.Ctx) error {
	query := `
		SELECT id, name, description, manufacturer, price, quantity, expiry_date, 
		       category, requires_prescription, created_at, updated_at
		FROM medicines
		ORDER BY created_at DESC
	`

	var medicines []models.Medicine
	err := database.DB.Select(&medicines, query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch medicines",
		})
	}

	return c.JSON(medicines)
}

// GetByID returns a medicine by ID
func (h *MedicineHandler) GetByID(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid medicine ID",
		})
	}

	query := `
		SELECT id, name, description, manufacturer, price, quantity, expiry_date, 
		       category, requires_prescription, created_at, updated_at
		FROM medicines
		WHERE id = $1
	`

	var medicine models.Medicine
	err = database.DB.Get(&medicine, query, id)
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

	return c.JSON(medicine)
}

// Create creates a new medicine
func (h *MedicineHandler) Create(c fiber.Ctx) error {
	var req models.CreateMedicineRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate required fields
	if req.Name == "" || req.Price <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name and price are required",
		})
	}

	query := `
		INSERT INTO medicines (name, description, manufacturer, price, quantity, expiry_date, 
		                      category, requires_prescription, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id, name, description, manufacturer, price, quantity, expiry_date, 
		          category, requires_prescription, created_at, updated_at
	`

	var medicine models.Medicine
	err := database.DB.QueryRow(
		query,
		req.Name,
		req.Description,
		req.Manufacturer,
		req.Price,
		req.Quantity,
		req.ExpiryDate,
		req.Category,
		req.RequiresPrescription,
		time.Now(),
		time.Now(),
	).Scan(
		&medicine.ID,
		&medicine.Name,
		&medicine.Description,
		&medicine.Manufacturer,
		&medicine.Price,
		&medicine.Quantity,
		&medicine.ExpiryDate,
		&medicine.Category,
		&medicine.RequiresPrescription,
		&medicine.CreatedAt,
		&medicine.UpdatedAt,
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create medicine: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(medicine)
}

// Update updates a medicine
func (h *MedicineHandler) Update(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid medicine ID",
		})
	}

	var req models.UpdateMedicineRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Build dynamic update query
	updates := []string{}
	args := []interface{}{}
	argCount := 1

	if req.Name != nil {
		updates = append(updates, "name = $"+strconv.Itoa(argCount))
		args = append(args, *req.Name)
		argCount++
	}
	if req.Description != nil {
		updates = append(updates, "description = $"+strconv.Itoa(argCount))
		args = append(args, *req.Description)
		argCount++
	}
	if req.Manufacturer != nil {
		updates = append(updates, "manufacturer = $"+strconv.Itoa(argCount))
		args = append(args, *req.Manufacturer)
		argCount++
	}
	if req.Price != nil {
		updates = append(updates, "price = $"+strconv.Itoa(argCount))
		args = append(args, *req.Price)
		argCount++
	}
	if req.Quantity != nil {
		updates = append(updates, "quantity = $"+strconv.Itoa(argCount))
		args = append(args, *req.Quantity)
		argCount++
	}
	if req.ExpiryDate != nil {
		updates = append(updates, "expiry_date = $"+strconv.Itoa(argCount))
		args = append(args, *req.ExpiryDate)
		argCount++
	}
	if req.Category != nil {
		updates = append(updates, "category = $"+strconv.Itoa(argCount))
		args = append(args, *req.Category)
		argCount++
	}
	if req.RequiresPrescription != nil {
		updates = append(updates, "requires_prescription = $"+strconv.Itoa(argCount))
		args = append(args, *req.RequiresPrescription)
		argCount++
	}

	if len(updates) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No fields to update",
		})
	}

	updates = append(updates, "updated_at = $"+strconv.Itoa(argCount))
	args = append(args, time.Now())
	argCount++

	args = append(args, id)

	query := `
		UPDATE medicines 
		SET ` + updates[0]
	for i := 1; i < len(updates); i++ {
		query += ", " + updates[i]
	}
	query += `
		WHERE id = $` + strconv.Itoa(argCount) + `
		RETURNING id, name, description, manufacturer, price, quantity, expiry_date, 
		          category, requires_prescription, created_at, updated_at
	`

	var medicine models.Medicine
	err = database.DB.QueryRow(query, args...).Scan(
		&medicine.ID,
		&medicine.Name,
		&medicine.Description,
		&medicine.Manufacturer,
		&medicine.Price,
		&medicine.Quantity,
		&medicine.ExpiryDate,
		&medicine.Category,
		&medicine.RequiresPrescription,
		&medicine.CreatedAt,
		&medicine.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Medicine not found",
		})
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update medicine: " + err.Error(),
		})
	}

	return c.JSON(medicine)
}

// Delete deletes a medicine
func (h *MedicineHandler) Delete(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid medicine ID",
		})
	}

	query := `DELETE FROM medicines WHERE id = $1`
	result, err := database.DB.Exec(query, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete medicine",
		})
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Medicine not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Medicine deleted successfully",
	})
}
