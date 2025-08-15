# SVC-LLT-Golang Implementation Notes

## Architecture Regulations (Based on svc-partnership-go)

### Layer Separation Rules

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

### File Naming Conventions
- API files: `api-{object}.go` 
- Processor files: `processor-{object}.go`
- Repository files: `mysql-{object}.go`
- Handler files: `{object}_handler.go`

### Response Handling
```go
// ✅ CORRECT - Usecase returns raw data
func (uc) Register(req) (string, error) { ... }

// Handler formats response
uuid, err := usecase.Register(req)
return response.Success(ctx, UserRegisterResponse{
    Message: "Success", UUID: uuid,
})
```

## Database & ID Generation

### Random Bigint IDs
- **All entities use `int64` primary keys** (not auto-increment)
- **Cryptographically secure random generation**
- Range: 1000000000000000000 to 9223372036854775807
- Generated at repository layer during create operations

### Entity Structure
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

## Technical Stack
- **Framework**: Fiber v2
- **ORM**: GORM v2  
- **Database**: MySQL
- **Authentication**: JWT
- **Architecture**: Clean Architecture (DDD)

## Project Structure
```
/entity/           ← All struct entities
/valueobject/      ← All DTOs/payloads/responses  
/domain/           ← Domain modules (masterdata, llt, etc.)
/utils/            ← Shared utilities
/config/           ← Configuration
```

## Key Implementation Details
- Single Usecase interface per domain (DDD compliant)
- Centralized entities and value objects
- Proper dependency inversion between layers
- MySQL with random bigint IDs for security
- Clean separation of concerns across all layers