package exception
type Exception interface {
	CheckError() bool
	Code() int
	Error() string
}
