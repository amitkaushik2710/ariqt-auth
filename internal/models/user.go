package models

// User
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Auth Request
type AuthReq struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// SignUp Response
type SignUpRes struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
