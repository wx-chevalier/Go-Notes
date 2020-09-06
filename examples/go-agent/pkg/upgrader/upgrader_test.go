

package upgrader

import (
	"go-agent.wx/pkg/config"
	"testing"
)

func init() {
	config.Init()
}

func Test_DoUpgradeAgent_01(t *testing.T) {
	err := DoUpgradeAgent()
	if err != nil {
		t.Error("err: ", err.Error())
	}
	t.Log("done")
}

func Test_DoUninstallAgent_01(t *testing.T) {
	err := DoUninstallAgent()
	if err != nil {
		t.Error("err: ", err.Error())
	}
	t.Log("done")
}
