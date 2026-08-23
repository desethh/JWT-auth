package rest

import authorization "jwt/internal/app/auth"

type AuthService interface {
	SignUp(d authorization.SignUpRequestDTO) error
}

type App struct {
	AuthService AuthService
}

func NewAdminService(opts ...Options) *App {
	app := new(App)
	for _, opt := range opts {
		opt(app)
	}
	return app
}
