package models

type UserDetails struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required" validate:"email"`
	Phone           string `json:"phone" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmpassword" binding:"required"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required" validate:"email"`
	Password string `json:"password" binding:"required"`
}

// user details shown after logging in
type UserDetailsResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type UserDetailsAtAdmin struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	BlockStatus bool   `json:"block_status"`
}

// user details along with embedded token which can be used by the user to access protected routes
type TokenUsers struct {
	Users UserDetailsResponse
	Token string
}

type UserSignInResponse struct {
	Id       uint   `json:"id"`
	UserID   uint   `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email" validate:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
