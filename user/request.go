package user

type UserRequest struct {
	Age      uint   `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}
