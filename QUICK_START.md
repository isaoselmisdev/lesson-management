# Quick Start Guide

## 1. Setup Environment

Create a `.env` file in the root directory:

```env
# Database
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=lesson_management
DB_PORT=5432

# Server
PORT=8080

# JWT Secret
JWT_SECRET=your-super-secret-jwt-key-change-in-production
```

## 2. Run the Server

```bash
go run cmd/server/main.go
```

You should see: `✅ Server running on port: 8080`

## 3. Quick Test Script

Run the provided test script:
```bash
./test_requests.sh
```

Or test manually:

### Step 1: Create an Admin
```bash
curl -X POST http://localhost:8080/api/auth/register/admin \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Admin User",
    "email": "admin@example.com",
    "password": "admin123"
  }'
```

### Step 2: Login as Admin
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@example.com",
    "password": "admin123",
    "role": "admin"
  }'
```

**Save the token** from the response!

### Step 3: Create a Teacher
```bash
curl -X POST http://localhost:8080/api/auth/register/teacher \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Teacher User",
    "email": "teacher@example.com",
    "password": "teacher123"
  }'
```

**Save the teacher ID** from the response!

### Step 4: Create a Lesson (Admin only)
```bash
curl -X POST http://localhost:8080/api/lessons \
  -H "Authorization: Bearer YOUR_ADMIN_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Math 101",
    "description": "Introduction to Mathematics",
    "teacher_id": TEACHER_ID_HERE
  }'
```

### Step 5: Test Teacher Endpoint
```bash
# Login as teacher first
TOKEN=$(curl -s -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"teacher@example.com","password":"teacher123","role":"teacher"}' \
  | jq -r '.token')

# View teacher's lessons
curl -X GET http://localhost:8080/api/teacher/lessons \
  -H "Authorization: Bearer $TOKEN"
```

## API Endpoints Summary

### Authentication
- `POST /api/auth/register/admin` - Create admin user
- `POST /api/auth/register/teacher` - Create teacher user  
- `POST /api/auth/login` - Login (returns JWT token)

### Admin Endpoints
- `POST /api/lessons` - Create lesson
- `PUT /api/lessons/{id}` - Update lesson
- `DELETE /api/lessons/{id}` - Delete lesson
- `POST /api/lessons/{id}/assign-teacher` - Assign teacher
- `POST /api/lessons/{id}/enroll-student` - Enroll student

### Teacher Endpoints
- `GET /api/teacher/lessons` - Get my lessons
- `GET /api/lessons/{id}/students` - Get students in my lesson
- `POST /api/lessons/{id}/students` - Add student to my lesson
- `DELETE /api/lessons/{id}/students/{studentId}` - Remove student

### Student Endpoints
- `GET /api/student/lessons` - Get my enrolled lessons

## Testing with Postman

1. Import the endpoints from `API_DOCUMENTATION.md`
2. Create a Collection "Lesson Management"
3. Add environment variables:
   - `base_url`: `http://localhost:8080`
   - `admin_token`: (set after login)
   - `teacher_token`: (set after login)
   - `student_token`: (set after login)
4. Test the sequence: Register → Login → Use endpoints

## Troubleshooting

### "Connection refused"
- Make sure server is running on port 8080
- Check database is running

### "Invalid credentials"
- Make sure you're using the correct role parameter
- Email must match exactly

### "Authorization required"
- Include the `Authorization: Bearer <token>` header
- Token expires after 24 hours

