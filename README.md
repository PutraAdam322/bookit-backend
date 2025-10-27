
---

## ⚙️ **Backend — `bookit-backend`**

```markdown
# Bookit (Backend)

A lightweight backend API for the Bookit prototype.  
Developed with **Go**, using **Gin** for routing and **GORM** for ORM database management.

---

## ✨ Features

1. **User Authentication**
   - User registration and login endpoints.
   - Token-based authentication using JWT.

2. **Facility Management**
   - CRUD operations for facilities (e.g., rooms, sports areas, labs).
   - Handles associated metadata such as name, capacity, and image.

3. **Booking System**
   - Manage booking slots for each facility.
   - Prevents overlapping bookings.
   - Calculates and returns total booking price.

4. **Booking API**
   - Receives `bookingSlotId` and `totalPrice` payloads from frontend.
   - Validates booking availability and saves data to the database.

5. **Database Integration**
   - Uses **GORM** for easy database migrations and relationships.
   - Supports MySQL / PostgreSQL / SQLite (depending on configuration).

6. **Encryption & Security**
   - Implements data encryption where necessary.
   - CORS enabled for frontend integration.
   - Follows RESTful API structure with JSON responses.

7. **Error Handling**
   - Centralized middleware for consistent error messages.
   - Validates input payloads using Gin’s binding and validator features.

## 🧩 Tech Stack

- **Go** (Gin Framework)
- **GORM** (ORM)
- **MySQL** (default database)
- **JWT Authentication**
