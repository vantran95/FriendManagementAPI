package exception

type Exception struct {
	Code    int
	Message string
}

func (e Exception) Error() string {
	panic("Implement me")
}

func (e Exception) Exception() string {
	return e.Message
}
