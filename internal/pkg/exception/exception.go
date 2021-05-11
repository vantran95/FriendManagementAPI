package exception

// Exception stores info to retrieve exception struct
type Exception struct {
	Code    int
	Message string
}

// Exception attempts to retrieve an error message.
func (e *Exception) Exception() string {
	return e.Message
}
