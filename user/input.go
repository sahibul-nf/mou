package user

type RegisterUserInput struct {
	Name         string `json:"name" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	Occupation   string `json:"occupation" binding:"required"`
	PasswordHash string `json:"password" binding:"required"`
}
