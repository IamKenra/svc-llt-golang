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
  - [ ] Implement elderly care CRUD operations
  - [ ] Add elderly care handlers
  - [ ] Create elderly care routes
- [ ] Database migrations
  - [ ] Create migration files for all entities
  - [ ] Set up auto-migration on startup
- [ ] Authentication & Authorization
  - [ ] Implement user registration
  - [ ] Add role-based access control
  - [ ] JWT refresh token mechanism

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

- Project follows `svc-partnership-go` architecture standards
- All struct entities centralized in `/entity/`
- All DTOs/payloads centralized in `/valueobject/`
- Clean separation between domain layers
- Ready for horizontal scaling with multiple domains

## 🔗 References

- Base architecture pattern: `svc-partnership-go`
- Framework: Fiber v2
- ORM: GORM v2
- Database: MySQL
- Authentication: JWT