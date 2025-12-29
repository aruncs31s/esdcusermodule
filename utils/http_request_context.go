package utils

type HTTPRequestContext struct {
	RequestID  string
	UserID     string
	UserRole   string
	Parameters *Parameters
}
type Parameters struct {
	Limit  int
	Offset int
}
