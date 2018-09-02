package errno

var (
	// Common errors
	OK                  = &ErrNo{Code: 0, Message: "OK"}
	InternalServerError = &ErrNo{Code: 10001, Message: "Internal server error."}
	ErrBind             = &ErrNo{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	ErrValidation = &ErrNo{Code: 20001, Message: "Validation failed"}
	ErrDatabase   = &ErrNo{Code: 20002, Message: "Database error"}
	ErrToken      = &ErrNo{Code: 20003, Message: "Error occurred while signing the JSON web token"}

	// user errors
	ErrEncrypt           = &ErrNo{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound      = &ErrNo{Code: 20102, Message: "The user was not found."}
	ErrTokenInvalid      = &ErrNo{Code: 20103, Message: "The token wa invalid."}
	ErrPasswordIncorrect = &ErrNo{Code: 20104, Message: "The password was incorrect."}
)
