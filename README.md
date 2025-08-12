# 🚀 Go Fiber API with Hot Reload and Database Configuration

## 📌 Overview
This project is a **Go Fiber** application with:
- **Hot Reload** using `air`
- **Database Configuration** via environment variables in `.sh` script
- **Modular Project Structure** following best practices

## 📂 Project Structure
```
project-root/
│── cmd/                  # Entry point of the application
│   ├── main.go
│
│── config/               # Application configuration
│   ├── config.go
│
│── internal/             # Business logic
│   ├── handlers/         # Controller / API handlers
│   ├── services/         # Business logic services
│   ├── repositories/     # Database queries
│   ├── models/           # Data structures / database models
│
│── pkg/                  # Utility functions (middleware, logging, response handling)
│
│── routes/               # API route definitions
│
│── database/             # Database connection
│
│── .env                  # Environment variables
│── run_example.sh        # Script to start the project
│── air.toml              # Hot reload configuration
│── go.mod                # Dependencies
│── go.sum                # Dependency checksums
│── README.md             # Documentation
```

---

## ⚙️ Installation & Setup
### 1️⃣ Clone Repository
```sh
git clone https://github.com/your-repo/project.git
cd project
```

### 2️⃣ Install Dependencies
```sh
go mod tidy
```

### 3️⃣ Install `air` for Hot Reload (if not installed)
```sh
go install github.com/cosmtrek/air@latest
export PATH=$(go env GOPATH)/bin:$PATH
```

### 4️⃣ Configure Database
Modify **`run_example.sh`** with your database credentials:
```sh
export DB_DRIVERNAME="postgres"
export DB_USERNAME="youruser"
export DB_PASSWORD="yourpassword"
export DB_HOST="localhost"
export DB_PORT="5432"
export DB_NAME="yourdatabase"
```

### 5️⃣ Start the Project with Hot Reload
```sh
./run_example.sh
```
This will automatically restart the server on code changes.

---

## 🔧 API Endpoints
### 📝 Example: LLM API
| Method | Endpoint  | Description |
|--------|----------|-------------|
| GET    | `/api/llm` | Fetch LLM response |

Example request:
```sh
curl -X GET http://localhost:3000/api/llm
```

---

## 🚀 Deployment
To deploy the application, build the binary:
```sh
go build -o main ./cmd/main.go
./main
```

For Docker deployment, create a `Dockerfile`:
```dockerfile
FROM golang:1.21
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o main ./cmd/main.go
CMD ["./main"]
EXPOSE 80
```

Build and run the Docker container:
```sh
docker build -t go-fiber-app .
docker run -p 80:80 go-fiber-app
```

---

## 📜 License
This project is licensed under the **MIT License**.

---

## 📞 Contact
For support, contact **[Your Name]** at **your.email@example.com**.

Happy coding! 🎯

