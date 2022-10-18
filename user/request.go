package user

type UserRequest struct {
	Age      uint   `json:"age" binding:"required,number,min=8"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Username string `json:"username" binding:"required"`
}
