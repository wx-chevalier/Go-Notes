package state

type op string

const (
	Add op = "add"
	Subtract = "sub"
	Multiply = "mult"
	Divide = "div"
)

type WorkRequest struct {
	Operation op
	Value1    int64
	Value2    int64
}

type WorkResponse struct {
	Wr     *WorkRequest
	Result int64
	Err    error
}
