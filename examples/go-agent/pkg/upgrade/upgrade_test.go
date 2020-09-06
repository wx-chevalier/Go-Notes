

package upgrade

import (
	"go-agent.wx/pkg/config"
	"testing"
)

func loadConfig() {
	config.LoadAgentConfig()
	config.LoadAgentEnv()
}

func Test_downloadFile_01(t *testing.T) {
	loadConfig()
	agentChanged, workAgentChanged, err := downloadFile()
	if err != nil {
		t.Error("err: ", err.Error())
		return
	}
	t.Log("agentChanged: ", agentChanged, ", workAgentChanged: ", workAgentChanged)
}
