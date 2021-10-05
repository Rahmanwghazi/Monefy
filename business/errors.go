package business

import "errors"

var (
	ErrorInvalidSigninInfo = errors.New("Username or password is invalid")
	ErrorUnauthorized      = errors.New("User Unauthorized")
	ErrorDuplicateUsername = errors.New("Username has already been taken")
	ErrorDuplicateEmail    = errors.New("Email has already been taken")
	ErrorInternal          = errors.New("An Error Has Occured ")
	ErrorEmpty             = errors.New("The Field cannot be empty ")
)
