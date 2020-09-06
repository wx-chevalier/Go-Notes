

package job

import (
	"os"
	"go-agent.wx/pkg/api"
	"testing"
)

func Test_writeStartBuildAgentScript_01(t *testing.T) {
	buildInfo := &api.ThirdPartyBuildInfo{"pid", "bid", "1", ""}
	file, err := writeStartBuildAgentScript(buildInfo)
	if err != nil {
		t.Error("error: ", err.Error())
	}
	dir, _ := os.Getwd()
	t.Log("workDir: ", dir)
	t.Log("fileName: ", file)
}
