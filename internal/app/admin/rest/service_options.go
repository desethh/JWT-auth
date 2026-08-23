package rest

type Options func(a *App)

func WithAuthService(AuthService AuthService) Options {
	return func(a *App) {
		a.AuthService = AuthService
	}
}
