package auth_usecase

type SignInDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenDTO struct {
	UserId  string
	IsAdmin bool
}

// структура представляет собой пару токенов, готовых к отправке
type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
