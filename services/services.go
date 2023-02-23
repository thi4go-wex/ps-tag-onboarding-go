package services

type ServiceErrorT int

type ServiceErrorReply struct {
	ErrorCode    ServiceErrorT `json:"errorCode"`
	ErrorContext string        `json:"errorContext"`
}

func (e ServiceErrorReply) Error() string {
	return e.ErrorContext
}
