#!/bin/bash

# Lesson Management System - API Test Script
# Make sure your server is running on localhost:8080

BASE_URL="http://localhost:8080"

echo "🚀 Testing Lesson Management API"
echo "================================="
echo ""

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# 1. Register Admin
echo -e "${BLUE}1. Registering Admin...${NC}"
ADMIN_RESPONSE=$(curl -s -X POST "$BASE_URL/api/auth/register/admin" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Admin User",
    "email": "admin@example.com",
    "password": "admin123"
  }')
echo "$ADMIN_RESPONSE" | jq '.'
echo ""

# 2. Register Teacher
echo -e "${BLUE}2. Registering Teacher...${NC}"
TEACHER_RESPONSE=$(curl -s -X POST "$BASE_URL/api/auth/register/teacher" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Teacher User",
    "email": "teacher@example.com",
    "password": "teacher123"
  }')
echo "$TEACHER_RESPONSE" | jq '.'
TEACHER_ID=$(echo "$TEACHER_RESPONSE" | jq -r '.id')
echo ""

# 3. Register Student
echo -e "${BLUE}3. Registering Student...${NC}"
STUDENT_RESPONSE=$(curl -s -X POST "$BASE_URL/api/auth/register/teacher" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Student User",
    "email": "student@example.com",
    "password": "student123"
  }')
echo "$STUDENT_RESPONSE" | jq '.'
STUDENT_ID=$(echo "$STUDENT_RESPONSE" | jq -r '.id')
echo ""

# 4. Login as Admin
echo -e "${BLUE}4. Logging in as Admin...${NC}"
ADMIN_LOGIN=$(curl -s -X POST "$BASE_URL/api/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@example.com",
    "password": "admin123",
    "role": "admin"
  }')
echo "$ADMIN_LOGIN" | jq '.'
ADMIN_TOKEN=$(echo "$ADMIN_LOGIN" | jq -r '.token')
echo ""

# 5. Login as Teacher
echo -e "${BLUE}5. Logging in as Teacher...${NC}"
TEACHER_LOGIN=$(curl -s -X POST "$BASE_URL/api/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "teacher@example.com",
    "password": "teacher123",
    "role": "teacher"
  }')
echo "$TEACHER_LOGIN" | jq '.'
TEACHER_TOKEN=$(echo "$TEACHER_LOGIN" | jq -r '.token')
echo ""

# 6. Login as Student
echo -e "${BLUE}6. Logging in as Student...${NC}"
STUDENT_LOGIN=$(curl -s -X POST "$BASE_URL/api/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "student@example.com",
    "password": "student123",
    "role": "student"
  }')
echo "$STUDENT_LOGIN" | jq '.'
STUDENT_TOKEN=$(echo "$STUDENT_LOGIN" | jq -r '.token')
echo ""

# 7. Admin creates a lesson
echo -e "${GREEN}7. Admin creates a lesson...${NC}"
LESSON_RESPONSE=$(curl -s -X POST "$BASE_URL/api/lessons" \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -H "Content-Type: application/json" \
  -d "{
    \"title\": \"Math 101\",
    \"description\": \"Introduction to Mathematics\",
    \"teacher_id\": $TEACHER_ID
  }")
echo "$LESSON_RESPONSE" | jq '.'
LESSON_ID=$(echo "$LESSON_RESPONSE" | jq -r '.id')
echo ""

# 8. List all lessons
echo -e "${GREEN}8. Listing all lessons...${NC}"
curl -s -X GET "$BASE_URL/api/lessons" | jq '.'
echo ""

# 9. Teacher views their lessons
echo -e "${GREEN}9. Teacher views their lessons...${NC}"
curl -s -X GET "$BASE_URL/api/teacher/lessons" \
  -H "Authorization: Bearer $TEACHER_TOKEN" | jq '.'
echo ""

# 10. Student views their lessons
echo -e "${GREEN}10. Student views their lessons...${NC}"
curl -s -X GET "$BASE_URL/api/student/lessons" \
  -H "Authorization: Bearer $STUDENT_TOKEN" | jq '.'
echo ""

echo -e "${GREEN}✅ Test completed!${NC}"

