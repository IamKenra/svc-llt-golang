# SVC-LLT-Golang Development Tasks

## Project Overview
Service untuk Layanan Lansia Terpadu (LLT) menggunakan Go dengan clean architecture pattern.

## ✅ Completed Tasks

### 1. Architecture Restructuring
- [x] Analyzed `svc-partnership-go` structure pattern
- [x] Implemented clean architecture following svc-partnership-go standards
- [x] Created centralized `/entity/` directory for all struct entities
- [x] Created centralized `/valueobject/` directory for all DTOs/payloads
- [x] Restructured domain layers to match pattern:
  ```
  /domain
  ├── masterdata/
  │   ├── repository.go      ← Interface repository
  │   ├── usecase.go         ← Interface usecase
  │   ├── repository/        ← Implementasi repository
  │   ├── usecase/           ← Implementasi usecase
  │   └── delivery/http/     ← Handler & routes
  └── llt/                   ← Domain lain (siap untuk pengembangan)
  ```

### 2. Code Organization
- [x] Removed duplicate `/internal/` structure
- [x] Moved middleware to `/utils/middleware/`
- [x] Renamed `/pkg/` to `/utils/` to match pattern
- [x] Updated all imports and references throughout codebase
- [x] Cleaned up final structure

### 3. Database Migration
- [x] Changed database driver from PostgreSQL to MySQL
- [x] Updated DSN format for MySQL connection
- [x] Updated table names (removed PostgreSQL schema prefix)
- [x] Updated `.env.example` for MySQL configuration
- [x] Tested MySQL connection and build process

### 4. Environment Configuration
- [x] Verified `.env` implementation is ready
- [x] Created `.env.example` template with all required variables
- [x] Configured environment variables for:
  - Database connection (MySQL)
  - JWT authentication
  - Internal API key
  - Server port

### 5. Code Standards & Best Practices
- [x] Refactored usecase structure to match svc-partnership-go regulation
  - Single Usecase interface per domain (DDD compliant)
  - Proper file structure: `usecase.go`, `usecase/app.go`, `usecase/api-*.go`
  - Method naming: `StoreUser`, `UpdateUser`, `DeleteUser` with payload objects
- [x] Updated repository pattern to match svc-partnership-go
  - MySQL-specific naming: `mysql-*.go` files
  - Constructor: `NewMysqlMasterdataRepository()`
  - Receiver parameter: `db *mysqlMasterdataRepository` (not `r`)
- [x] Refactored handler parameter naming
  - Receiver: `handler *HandlerStruct` (not `h`)
  - Context: `ctx *fiber.Ctx` (not `c`)
- [x] Implemented proper clean architecture for HealthHandler
  - Removed direct database dependency from handler
  - Added health check to usecase and repository layers
  - Handler now uses usecase interface instead of direct GORM access
- [x] **Business Logic & Processor Separation** ⭐
  - **API Layer** (`api-*.go`): Contains business logic, validation, orchestration
  - **Processor Layer** (`processor-*.go`): Database operations preparation for ORM
  - **Handler Layer**: HTTP request/response handling and response formatting
  - Clear separation following svc-partnership-go pattern with ORM adaptations
- [x] **Random Bigint ID Implementation** 🔒
  - Changed all entity IDs from `uint` to `int64` with `autoIncrement:false`
  - Cryptographically secure random ID generation (19-digit minimum)
  - Repository layer generates IDs on create operations
  - Prevents enumeration attacks and distributed system ID conflicts

### 6. Project Maintenance
- [x] Created comprehensive `.gitignore` file
- [x] Removed `/bin` directory (40MB of build artifacts)
- [x] Added `/bin/` to `.gitignore` to prevent future binary commits

## 📁 Final Project Structure

```
/svc-llt-golang
├── entity/                 ← 🎯 Semua struct entity terpusat
│   ├── user.go
│   ├── identity.go
│   └── elderly_care.go
├── valueobject/           ← 🎯 Semua DTO/payload/response terpusat  
│   ├── user.go
│   └── elderly_care.go
├── domain/                ← 🎯 Domain modules
│   ├── masterdata/
│   │   ├── repository.go      ← Interface repository
│   │   ├── usecase.go         ← Interface usecase
│   │   ├── repository/        ← Implementasi repository
│   │   ├── usecase/           ← Implementasi usecase
│   │   └── delivery/http/     ← Handler & routes
│   └── llt/               ← Domain untuk elderly care (ready for dev)
├── cmd/                   ← Binary entry points
│   ├── api/main.go        ← Main API server
│   ├── server/main.go     ← Alternative server
│   └── seed/              ← Database seeding
├── utils/                 ← Shared utilities
│   ├── logger/
│   ├── response/
│   ├── middleware/        ← Auth, CORS, etc.
│   └── utils/             ← Hash, etc.
├── config/                ← Configuration
├── bin/                   ← Compiled binaries
├── .env.example          ← Environment template
└── _document/            ← Project documentation (preserved)
```


## 🔧 Technical Configuration

### Database
- **Type**: MySQL
- **Port**: 3306
- **DSN Format**: `user:pass@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local`
- **Tables**: `users`, `identities`, `elderly_care`

### Environment Variables
```env
# Database Configuration (MySQL)
DB_HOST=localhost
DB_USERNAME=root
DB_PASSWORD=password
DB_NAME=svc_llt_golang
DB_PORT=3306

# JWT Configuration
JWT_SECRET=your-jwt-secret-key-here

# API Configuration
INTERNAL_API_KEY=your-internal-api-key-here

# Server Configuration
PORT=3000
```

## 📋 Next Development Tasks (TODO)

### High Priority
- [ ] Complete LLT domain implementation
  - [x] Basic structure created (usecase, repository, entity)
  - [ ] Implement elderly care CRUD handlers
  - [ ] Create elderly care routes
  - [ ] Add business logic validation
- [ ] Database migrations & seeding
  - [ ] Create migration files for all entities
  - [ ] Set up auto-migration on startup
  - [ ] Create database seeder for test data
- [ ] Authentication & Authorization enhancements
  - [ ] Implement user registration
  - [ ] Add role-based access control
  - [ ] JWT refresh token mechanism
  - [ ] Password reset functionality

### Medium Priority
- [ ] API Documentation
  - [ ] Set up Swagger/OpenAPI documentation
  - [ ] Document all endpoints
  - [ ] Add request/response examples
- [ ] Testing
  - [ ] Unit tests for repositories
  - [ ] Unit tests for usecases
  - [ ] Integration tests for APIs
  - [ ] Add test database configuration
- [ ] Logging & Monitoring
  - [ ] Improve logging structure
  - [ ] Add request tracing
  - [ ] Performance monitoring

### Low Priority
- [ ] Docker containerization
  - [ ] Create Dockerfile
  - [ ] Docker compose for development
  - [ ] Environment-specific configs
- [ ] CI/CD Pipeline
  - [ ] GitHub Actions setup
  - [ ] Automated testing
  - [ ] Deployment scripts
- [ ] Performance Optimization
  - [ ] Database query optimization
  - [ ] Caching implementation
  - [ ] API response optimization

## 🚀 Quick Start

1. **Setup Environment**
   ```bash
   cp .env.example .env
   # Edit .env with your database credentials
   ```

2. **Install Dependencies**
   ```bash
   go mod tidy
   ```

3. **Setup Database**
   ```sql
   CREATE DATABASE svc_llt_golang;
   ```

4. **Build & Run**
   ```bash
   go build -o bin/api ./cmd/api
   ./bin/api
   ```

## 📝 Notes

- Project follows `svc-partnership-go` architecture standards with 100% compliance
- All struct entities centralized in `/entity/` with random bigint IDs
- All DTOs/payloads centralized in `/valueobject/`
- Clean separation between domain layers with proper dependency inversion
- Ready for horizontal scaling with multiple domains
- **Security-first approach**: Random IDs prevent enumeration attacks
- Build successful with no compilation errors

> **📋 For detailed implementation guidelines, see [notes.md](./notes.md)**

## 🔗 References

- Base architecture pattern: `svc-partnership-go`
- Framework: Fiber v2
- ORM: GORM v2
- Database: MySQL
- Authentication: JWT