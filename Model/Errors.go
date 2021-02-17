package Model

type CustomError struct {
	Code      int
	ErrorMsg  string
	OrigError error
}

func (err CustomError) Error() string {
	return err.ErrorMsg
}
func (err CustomError) StatusCode() int {
	return err.Code
}

func (err *CustomError) SetOrigError(origError error) {
	err.OrigError = origError
}

var BAD_REQUEST = CustomError{Code: 400, ErrorMsg: "Bad Request. Unable to parse Request"}
var NOT_FOUND = CustomError{Code: 404, ErrorMsg: "No records Found"}
var DB_UNREACHABLE = CustomError{Code: 500, ErrorMsg: "Internal Server Error. Unable to connect to DB"}
var DB_ERROR = CustomError{Code: 500, ErrorMsg: "Internal Server Error. Error accessing to DB"}
