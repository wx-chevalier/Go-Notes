

package upgrade

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"go-agent.wx/pkg/api"
	"go-agent.wx/pkg/config"
	"go-agent.wx/pkg/util/command"
	"go-agent.wx/pkg/util/fileutil"
	"go-agent.wx/pkg/util/systemutil"
	"os"
)

func UninstallAgent() {
	logs.Info("start uninstall agent")

	err := runUpgrader(config.ActionUninstall)
	if err != nil {
		logs.Error("start upgrader failed")
		return
	}
	logs.Warning("agent process exiting")
	systemutil.ExitProcess(0)
}

func runUpgrader(action string) error {
	logs.Info("start upgrader process")

	scripPath := systemutil.GetUpgradeDir() + "/" + config.GetClientUpgraderFile()

	if !systemutil.IsWindows() {
		err := os.Chmod(scripPath, 0777)
		if err != nil {
			logs.Error("chmod failed: ", err.Error())
			return errors.New("chmod failed: ")
		}
	}

	if action != config.ActionUninstall {
		action = config.ActionUpgrade
	}
	args := []string{"-action=" + action}

	pid, err := command.StartProcess(scripPath, args, systemutil.GetWorkDir(), nil, "")
	if err != nil {
		logs.Error("run upgrader failed: ", err.Error())
		return errors.New("run upgrader failed")
	}
	logs.Info("start process success, pid: ", pid)

	logs.Warning("agent process exiting")
	systemutil.ExitProcess(0)
	return nil
}

func DoUpgradeOperation(agentChanged bool, workAgentChanged bool) error {
	logs.Info("start upgrade, agent changed: ", agentChanged, ", work agent changed: ", workAgentChanged)
	config.GIsAgentUpgrading = true
	defer func() {
		config.GIsAgentUpgrading = false
		api.FinishUpgrade(true)
	}()

	if !agentChanged && !workAgentChanged {
		logs.Info("no change to upgrade, skip")
		return nil
	}

	if workAgentChanged {
		logs.Info("work agent changed, replace work agent file")
		_, err := fileutil.CopyFile(
			systemutil.GetUpgradeDir()+"/"+config.WorkAgentFile,
			systemutil.GetWorkDir()+"/"+config.WorkAgentFile,
			true)
		if err != nil {
			logs.Error("replace work agent file failed: ", err.Error())
			return errors.New("replace work agent file failed")
		}
		logs.Info("relace agent file done")

		config.GAgentEnv.SlaveVersion = config.DetectWorkerVersion()
	}

	if agentChanged {
		logs.Info("agent changed, start upgrader")
		err := runUpgrader(config.ActionUpgrade)
		if err != nil {
			return err
		}
	} else {
		logs.Info("agent not changed, skip agent upgrade")
	}
	return nil
}
