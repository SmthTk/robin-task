# Robin Task API

Robin Task API is a **task management system** built using **Go** and **Gin Framework**, featuring **role-based access control** and **task archiving**.

---

## **🚀 Getting Started**

### **📌 Prerequisites**
Before running the project, make sure you have the following installed:
- **[Docker](https://www.docker.com/)**
- **[Go](https://go.dev/)** (if running locally)
- **[MySQL](https://www.mysql.com/)** (if not using Docker)

---

## **🛠 Running the Project with Docker**
To start the API using **Docker Compose**, follow these steps:

1. **Navigate to the project directory:**
   ```bash
   cd robin-task
   ```

2. **Start the services using Docker:**
   ```bash
   docker-compose up -d --build
   ```
   This will **build and run the API and MySQL database in Docker containers**.

3. **Verify the API is running:**
   - Open in your browser or Postman:
     ```
     http://localhost:5000
     ```

---

## **📂 Database & Example Data**
- **Database Initialization:**
   - Example data is preloaded from:
     ```
     init-database/init.sql
     ```
- **Default Login Credentials:**
- 
  | Role | Username | Password |
- 
  | **Admin** | `admin` | `password` |
- 
  | **User**  | `user1` | `password` |
-  
  | **User**  | `user2` | `password` |

- **Role Differences:**
   - **Admin**: Can **view archived tasks**.
   - **User**: Cannot see archived tasks.

---

## **📌 API Base URL**
```
http://localhost:5000
```

---

## **📬 Postman Collection**
To test API endpoints easily, **import the Postman collection**:
- File:
  ```
  /robin-task/postman/Robin Task API Collection.postman_collection.json
  ```
- **Import into Postman** → Start testing the API!

---

## **🚀 Features**
✅ **Task Management (CRUD)**  
✅ **User Authentication & JWT Tokens**  
✅ **Role-Based Access (Admin & User)**  
✅ **Task Archiving (Only Admins can see archived tasks)**  
✅ **Database Initialization with Example Data**  
✅ **Postman Collection for Easy Testing**

---
