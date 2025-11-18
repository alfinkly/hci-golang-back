#!/bin/bash

# API Testing Script for Pharmacy Backend
# This script tests the main endpoints of the API

API_URL="${API_URL:-http://localhost:8080}"

echo "=== Pharmacy Backend API Testing ==="
echo "API URL: $API_URL"
echo ""

# Test health endpoint
echo "1. Testing Health Endpoint..."
curl -s "$API_URL/health" | jq .
echo ""

# Register a new user
echo "2. Registering a new user..."
REGISTER_RESPONSE=$(curl -s -X POST "$API_URL/api/auth/register" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "email": "admin@pharmacy.com",
    "password": "admin123",
    "role": "admin"
  }')
echo "$REGISTER_RESPONSE" | jq .
echo ""

# Login
echo "3. Logging in..."
LOGIN_RESPONSE=$(curl -s -X POST "$API_URL/api/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "admin123"
  }')
echo "$LOGIN_RESPONSE" | jq .

# Extract token
TOKEN=$(echo "$LOGIN_RESPONSE" | jq -r '.token')
echo "Token: ${TOKEN:0:50}..."
echo ""

# Get profile
echo "4. Getting user profile..."
curl -s "$API_URL/api/profile" \
  -H "Authorization: Bearer $TOKEN" | jq .
echo ""

# Create a medicine
echo "5. Creating a medicine..."
MEDICINE_RESPONSE=$(curl -s -X POST "$API_URL/api/medicines" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "name": "Aspirin",
    "description": "Pain reliever and anti-inflammatory",
    "manufacturer": "Bayer",
    "price": 150.50,
    "quantity": 100,
    "expiry_date": "2025-12-31T00:00:00Z",
    "category": "Pain Relievers",
    "requires_prescription": false
  }')
echo "$MEDICINE_RESPONSE" | jq .
MEDICINE_ID=$(echo "$MEDICINE_RESPONSE" | jq -r '.id')
echo ""

# Get all medicines
echo "6. Getting all medicines..."
curl -s "$API_URL/api/medicines" \
  -H "Authorization: Bearer $TOKEN" | jq .
echo ""

# Create a supplier
echo "7. Creating a supplier..."
SUPPLIER_RESPONSE=$(curl -s -X POST "$API_URL/api/suppliers" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "name": "Pharma Supply Co.",
    "contact_person": "John Doe",
    "phone": "+1-555-0100",
    "email": "contact@pharmasupply.com",
    "address": "123 Medical Street, NY"
  }')
echo "$SUPPLIER_RESPONSE" | jq .
SUPPLIER_ID=$(echo "$SUPPLIER_RESPONSE" | jq -r '.id')
echo ""

# Create a purchase
if [ "$MEDICINE_ID" != "null" ] && [ "$SUPPLIER_ID" != "null" ]; then
  echo "8. Creating a purchase..."
  curl -s -X POST "$API_URL/api/purchases" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $TOKEN" \
    -d "{
      \"medicine_id\": $MEDICINE_ID,
      \"supplier_id\": $SUPPLIER_ID,
      \"quantity\": 50,
      \"unit_price\": 120.00
    }" | jq .
  echo ""
fi

# Create a sale
if [ "$MEDICINE_ID" != "null" ]; then
  echo "9. Creating a sale..."
  curl -s -X POST "$API_URL/api/sales" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $TOKEN" \
    -d "{
      \"medicine_id\": $MEDICINE_ID,
      \"quantity\": 2
    }" | jq .
  echo ""
fi

# Get all sales
echo "10. Getting all sales..."
curl -s "$API_URL/api/sales" \
  -H "Authorization: Bearer $TOKEN" | jq .
echo ""

echo "=== Testing Complete ==="
