package authorization

import (
	"regexp"
)

func validateSignUpRequest(d *SignUpRequestDTO) error {
	if d == nil {
		return ErrBodyIsNil
	}
	switch {
	case d.Name == "":
		return ErrNameIsEmpty
	case d.Email == "":
		return ErrEmailIsEmpty
	case d.Password == "":
		return ErrPasswordIsEmpty
	}

	// mail string validation
	if !isEmailValid(d.Email) {
		return ErrEmailName
	}
	return nil
}

func validateLoginInRequest(d *LoginInRequestDTO) error {
	if d == nil {
		return ErrBodyIsNil
	}
	switch {
	case d.Email == "" && d.Name == "":
		return ErrEmailNameIsEmpty
	case d.Password == "":
		return ErrPasswordIsEmpty
	}
	return nil
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
