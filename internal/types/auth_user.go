package types

type RegisterRequest struct {
	Name     int    `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Variables are passed by address so that their actual value is mutated.
// We pass by value when we dont want to mutate actual value.