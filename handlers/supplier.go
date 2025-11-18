package handlers

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/alfinkly/hci-golang-back/database"
	"github.com/alfinkly/hci-golang-back/models"
	"github.com/gofiber/fiber/v3"
)

type SupplierHandler struct{}

func NewSupplierHandler() *SupplierHandler {
	return &SupplierHandler{}
}

// GetAll returns all suppliers
func (h *SupplierHandler) GetAll(c fiber.Ctx) error {
	query := `
		SELECT id, name, contact_person, phone, email, address, created_at, updated_at
		FROM suppliers
		ORDER BY created_at DESC
	`

	var suppliers []models.Supplier
	err := database.DB.Select(&suppliers, query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch suppliers",
		})
	}

	return c.JSON(suppliers)
}

// GetByID returns a supplier by ID
func (h *SupplierHandler) GetByID(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid supplier ID",
		})
	}

	query := `
		SELECT id, name, contact_person, phone, email, address, created_at, updated_at
		FROM suppliers
		WHERE id = $1
	`

	var supplier models.Supplier
	err = database.DB.Get(&supplier, query, id)
	if err == sql.ErrNoRows {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Supplier not found",
		})
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch supplier",
		})
	}

	return c.JSON(supplier)
}

// Create creates a new supplier
func (h *SupplierHandler) Create(c fiber.Ctx) error {
	var req models.CreateSupplierRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate required fields
	if req.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name is required",
		})
	}

	query := `
		INSERT INTO suppliers (name, contact_person, phone, email, address, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, name, contact_person, phone, email, address, created_at, updated_at
	`

	var supplier models.Supplier
	err := database.DB.QueryRow(
		query,
		req.Name,
		req.ContactPerson,
		req.Phone,
		req.Email,
		req.Address,
		time.Now(),
		time.Now(),
	).Scan(
		&supplier.ID,
		&supplier.Name,
		&supplier.ContactPerson,
		&supplier.Phone,
		&supplier.Email,
		&supplier.Address,
		&supplier.CreatedAt,
		&supplier.UpdatedAt,
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create supplier: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(supplier)
}

// Update updates a supplier
func (h *SupplierHandler) Update(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid supplier ID",
		})
	}

	var req models.UpdateSupplierRequest
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
	if req.ContactPerson != nil {
		updates = append(updates, "contact_person = $"+strconv.Itoa(argCount))
		args = append(args, *req.ContactPerson)
		argCount++
	}
	if req.Phone != nil {
		updates = append(updates, "phone = $"+strconv.Itoa(argCount))
		args = append(args, *req.Phone)
		argCount++
	}
	if req.Email != nil {
		updates = append(updates, "email = $"+strconv.Itoa(argCount))
		args = append(args, *req.Email)
		argCount++
	}
	if req.Address != nil {
		updates = append(updates, "address = $"+strconv.Itoa(argCount))
		args = append(args, *req.Address)
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
		UPDATE suppliers 
		SET ` + updates[0]
	for i := 1; i < len(updates); i++ {
		query += ", " + updates[i]
	}
	query += `
		WHERE id = $` + strconv.Itoa(argCount) + `
		RETURNING id, name, contact_person, phone, email, address, created_at, updated_at
	`

	var supplier models.Supplier
	err = database.DB.QueryRow(query, args...).Scan(
		&supplier.ID,
		&supplier.Name,
		&supplier.ContactPerson,
		&supplier.Phone,
		&supplier.Email,
		&supplier.Address,
		&supplier.CreatedAt,
		&supplier.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Supplier not found",
		})
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update supplier: " + err.Error(),
		})
	}

	return c.JSON(supplier)
}

// Delete deletes a supplier
func (h *SupplierHandler) Delete(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid supplier ID",
		})
	}

	query := `DELETE FROM suppliers WHERE id = $1`
	result, err := database.DB.Exec(query, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete supplier",
		})
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Supplier not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Supplier deleted successfully",
	})
}
