package errno

import "fmt"

type ErrNo struct {
	Code    int
	Message string
}

func (e *ErrNo) Error() string {
	return e.Message
}

// Err represents on error
type Err struct {
	Code    int
	Message string
	Err     error
}

func New(errNo *ErrNo, err error) *Err {
	return &Err{Code: errNo.Code, Message: errNo.Message, Err: err}
}

func (err *Err) Add(message string) error {
	err.Message += " " + message
	return err
}

func (err *Err) Addf(format string, args ...interface{}) error {
	//return err.Message = fmt.Sprintf("%s %s", err.Message, fmt.Sprintf(format, args...))
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

func IsErrUserNotFound(err error) bool {
	code, _ := DecodeErr(err)
	return code == ErrUserNotFound.Code
}

func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *ErrNo:
		return typed.Code, typed.Message
	default:
	}

	return InternalServerError.Code, err.Error()
}
