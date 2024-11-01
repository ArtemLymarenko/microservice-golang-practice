package dto

type AuthProjectResponseDto struct {
	ProjectToken string `json:"projectToken"`
	ExpiresIn    string `json:"expiresIn"`
}
