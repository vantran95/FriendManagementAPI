package exception

type Exception struct {
	Code    int
	Message string
}

func (e *Exception) Exception() string {
	return e.Message
}
