package exception

type MismatchTypeException struct {
	Message string
}

func (r MismatchTypeException) Error() string {
	return r.Message
}

func NewMismatchException() MismatchTypeException {
	return MismatchTypeException{Message: MismatchedTypeError}
}
