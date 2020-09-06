package command

import (
	"testing"
)

func Test_RunCommand_01(t *testing.T) {
	output, err := RunCommand("ipconfig", []string{"/all"}, "", nil)
	if err != nil {
		t.Error("err: ", err.Error())
	}
	t.Log("output: ", string(output))
}

func Test_RunCommand_02(t *testing.T) {
	output, err := RunCommand("bash", []string{"/Users/huangou/workspace/agent/test/devops_pipeline_oamyqvmd_COMMAND.sh"}, "/Users/huangou/workspace/agent/test", nil)
	if err != nil {
		t.Error("err: ", err.Error())
	}
	t.Log("output: ", string(output))
}

func Test_StartProcess_01(t *testing.T) {
	_, err := StartProcess("/a/tme.exe", nil, "", nil, "")
	if err != nil {
		t.Error("err: ", err.Error())
	}
}
