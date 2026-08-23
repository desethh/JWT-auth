package authorization

type SignUpRequestDTO struct {
	Name     string
	Email    string
	Password string
}

type LoginInRequestDTO struct {
	Name     string
	Email    string
	Password string
}
