

package upgrade

import (
	"testing"

	"go-agent.wx/pkg/config"
)

func Test_startUpgrader_01(t *testing.T) {
	err := runUpgrader(config.ActionUpgrade)
	if err != nil {
		t.Error("err: ", err.Error())
		return
	}
	t.Log("success")
}

func Test_startUpgrader_02(t *testing.T) {
	err := runUpgrader(config.ActionUninstall)
	if err != nil {
		t.Error("err: ", err.Error())
		return
	}
	t.Log("success")
}
