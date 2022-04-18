package exception

type OutOfBoundsException struct {
	Message string
}

func (r OutOfBoundsException) Error() string {
	return r.Message
}

func NewOutOfBoundsException() OutOfBoundsException {
	return OutOfBoundsException{Message: OutOfBoundsError}
}
