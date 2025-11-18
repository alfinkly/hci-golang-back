package models

import (
	"time"
)

type User struct {
	ID           int       `json:"id" db:"id"`
	Username     string    `json:"username" db:"username"`
	Email        string    `json:"email" db:"email"`
	PasswordHash string    `json:"-" db:"password_hash"`
	Role         string    `json:"role" db:"role"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type Medicine struct {
	ID                   int       `json:"id" db:"id"`
	Name                 string    `json:"name" db:"name"`
	Description          string    `json:"description" db:"description"`
	Manufacturer         string    `json:"manufacturer" db:"manufacturer"`
	Price                float64   `json:"price" db:"price"`
	Quantity             int       `json:"quantity" db:"quantity"`
	ExpiryDate           time.Time `json:"expiry_date" db:"expiry_date"`
	Category             string    `json:"category" db:"category"`
	RequiresPrescription bool      `json:"requires_prescription" db:"requires_prescription"`
	CreatedAt            time.Time `json:"created_at" db:"created_at"`
	UpdatedAt            time.Time `json:"updated_at" db:"updated_at"`
}

type Supplier struct {
	ID            int       `json:"id" db:"id"`
	Name          string    `json:"name" db:"name"`
	ContactPerson string    `json:"contact_person" db:"contact_person"`
	Phone         string    `json:"phone" db:"phone"`
	Email         string    `json:"email" db:"email"`
	Address       string    `json:"address" db:"address"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

type Purchase struct {
	ID           int       `json:"id" db:"id"`
	MedicineID   int       `json:"medicine_id" db:"medicine_id"`
	SupplierID   int       `json:"supplier_id" db:"supplier_id"`
	Quantity     int       `json:"quantity" db:"quantity"`
	UnitPrice    float64   `json:"unit_price" db:"unit_price"`
	TotalPrice   float64   `json:"total_price" db:"total_price"`
	PurchaseDate time.Time `json:"purchase_date" db:"purchase_date"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

type Sale struct {
	ID         int       `json:"id" db:"id"`
	MedicineID int       `json:"medicine_id" db:"medicine_id"`
	UserID     int       `json:"user_id" db:"user_id"`
	Quantity   int       `json:"quantity" db:"quantity"`
	UnitPrice  float64   `json:"unit_price" db:"unit_price"`
	TotalPrice float64   `json:"total_price" db:"total_price"`
	SaleDate   time.Time `json:"sale_date" db:"sale_date"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

// Request/Response DTOs
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type CreateMedicineRequest struct {
	Name                 string    `json:"name"`
	Description          string    `json:"description"`
	Manufacturer         string    `json:"manufacturer"`
	Price                float64   `json:"price"`
	Quantity             int       `json:"quantity"`
	ExpiryDate           time.Time `json:"expiry_date"`
	Category             string    `json:"category"`
	RequiresPrescription bool      `json:"requires_prescription"`
}

type UpdateMedicineRequest struct {
	Name                 *string    `json:"name,omitempty"`
	Description          *string    `json:"description,omitempty"`
	Manufacturer         *string    `json:"manufacturer,omitempty"`
	Price                *float64   `json:"price,omitempty"`
	Quantity             *int       `json:"quantity,omitempty"`
	ExpiryDate           *time.Time `json:"expiry_date,omitempty"`
	Category             *string    `json:"category,omitempty"`
	RequiresPrescription *bool      `json:"requires_prescription,omitempty"`
}

type CreateSupplierRequest struct {
	Name          string `json:"name"`
	ContactPerson string `json:"contact_person"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	Address       string `json:"address"`
}

type UpdateSupplierRequest struct {
	Name          *string `json:"name,omitempty"`
	ContactPerson *string `json:"contact_person,omitempty"`
	Phone         *string `json:"phone,omitempty"`
	Email         *string `json:"email,omitempty"`
	Address       *string `json:"address,omitempty"`
}

type CreatePurchaseRequest struct {
	MedicineID int     `json:"medicine_id"`
	SupplierID int     `json:"supplier_id"`
	Quantity   int     `json:"quantity"`
	UnitPrice  float64 `json:"unit_price"`
}

type CreateSaleRequest struct {
	MedicineID int `json:"medicine_id"`
	Quantity   int `json:"quantity"`
}
