

package upgrader

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gofrs/flock"
	"go-agent.wx/pkg/api"
	"go-agent.wx/pkg/config"
	"go-agent.wx/pkg/util/command"
	"go-agent.wx/pkg/util/fileutil"
	"go-agent.wx/pkg/util/systemutil"
	"time"
)

func DoUpgradeAgent() error {
	logs.Info("start upgrade agent")
	config.Init()

	totalLock := flock.New(fmt.Sprintf("%s/%s.lock", systemutil.GetRuntimeDir(), systemutil.TotalLock))
	err := totalLock.Lock()
	if err = totalLock.Lock(); err != nil {
		logs.Error("get total lock failed, exit", err.Error())
		return errors.New("get total lock failed")
	}

	logs.Info("wait 10 seconds for agent to stop")
	time.Sleep(10 * time.Second)

	// GO_20190807 版非windows agent做重装升级替换daemon，其他只替换devopsAgent
	currentAgentVersion := config.GAgentEnv.AgentVersion
	if !systemutil.IsWindows() && currentAgentVersion == "GO_20190807" {
		err := UninstallAgent()
		if err != nil {
			return errors.New("uninstall agent failed")
		}

		fileutil.TryRemoveFile(systemutil.GetWorkDir() + "/agent.zip")
		api.DownloadAgentInstallScript(systemutil.GetWorkDir() + "/" + config.GetInstallScript())

		totalLock.Unlock()
		logs.Info(totalLock.Unlock())
		err = InstallAgent()
		if err != nil {
			logs.Error("install agent failed: ", err)
			return errors.New("install agent failed")
		}

		logs.Info("reinstall agent done, upgrade process exiting")
		return nil
	} else {
		err = replaceAgentFile()
		if err != nil {
			logs.Error("replace agent file failed: ", err.Error())
			return errors.New("replace agent file failed")
		}
		totalLock.Unlock()
	}
	logs.Info("agent upgrade done, upgrade process exiting")
	return nil
}

func DoUninstallAgent() error {
	err := UninstallAgent()
	if err != nil {
		logs.Error("uninstall agent failed: ", err.Error())
		return errors.New("uninstall agent failed")
	}
	return nil
}

func UninstallAgent() error {
	logs.Info("start uninstall agent")

	workDir := systemutil.GetWorkDir()
	startCmd := workDir + "/" + config.GetUninstallScript()
	_, err := command.RunCommand(startCmd, []string{} /*args*/, workDir, nil)
	if err != nil {
		logs.Error("run uninstall script failed: ", err.Error())
		return errors.New("run uninstall script failed")
	}
	return nil
}

func StopAgent() error {
	logs.Info("start stop agent")

	workDir := systemutil.GetWorkDir()
	startCmd := workDir + "/" + config.GetStopScript()
	_, err := command.RunCommand(startCmd, []string{} /*args*/, workDir, nil)
	if err != nil {
		logs.Error("run uninstall script failed: ", err.Error())
		return errors.New("run uninstall script failed")
	}
	return nil
}

func StartAgent() error {
	logs.Info("start agent")

	workDir := systemutil.GetWorkDir()
	startCmd := workDir + "/" + config.GetStartScript()
	_, err := command.RunCommand(startCmd, []string{} /*args*/, workDir, nil)
	if err != nil {
		logs.Error("run uninstall script failed: ", err.Error())
		return errors.New("run uninstall script failed")
	}
	return nil
}

func replaceAgentFile() error {
	logs.Info("replace agent file")
	src := systemutil.GetUpgradeDir() + "/" + config.GetClienAgentFile()
	dst := systemutil.GetWorkDir() + "/" + config.GetClienAgentFile()
	_, err := fileutil.CopyFile(src, dst, true)
	return err
}

func InstallAgent() error {
	logs.Info("start install agent")

	workDir := systemutil.GetWorkDir()
	startCmd := workDir + "/" + config.GetInstallScript()

	err := fileutil.SetExecutable(startCmd)
	if err != nil {
		return fmt.Errorf("chmod install script failed: %s", err.Error())
	}

	_, err = command.RunCommand(startCmd, []string{} /*args*/, workDir, nil)
	if err != nil {
		logs.Error("run install script failed: ", err.Error())
		return errors.New("run install script failed")
	}
	return nil
}
