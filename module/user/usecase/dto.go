package userusecase

type EmailPasswordRegistrationDTO struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type EmailPasswordLoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
