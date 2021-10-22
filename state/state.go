package state

type Op string

const (
	Add       = "add"
	Subtract  = "sub"
	Multiply  = "mul"
	Divide    = "div"
)

type WorkRequest struct {
	Operation Op
	Value1    int64
	Value2    int64
}

type WorkResponse struct {
	Wr     *WorkRequest
	Result int64
	Err    error
}
