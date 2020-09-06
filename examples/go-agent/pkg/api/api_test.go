package api

import (
	"testing"

	"go-agent.wx/pkg/config"
)

func loadConfig() {
	config.LoadAgentConfig()
	config.LoadAgentEnv()
}

func Test_buildUrl_01(t *testing.T) {
	loadConfig()
	url := buildUrl("/abc")
	t.Log("url: ", url)
}

func Test_CheckUpgrade_01(t *testing.T) {
	loadConfig()
	data, err := CheckUpgrade()
	if err != nil {
		t.Error("err: ", err.Error())
		return
	}
	t.Log("data: ", data)
}

func Test_AgentStartup_01(t *testing.T) {
	loadConfig()
	data, err := AgentStartup()
	if err != nil {
		t.Error("err: ", err.Error())
		return
	}
	t.Log("data: ", data)
}

func Test_GetAgentStatus_01(t *testing.T) {
	loadConfig()
	data, err := GetAgentStatus()
	if err != nil {
		t.Error("err: ", err.Error())
		return
	}
	t.Log("data: ", data)
}

func Test_GetBuild_01(t *testing.T) {
	loadConfig()
	data, err := GetBuild()
	if err != nil {
		t.Error("err: ", err.Error())
		return
	}
	t.Log("data: ", data)
}

func Test_loadConfig(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loadConfig()
		})
	}
}
