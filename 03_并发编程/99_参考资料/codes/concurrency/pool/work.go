package pool

import "errors"

type op string

const (
	Hash op = "encrypt"
	Compare = "decrypt"
)

type WorkRequest struct {
	Op      op
	Text    []byte
	Compare []byte 
}

type WorkResponse struct {
	Wr      WorkRequest
	Result  []byte
	Matched bool
	Err     error
}

func Process(wr WorkRequest) WorkResponse {
	switch wr.Op {
	case Hash:
		return hashWork(wr)
	case Compare:
		return compareWork(wr)
	default:
		return WorkResponse{Err: errors.New("unsupported operation")}
	}
}
