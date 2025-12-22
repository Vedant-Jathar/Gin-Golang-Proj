# Copilot Instructions for Gin-Golang-Proj

## Project Overview
A **Gin-based REST API** with layered architecture (MVC pattern) using GORM for ORM and MySQL for persistence. The project emphasizes separation of concerns with distinct layers for database, models, services, and controllers.

## Architecture & Key Components

### Layered Structure
```
Controllers → Services → Models ↔ Database
```

**Controllers** (`internal/controllers/`): HTTP endpoint handlers that parse requests and delegate to services.
- Example: `user.go` defines `UserController` with handlers returning `gin.HandlerFunc`
- Pattern: Each endpoint wraps logic in a function returning `gin.HandlerFunc` for Gin routes

**Services** (`internal/services/`): Business logic layer that operates on models and interacts with database.
- Example: `UserService` holds `*gorm.DB` and methods like `CreateUser()`, `GetUserById()`
- Pattern: Receivers are pointer types that encapsulate dependency injection (`db`)

**Models** (`internal/models/`): GORM data structures with tags for validation and ORM mappings.
- Example: `User` struct includes `gorm.Model`, field tags for `json:`, `gorm:` (constraints), and `binding:` (validation)
- Critical: Duplicate `Id` field alongside `gorm.Model` - both define primary keys (anti-pattern but current approach)

**Database** (`internal/database/db.go`): Single connection initialization function.
- Hardcoded MySQL DSN: `root:Vedant@12345@tcp(localhost:3306)/user_db`
- Also supports Postgres (commented out) - migrations use `db.AutoMigrate()`

### Cross-Layer Communication
- Controllers receive services via constructor-style initializers (e.g., `NewUserController()`)
- Services receive `*gorm.DB` as dependency
- Example flow: `UserController.GetUsers()` → `UserService.GetUsers()` → `db.Find()`

## Critical Developer Workflows

### Running the Application
1. **Start MySQL**: Use Docker (`docker-compose up postgres` - note: currently only Postgres configured)
   - Current code uses MySQL on `localhost:3306`, user `root`, password `Vedant@12345`
   - Database name: `user_db` (hardcoded in `db.go`)

2. **Build & Run**: 
   ```bash
   go build -o main . && ./main
   ```
   - Application starts on port `:8000` (hardcoded in `main.go`)
   - Auto-migrates tables for `User` and `AuthUser` models

### Testing
- `test.go` file exists but is empty - no test infrastructure currently in place
- Manual API testing required (Postman/curl for endpoints like `GET /user`, `POST /user`)

### Database Concerns
- **Current state**: MySQL hardcoded; Postgres configuration exists but commented out
- Migrations are automatic via `AutoMigrate()` on startup
- **Note**: `docker-compose.yaml` defines Postgres service but codebase uses MySQL - reconciliation needed

## Project-Specific Patterns & Conventions

### Naming & Package Structure
- **Typo**: Package named `contollers` (missing 'n') - used throughout; maintain for consistency
- **Typo**: Middleware package `midlleware` (missing 'd') - do not correct without updating imports

### Initialization Pattern (Dependency Injection)
Services and controllers use "constructor" methods rather than struct literals:
```go
// Correct pattern in this project:
UserCtlr := contollers.UserController{}
newUserCtlr := UserCtlr.NewUserController(*newUserSrv)

// Services are initialized similarly:
newUserSrv := UserSrv.NewUserService(db)
```
Follow this pattern for any new services/controllers.

### Gin Handler Pattern
All endpoints return `gin.HandlerFunc` closures that capture the service instance:
```go
func (u *UserController) GetUsers() gin.HandlerFunc {
    return func(c *gin.Context) {
        // c is *gin.Context, access parameters/body via c.Param(), c.BindJSON()
        // Respond with c.JSON(http.StatusCode, gin.H{...})
    }
}
```

### Error Handling & Responses
Standard response format uses `gin.H` with `status` boolean and `error` field:
```go
c.JSON(http.StatusBadRequest, gin.H{
    "status": false,
    "error":  err.Error(),
})
```

### Middleware
- Currently minimal: `AuthMiddleware` in `internal/midlleware/authMiddleware.go` sets `userId` to hardcoded `1`
- Not yet integrated into routes (commented out in controller)
- Pattern: Use `router.Use()` or group-level middleware before route definition

## Integration Points & External Dependencies

### Core Dependencies
- **Gin**: Web framework (`github.com/gin-gonic/gin`)
- **GORM**: ORM with MySQL/Postgres drivers
  - MySQL driver: `gorm.io/driver/mysql`
  - Postgres driver: `gorm.io/driver/postgres` (commented out but available)

### Database Models in Use
- `models.User`: Basic user info (Email, Name, Age)
- `models.AuthUser`: Authentication model (defined in `auth_user.go`)
- Both auto-migrated on startup

### Ports & Configuration
- **API Server**: `:8000` (hardcoded)
- **MySQL**: `localhost:3306` with credentials `root:Vedant@12345`
- **Postgres** (alternative, commented): `localhost:5433`

## Common Modifications & Extension Points

### Adding a New Endpoint
1. Create handler method in controller (return `gin.HandlerFunc`)
2. Register in `InitRoutes()` method (e.g., `routes.GET("/path", u.Handler())`)
3. Implement logic via service method if needed
4. Follow response format: `gin.H{"status": bool, "error": errorOrData}`

### Adding New Models
1. Define struct in `internal/models/` with GORM tags
2. Add to `db.AutoMigrate()` call in `main.go`
3. Create service methods in `internal/services/`
4. Expose via new controller in `internal/controllers/`

### Database Switching
- Modify DSN string in `db.go`
- Uncomment Postgres driver and switch import
- Note: `docker-compose.yaml` needs alignment with chosen database

## Known Issues & Anti-Patterns
- **Duplicate primary key**: `User` model has both `gorm.Model` (includes `ID`) and explicit `Id` field
- **Hardcoded credentials**: Database DSN hardcoded in source; use environment variables
- **Missing auth implementation**: `AuthController` and `AuthService` are shells
- **Commented main.go**: Original main function is entirely commented out; currently non-functional
- **Unused types**: `Student` type in `internal/types/student.go` not referenced
- **Database mismatch**: Code uses MySQL, but Docker only configures Postgres

---

**Last Updated**: December 2025
