# Lesson Management System - API Documentation

## Architecture Overview

The system has **3 user roles** (Admin, Teacher, Student), each with different access levels:

- **Admin**: Full system access
- **Teacher**: Can manage their assigned lessons and students
- **Student**: Read-only access to enrolled lessons

## Database Tables

1. `admins` - Admin users
2. `teachers` - Teacher users  
3. `students` - Student users
4. `lessons` - Lesson/course records
5. `lesson_students` - Many-to-many join table for student enrollment

## Authentication

### Register Admin
```
POST /api/auth/register/admin
```

**Request Body:**
```json
{
  "name": "Admin User",
  "email": "admin@example.com",
  "password": "password123"
}
```

### Register Teacher
```
POST /api/auth/register/teacher
```

**Request Body:**
```json
{
  "name": "Teacher User",
  "email": "teacher@example.com",
  "password": "password123"
}
```

### Login Endpoint
```
POST /api/auth/login
```

**Request Body:**
```json
{
  "email": "admin@example.com",
  "password": "password123",
  "role": "admin"  // or "teacher" or "student"
}
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "role": "admin",
  "name": "John Doe"
}
```

**Use the token in subsequent requests:**
```
Authorization: Bearer <token>
```

## Role-Based Endpoints

### Public Endpoints (No Auth Required)
```
GET  /api/lessons               - List all lessons
GET  /api/lessons/{id}          - Get lesson by ID
```

### Admin-Only Endpoints
```
POST /api/lessons                    - Create lesson
PUT  /api/lessons/{id}                - Update lesson
DELETE /api/lessons/{id}             - Delete lesson
POST /api/lessons/{id}/assign-teacher - Assign teacher to lesson
POST /api/lessons/{id}/enroll-student - Enroll student in lesson
```

### Teacher-Only Endpoints
```
GET  /api/teacher/lessons                 - Get teacher's lessons
GET  /api/lessons/{id}/students           - Get students in lesson
POST /api/lessons/{id}/students           - Add student to lesson
DELETE /api/lessons/{id}/students/{studentId} - Remove student from lesson
```

### Student-Only Endpoints
```
GET  /api/student/lessons  - Get enrolled lessons
```

## How It Works

### 1. Admin Flow
1. Admin logs in: `POST /api/auth/login` with `role: "admin"`
2. Receives JWT token
3. Creates a lesson: `POST /api/lessons` (includes teacher_id in request)
4. Assigns teacher: `POST /api/lessons/{id}/assign-teacher`
5. Enrolls students: `POST /api/lessons/{id}/enroll-student`

### 2. Teacher Flow
1. Teacher logs in: `POST /api/auth/login` with `role: "teacher"`
2. Receives JWT token  
3. Views their lessons: `GET /api/teacher/lessons`
4. Views students: `GET /api/lessons/{id}/students`
5. Manages students: add/remove from their lessons

### 3. Student Flow
1. Student logs in: `POST /api/auth/login` with `role: "student"`
2. Receives JWT token
3. Views enrolled lessons: `GET /api/student/lessons`

## Where the Logic Lives

### Entities (Database Models)
- `entities/admin.go` - Admin table structure
- `entities/teacher.go` - Teacher table structure
- `entities/student.go` - Student table structure
- `entities/lesson.go` - Lesson table with TeacherID foreign key

### Authentication Module
- `internal/modules/auth/handler.go` - Login handler
- `internal/modules/auth/service.go` - JWT generation, password hashing
- `internal/modules/auth/repository.go` - Database queries for all user types
- `internal/modules/auth/router.go` - Registers login endpoint

### Middleware
- `pkg/middleware/auth.go` - JWT validation and role enforcement

### Lessons Module
- `internal/modules/lessons/handler.go` - All CRUD operations + role-specific endpoints
- `internal/modules/lessons/service.go` - Business logic with authorization checks
- `internal/modules/lessons/repo.go` - Database operations
- `internal/modules/lessons/router.go` - Routes with role-based middleware

## Security

- JWT tokens expire after 24 hours
- Passwords are hashed with bcrypt
- Role-based access enforced at route level
- Teachers can only access their own lessons
- Students can only see their enrolled lessons

## Example Usage

### Create Lesson (Admin)
```bash
curl -X POST http://localhost:8080/api/lessons \
  -H "Authorization: Bearer YOUR_ADMIN_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Math 101",
    "description": "Introduction to Mathematics",
    "teacher_id": 1
  }'
```

### Login (Teacher)
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "teacher@example.com",
    "password": "password123",
    "role": "teacher"
  }'
```

### View Teacher's Lessons
```bash
curl http://localhost:8080/api/teacher/lessons \
  -H "Authorization: Bearer YOUR_TEACHER_TOKEN"
```

