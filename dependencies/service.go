package dependencies

import authorization "jwt/internal/app/auth"

type Dependencies struct {
	AuthService *authorization.AuthService
}

func NewDependencies(authService *authorization.AuthService) Dependencies {
	return Dependencies{
		AuthService: authService,
	}
}
