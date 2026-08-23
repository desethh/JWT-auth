package authorization

import "fmt"

var (
	ErrBodyIsNil        = fmt.Errorf("request body is nil")
	ErrEmailIsEmpty     = fmt.Errorf("email is empty")
	ErrEmailNameIsEmpty = fmt.Errorf("email and name are empty")
	ErrEmailName        = fmt.Errorf("wrong email format")
	ErrPasswordIsEmpty  = fmt.Errorf("password is empty")
	ErrNameIsEmpty      = fmt.Errorf("name is empty")
)
