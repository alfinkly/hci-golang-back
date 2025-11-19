# API Documentation

## Base URL

```
http://localhost:8080/api
```

## Authentication

Most endpoints require JWT authentication. After logging in, include the JWT token in the Authorization header:

```
Authorization: Bearer <your-jwt-token>
```

## Endpoints

### Health Check

#### GET /health

Check if the API is running.

**No authentication required**

**Response:**
```json
{
  "status": "ok",
  "message": "Pharmacy API is running"
}
```

---

## Authentication Endpoints

### Register User

#### POST /api/auth/register

Create a new user account.

**No authentication required**

**Request Body:**
```json
{
  "username": "string (required)",
  "email": "string (required)",
  "password": "string (required)",
  "role": "string (optional, default: 'user')"
}
```

**Response (201 Created):**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "username": "admin",
    "email": "admin@pharmacy.com",
    "role": "admin",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

### Login

#### POST /api/auth/login

Authenticate a user and receive a JWT token.

**No authentication required**

**Request Body:**
```json
{
  "username": "string (required)",
  "password": "string (required)"
}
```

**Response (200 OK):**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "username": "admin",
    "email": "admin@pharmacy.com",
    "role": "admin",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

### Get Profile

#### GET /api/profile

Get the current user's profile.

**Authentication required**

**Response (200 OK):**
```json
{
  "id": 1,
  "username": "admin",
  "email": "admin@pharmacy.com",
  "role": "admin",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

---

## Medicine Endpoints

### Get All Medicines

#### GET /api/medicines

Retrieve all medicines.

**Authentication required**

**Response (200 OK):**
```json
[
  {
    "id": 1,
    "name": "Aspirin",
    "description": "Pain reliever",
    "manufacturer": "Bayer",
    "price": 150.50,
    "quantity": 100,
    "expiry_date": "2025-12-31T00:00:00Z",
    "category": "Pain Relievers",
    "requires_prescription": false,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
]
```

### Get Medicine by ID

#### GET /api/medicines/:id

Retrieve a specific medicine by ID.

**Authentication required**

**Response (200 OK):**
```json
{
  "id": 1,
  "name": "Aspirin",
  "description": "Pain reliever",
  "manufacturer": "Bayer",
  "price": 150.50,
  "quantity": 100,
  "expiry_date": "2025-12-31T00:00:00Z",
  "category": "Pain Relievers",
  "requires_prescription": false,
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

### Create Medicine

#### POST /api/medicines

Create a new medicine.

**Authentication required**

**Request Body:**
```json
{
  "name": "string (required)",
  "description": "string (optional)",
  "manufacturer": "string (optional)",
  "price": 150.50, // number (required)
  "quantity": 100, // integer (optional, default: 0)
  "expiry_date": "2025-12-31T00:00:00Z", // ISO 8601 date (optional)
  "category": "string (optional)",
  "requires_prescription": false // boolean (optional, default: false)
}
```

**Response (201 Created):**
```json
{
  "id": 1,
  "name": "Aspirin",
  "description": "Pain reliever",
  "manufacturer": "Bayer",
  "price": 150.50,
  "quantity": 100,
  "expiry_date": "2025-12-31T00:00:00Z",
  "category": "Pain Relievers",
  "requires_prescription": false,
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

### Update Medicine

#### PUT /api/medicines/:id

Update an existing medicine. All fields are optional.

**Authentication required**

**Request Body:**
```json
{
  "name": "string (optional)",
  "description": "string (optional)",
  "manufacturer": "string (optional)",
  "price": 150.50, // number (optional)
  "quantity": 100, // integer (optional)
  "expiry_date": "2025-12-31T00:00:00Z", // ISO 8601 date (optional)
  "category": "string (optional)",
  "requires_prescription": false // boolean (optional)
}
```

**Response (200 OK):**
```json
{
  "id": 1,
  "name": "Aspirin Updated",
  "description": "Pain reliever",
  "manufacturer": "Bayer",
  "price": 160.00,
  "quantity": 100,
  "expiry_date": "2025-12-31T00:00:00Z",
  "category": "Pain Relievers",
  "requires_prescription": false,
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T12:00:00Z"
}
```

### Delete Medicine

#### DELETE /api/medicines/:id

Delete a medicine.

**Authentication required**

**Response (200 OK):**
```json
{
  "message": "Medicine deleted successfully"
}
```

---

## Supplier Endpoints

### Get All Suppliers

#### GET /api/suppliers

Retrieve all suppliers.

**Authentication required**

**Response (200 OK):**
```json
[
  {
    "id": 1,
    "name": "Pharma Supply Co.",
    "contact_person": "John Doe",
    "phone": "+1-555-0100",
    "email": "contact@pharmasupply.com",
    "address": "123 Medical Street, NY",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
]
```

### Get Supplier by ID

#### GET /api/suppliers/:id

Retrieve a specific supplier by ID.

**Authentication required**

**Response (200 OK):**
```json
{
  "id": 1,
  "name": "Pharma Supply Co.",
  "contact_person": "John Doe",
  "phone": "+1-555-0100",
  "email": "contact@pharmasupply.com",
  "address": "123 Medical Street, NY",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

### Create Supplier

#### POST /api/suppliers

Create a new supplier.

**Authentication required**

**Request Body:**
```json
{
  "name": "string (required)",
  "contact_person": "string (optional)",
  "phone": "string (optional)",
  "email": "string (optional)",
  "address": "string (optional)"
}
```

**Response (201 Created):**
```json
{
  "id": 1,
  "name": "Pharma Supply Co.",
  "contact_person": "John Doe",
  "phone": "+1-555-0100",
  "email": "contact@pharmasupply.com",
  "address": "123 Medical Street, NY",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

### Update Supplier

#### PUT /api/suppliers/:id

Update an existing supplier. All fields are optional.

**Authentication required**

**Request Body:**
```json
{
  "name": "string (optional)",
  "contact_person": "string (optional)",
  "phone": "string (optional)",
  "email": "string (optional)",
  "address": "string (optional)"
}
```

**Response (200 OK):**
```json
{
  "id": 1,
  "name": "Pharma Supply Co. Updated",
  "contact_person": "Jane Doe",
  "phone": "+1-555-0100",
  "email": "contact@pharmasupply.com",
  "address": "123 Medical Street, NY",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T12:00:00Z"
}
```

### Delete Supplier

#### DELETE /api/suppliers/:id

Delete a supplier.

**Authentication required**

**Response (200 OK):**
```json
{
  "message": "Supplier deleted successfully"
}
```

---

## Purchase Endpoints

### Get All Purchases

#### GET /api/purchases

Retrieve all purchases.

**Authentication required**

**Response (200 OK):**
```json
[
  {
    "id": 1,
    "medicine_id": 1,
    "supplier_id": 1,
    "quantity": 50,
    "unit_price": 120.00,
    "total_price": 6000.00,
    "purchase_date": "2024-01-01T10:00:00Z",
    "created_at": "2024-01-01T10:00:00Z"
  }
]
```

### Get Purchase by ID

#### GET /api/purchases/:id

Retrieve a specific purchase by ID.

**Authentication required**

**Response (200 OK):**
```json
{
  "id": 1,
  "medicine_id": 1,
  "supplier_id": 1,
  "quantity": 50,
  "unit_price": 120.00,
  "total_price": 6000.00,
  "purchase_date": "2024-01-01T10:00:00Z",
  "created_at": "2024-01-01T10:00:00Z"
}
```

### Create Purchase

#### POST /api/purchases

Create a new purchase and automatically update medicine quantity.

**Authentication required**

**Request Body:**
```json
{
  "medicine_id": 1, // integer (required)
  "supplier_id": 1, // integer (required)
  "quantity": 50, // integer (required, must be > 0)
  "unit_price": 120.00 // number (required, must be > 0)
}
```

**Response (201 Created):**
```json
{
  "id": 1,
  "medicine_id": 1,
  "supplier_id": 1,
  "quantity": 50,
  "unit_price": 120.00,
  "total_price": 6000.00,
  "purchase_date": "2024-01-01T10:00:00Z",
  "created_at": "2024-01-01T10:00:00Z"
}
```

**Note:** The medicine quantity is automatically increased by the purchase quantity.

### Delete Purchase

#### DELETE /api/purchases/:id

Delete a purchase.

**Authentication required**

**Response (200 OK):**
```json
{
  "message": "Purchase deleted successfully"
}
```

---

## Sale Endpoints

### Get All Sales

#### GET /api/sales

Retrieve all sales.

**Authentication required**

**Response (200 OK):**
```json
[
  {
    "id": 1,
    "medicine_id": 1,
    "user_id": 1,
    "quantity": 2,
    "unit_price": 150.50,
    "total_price": 301.00,
    "sale_date": "2024-01-01T15:00:00Z",
    "created_at": "2024-01-01T15:00:00Z"
  }
]
```

### Get Sale by ID

#### GET /api/sales/:id

Retrieve a specific sale by ID.

**Authentication required**

**Response (200 OK):**
```json
{
  "id": 1,
  "medicine_id": 1,
  "user_id": 1,
  "quantity": 2,
  "unit_price": 150.50,
  "total_price": 301.00,
  "sale_date": "2024-01-01T15:00:00Z",
  "created_at": "2024-01-01T15:00:00Z"
}
```

### Create Sale

#### POST /api/sales

Create a new sale and automatically update medicine quantity.

**Authentication required**

**Request Body:**
```json
{
  "medicine_id": 1, // integer (required)
  "quantity": 2 // integer (required, must be > 0)
}
```

**Response (201 Created):**
```json
{
  "id": 1,
  "medicine_id": 1,
  "user_id": 1,
  "quantity": 2,
  "unit_price": 150.50,
  "total_price": 301.00,
  "sale_date": "2024-01-01T15:00:00Z",
  "created_at": "2024-01-01T15:00:00Z"
}
```

**Note:** 
- The medicine quantity is automatically decreased by the sale quantity.
- The unit price is fetched from the current medicine price.
- The user ID is taken from the JWT token.
- Returns 400 error if insufficient quantity is available.

### Delete Sale

#### DELETE /api/sales/:id

Delete a sale.

**Authentication required**

**Response (200 OK):**
```json
{
  "message": "Sale deleted successfully"
}
```

---

## Error Responses

All endpoints may return the following error responses:

### 400 Bad Request
```json
{
  "error": "Invalid request body"
}
```

### 401 Unauthorized
```json
{
  "error": "Missing authorization header"
}
```

### 403 Forbidden
```json
{
  "error": "Insufficient permissions"
}
```

### 404 Not Found
```json
{
  "error": "Resource not found"
}
```

### 500 Internal Server Error
```json
{
  "error": "Internal server error message"
}
```

## Rate Limiting

Currently, there is no rate limiting implemented. Consider implementing rate limiting for production use.

## CORS

CORS is enabled for all origins (`*`). Update the middleware configuration for production to restrict origins.
