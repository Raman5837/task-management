# 📝 Task Management Service

## **📌 Overview**

The **Task Management Microservice** provides CRUD operations for managing tasks, including:

- **Create, Read, Update, Delete, and List Tasks**
- **Logging with Zerolog for better observability**
- **Flexible Deployment with Binary Executable or Docker**
- **Database using SQLite & Bun ORM (Extendable to other databases)**
- **Supports Pagination & Filtering (`GET /tasks?status=completed`)**


---

## **📌 Problem Breakdown & Design Decisions**
### **🎯 Problem Statement**
The Task Management Service is designed to:
- Provide a **RESTful API** for task management (**Create, Read, Update, Delete, List**).
- Support **pagination and filtering** (`GET /tasks?status=completed`).
- Be **lightweight, modular, and scalable** for future enhancements.

### **🛠 Design Decisions**
1. **Layered Architecture:**
   - **Models:** Define the database schema.
   - **Services:** Implement business logic.
   - **Repositories:** Handle database operations.
   - **Handlers:** Manage API request-response cycles.
   - **Middleware:** Process pagination, validation, and request modifications.

2. **Choice of SQLite & Bun ORM**
   - Chosen for its **lightweight and efficient** setup in a microservice environment.
   - The database manager (`DBManager`) is designed to be **extensible**, allowing easy migration to **PostgreSQL, MySQL, or other databases** as needed.

3. **Scalability & Deployment**
   - Can be deployed as a **standalone binary** (`go build`) or as a **Docker container**.
   - Supports **horizontal scaling** by running multiple instances behind a load balancer if required.

---

### **Future Enhancements (Requires Further Development)**
1. **Inter-Service Communication**
   - For interaction with other microservices (e.g., a User Service):
     - **REST API Calls** can be used for **synchronous communication**, such as retrieving metadata (`GET /users/:id`).
     - **Event-Driven Communication** (Kafka, RabbitMQ) can be implemented for **asynchronous workflows**, such as **state changes, notifications, or background processing**.


## **📌 Installation & Running the Service**

### **Prerequisites**

- **Go 1.23+**
- **SQLite**
- **Docker (Optional for Containerized Deployment)**

### **Steps to Run**

```sh
Step 1: Clone the Repository
git clone git@github.com:Raman5837/task-management.git
cd task-management

Step 2: Install Dependencies
go mod tidy
make install-requirements (Optional) (To install extra dependencies such as Air, Linters)

Option 1: Run locally
    go run main.go

Option 2: Run with live reload
    make watch-local         

Option 3: Run using Docker
    docker-compose up --build -d
```

## **📌 API Documentation**

```
Task Endpoints
Method	            Endpoint	                                        Description
POST	            /api/v1/tasks/	                                Create a new task
GET	            /api/v1/tasks/:id	                                Get a task by ID
PUT	            /api/v1/tasks/:id	                                Update a task
DELETE	            /api/v1/tasks/:id	                                Delete a task
GET	            /api/v1/tasks?status=completed&limit=25&offset=0	List tasks with pagination & filtering
```

## **📌 Sample API Requests**
You can also refer to the `sample.http` file in the repository to see various API requests and responses in action.
