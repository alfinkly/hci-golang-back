package database

import (
	"log"

	"github.com/alfinkly/hci-golang-back/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

// Connect initializes the database connection
func Connect(cfg *config.Config) error {
	var err error
	DB, err = sqlx.Connect("postgres", cfg.GetDBConnectionString())
	if err != nil {
		return err
	}

	// Test the connection
	if err = DB.Ping(); err != nil {
		return err
	}

	log.Println("Database connected successfully")
	return nil
}

// Close closes the database connection
func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}

// InitSchema creates the database schema
func InitSchema() error {
	schema := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(100) UNIQUE NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		password_hash VARCHAR(255) NOT NULL,
		role VARCHAR(50) DEFAULT 'user',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS suppliers (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		contact_person VARCHAR(255),
		phone VARCHAR(50),
		email VARCHAR(100),
		address TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS medicines (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description TEXT,
		manufacturer VARCHAR(255),
		price DECIMAL(10, 2) NOT NULL,
		quantity INTEGER DEFAULT 0,
		expiry_date DATE,
		category VARCHAR(100),
		requires_prescription BOOLEAN DEFAULT false,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS purchases (
		id SERIAL PRIMARY KEY,
		medicine_id INTEGER REFERENCES medicines(id) ON DELETE CASCADE,
		supplier_id INTEGER REFERENCES suppliers(id) ON DELETE CASCADE,
		quantity INTEGER NOT NULL,
		unit_price DECIMAL(10, 2) NOT NULL,
		total_price DECIMAL(10, 2) NOT NULL,
		purchase_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS sales (
		id SERIAL PRIMARY KEY,
		medicine_id INTEGER REFERENCES medicines(id) ON DELETE CASCADE,
		user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
		quantity INTEGER NOT NULL,
		unit_price DECIMAL(10, 2) NOT NULL,
		total_price DECIMAL(10, 2) NOT NULL,
		sale_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE INDEX IF NOT EXISTS idx_medicines_name ON medicines(name);
	CREATE INDEX IF NOT EXISTS idx_medicines_category ON medicines(category);
	CREATE INDEX IF NOT EXISTS idx_suppliers_name ON suppliers(name);
	CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);
	CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
	`

	_, err := DB.Exec(schema)
	return err
}
