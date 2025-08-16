# SVC-LLT-Golang Technical Guidelines

## Table of Contents
- [Architecture Patterns](#architecture-patterns-based-on-svc-partnership-go)
- [Database Design](#database-design)
- [Security Implementation](#security-implementation)
- [Database Safety](#database-safety-protocols)
- [Feature Development](#feature-development-process)
- [Tech Stack](#tech-stack-reference)

## Architecture Patterns (Based on svc-partnership-go)

### Clean Architecture Layer Rules

#### **Handler Layer** (`delivery/http/`)
- HTTP request/response handling ONLY
- Parse requests, format responses, logging
- **NO business logic or direct repository calls**

#### **Usecase API Layer** (`usecase/api-*.go`)
- Business logic, validation, orchestration
- JWT generation, password hashing, UUID generation
- **Returns raw data (primitives, entities) - NO response formatting**

#### **Processor Layer** (`usecase/processor-*.go`) 
- Database operation preparation for ORM
- Parameter mapping, batch operations
- **NO business logic or validation**

#### **Repository Layer** (`repository/mysql-*.go`)
- GORM database operations only
- **NO business logic**

### Naming Conventions
- API files: `api-{object}.go` 
- Processor files: `processor-{object}.go`
- Repository files: `mysql-{object}.go`
- Handler files: `{object}_handler.go`

### Response Pattern
```go
// ✅ Usecase returns raw data
func (uc) Register(req) (string, error) { ... }

// Handler formats response
uuid, err := usecase.Register(req)
return response.Success(ctx, UserRegisterResponse{
    Message: "Success", UUID: uuid,
})
```

## Database Design

### Secure ID Generation
- **All entities use `int64` primary keys** (not auto-increment)
- **Cryptographically secure random generation**
- Range: 1000000000000000000 to 9223372036854775807
- Generated at repository layer during create operations

### Entity Pattern
```go
type Entity struct {
    ID int64 `gorm:"primaryKey;autoIncrement:false"`
    // other fields...
}
```

### Create Operations
- `utils.GenerateRandomID()` generates secure random int64
- Repository handles ID generation before database insert
- Prevents enumeration attacks and distributed system conflicts

## Security Implementation

### X-Member Header Pattern
```go
// utils/header/header.go
func ExtractXMember(ctx *fiber.Ctx) (string, error)
func ValidateAndExtractXMember(ctx *fiber.Ctx) (string, error)

// Handler usage
xMember, err := header.ValidateAndExtractXMember(ctx)
if err != nil {
    return err
}
req.User = xMember
```

### Endpoint Security
- Use query parameters instead of path parameters
- Before: `/lansia:uuid` → After: `/lansia/detail?uuid=xxx`
- Prevents UUID exposure in logs, browser history, referrer headers
- All payload structs include `User` field for audit trail

## Database Safety Protocols

### Critical Rules
**⚠️ NEVER DROP EXISTING TABLES WITHOUT EXPLICIT USER PERMISSION ⚠️**

### Safety Practices
1. Always ask before destructive database operations
2. Never use `DropTable()` unless explicitly requested
3. Use `HasTable()` to check existence before migration
4. Prefer `AutoMigrate()` only (adds columns/tables, doesn't drop)
5. Assume production data exists - treat all databases as important

### Safe Migration Pattern
```go
// ✅ SAFE - Only add/modify, never drop
if !db.Migrator().HasTable(&entity.TableName{}) {
    log.Println("Creating new table...")
    err = db.AutoMigrate(&entity.TableName{})
} else {
    log.Println("Table exists, skipping creation")
}
```

### Dangerous Operations
```go
// ❌ REQUIRE EXPLICIT PERMISSION
db.Migrator().DropTable()     // Deletes entire table and data
db.Exec("DROP TABLE...")      // Direct SQL drop
db.Migrator().DropColumn()    // Deletes column and data
```

## Feature Development Process

### 8-Step Implementation Guide

#### 1. Entity Definition (`/entity/`)
```go
type Entity struct {
    ID           int64   `gorm:"primaryKey;autoIncrement:false"`
    UUID         string  `gorm:"type:varchar(255);uniqueIndex;not null"`
    Field        *string `gorm:"column:field_name;type:varchar(255)"`
    CreatedAt    time.Time
    UpdatedAt    time.Time
    UserInput    string  `gorm:"column:user_input;type:varchar(255);not null"`
    UserUpdate   *string `gorm:"column:user_update;type:varchar(255)"`
}

func (Entity) TableName() string {
    return "table_name"
}
```

#### 2. Value Objects (`/valueobject/`)
```go
type EntityPayloadInsert struct {
    Data []Entity `json:"data"`
    User string   `json:"user"`
}

type EntityPayloadUpdate struct {
    Data EntityUpdateData `json:"data"`
    User string           `json:"user"`
}
```

#### 3. Repository Interface (`/domain/{module}/repository.go`)
```go
type Repository interface {
    GetAllEntity(param map[string]interface{}) ([]valueobject.Entity, error)
    GetOneEntity(param map[string]interface{}) (valueobject.Entity, error)
    CreateEntity(params...) error
    UpdateEntity(param, data map[string]interface{}) error
    DeleteEntity(param map[string]interface{}) error
}
```

#### 4. Repository Implementation (`/domain/{module}/repository/mysql-{entity}.go`)
- Use `utils.GenerateRandomID()` for ID generation
- Handle nullable fields with pointers
- Use query parameter maps for flexible filtering

#### 5. Processor Layer (`/domain/{module}/usecase/processor-{entity}.go`)
- Parameter mapping and batch operations
- NO business logic or validation

#### 6. API Layer (`/domain/{module}/usecase/api-{entity}.go`)
- Business logic, validation, orchestration
- Return raw data only

#### 7. Handler (`/domain/{module}/delivery/http/{entity}_handler.go`)
- Extract X-Member header: `header.ValidateAndExtractXMember(ctx)`
- Format JSON responses
- Use query parameters for security

#### 8. Routes (`/domain/{module}/delivery/http/routes.go`)
- Register endpoints with security query parameters

### Implementation Patterns

#### Nullable Fields
```go
// Use pointers for optional database fields
Field *string `gorm:"column:field_name;type:varchar(255)"`

// Assignment
if value != "" {
    entity.Field = &value
}
```

#### User Tracking
```go
// All payloads include User field
type Payload struct {
    Data interface{} `json:"data"`
    User string      `json:"user"`
}

// Extract from X-Member header
xMember, err := header.ValidateAndExtractXMember(ctx)
req.User = xMember
```

#### Security Patterns
```go
// Query parameters instead of path parameters
// Before: /entity/:uuid
// After:  /entity/detail?uuid=xxx

// X-Member header validation
xMember, err := header.ValidateAndExtractXMember(ctx)
if err != nil {
    return err
}
```

## Tech Stack Reference
- **Framework**: Fiber v2
- **ORM**: GORM v2  
- **Database**: MySQL
- **Authentication**: JWT
- **Architecture**: Clean Architecture (DDD)