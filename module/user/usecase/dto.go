package userusecase

type EmailPasswordRegistrationDTO struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Status      string `json:"status"`
}

type EmailPasswordLoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponseDTO struct {
	AccessToken       string `json:"access_token"`
	AccessTokenExpIn  int    `json:"access_token_exp_in"`
	RefreshToken      string `json:"refresh_token"`
	RefreshTokenExpIn int    `json:"refresh_token_exp_in"`
}
