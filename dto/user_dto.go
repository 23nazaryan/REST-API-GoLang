package dto

type RegisterDTO struct {
	Name  string `json:"name" form:"name" binding:"required"`
	Email string `json:"email" form:"email" binding:"required,email"`
	Role  string `json:"role" form:"role" binding:"required"`
}

type PwdDTO struct {
	Hash     string `json:"hash" form:"hash" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginDTO struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type UserUpdateDTO struct {
	ID    uint64 `json:"id" form:"id"`
	Name  string `json:"name" form:"name" binding:"required"`
	Email string `json:"email" form:"email" binding:"required,email"`
	Role  string `json:"role" form:"role" binding:"required"`
}

type ForgotDTO struct {
	Email string `json:"email" form:"email" binding:"required,email"`
}

type HashDTO struct {
	Hash   string `json:"hash" form:"hash" binding:"required"`
	Action string `json:"action" form:"action" binding:"required"`
}
