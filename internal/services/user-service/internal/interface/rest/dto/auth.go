package dto

type RegisterUserRequest struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `form:"email" json:"email" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
}

type LoginUserRequest struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type AuthResponse struct {
	AccessToken      string `json:"accessToken"`
	AccessExpiresIn  string `json:"expiresIn"`
	RefreshToken     string `json:"refreshToken"`
	RefreshExpiresIn string `json:"refreshExpiresIn"`
}

type IssueTokens struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}
