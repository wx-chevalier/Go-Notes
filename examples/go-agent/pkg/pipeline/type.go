

package pipeline

var COMMAND = "COMMAND"
var FILE = "FILE"

const (
	StatusPending   = "PENDING"
	StatusExecuting = "EXECUTING"
	StatusTimeout   = "TIMEOUT"
	StatusSuccess   = "SUCCESS"
	StatusFailure   = "FAILURE"
)

type CommandPipeline struct {
	SeqId   string `json:"seqId"`
	Type    string `json:"type"`
	Command string `json:"command"`
}

type FilePipeline struct {
	SeqId     string `json:"seqId"`
	Type      string `json:"type"`
	Operation string `json:"operation"`
	File      string `json:"file"`
}
