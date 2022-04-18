package exception

type NotFoundException struct {
	Message string
}

func (r NotFoundException) Error() string {
	return r.Message
}

func NewNotFoundException() NotFoundException {
	return NotFoundException{Message: NotFoundError}
}
