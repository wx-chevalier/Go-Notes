

package upgrade

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"os"
	"go-agent.wx/pkg/api"
	"go-agent.wx/pkg/config"
	"go-agent.wx/pkg/util/fileutil"
	"go-agent.wx/pkg/util/systemutil"
	"time"
)

func DoPollAndUpgradeAgent() {
	for {
		time.Sleep(20 * time.Second)
		logs.Info("try upgrade")
		agentUpgrade()
		logs.Info("upgrade done")
	}
}

func agentUpgrade() {
	checkResult, err := api.CheckUpgrade()
	if err != nil {
		logs.Error("check upgrade err: ", err.Error())
		return
	}
	if !checkResult.IsOk() {
		logs.Error("check upgrade failed: ", checkResult.Message)
		return
	}

	if checkResult.IsAgentDelete() {
		logs.Info("agent is deleted, skip")
		return
	}

	if !(checkResult.Data).(bool) {
		logs.Info("no need to upgrade agent, skip")
		return
	}

	logs.Info("download upgrade files start")
	agentChanged, workerChanged, err := downloadUpgradeFiles()
	if err != nil {
		logs.Error("download upgrade files failed", err.Error())
		return
	}
	logs.Info("download upgrade files done")

	err = DoUpgradeOperation(agentChanged, workerChanged)
	if err != nil {
		logs.Error("do upgrade operation failed", err)
	}
}

func downloadUpgradeFiles() (agentChanged bool, workAgentChanged bool, err error) {
	workDir := systemutil.GetWorkDir()
	upgradeDir := systemutil.GetUpgradeDir()
	os.MkdirAll(upgradeDir, os.ModePerm)

	logs.Info("download upgrader start")
	_, err = api.DownloadUpgradeFile("upgrade/"+config.GetServerUpgraderFile(), upgradeDir+"/"+config.GetClientUpgraderFile())
	if err != nil {
		logs.Error("download upgrader failed", err)
		return false, false, errors.New("download upgrader failed")
	}
	logs.Info("download upgrader done")

	logs.Info("download agent start")
	newAgentMd5, err := api.DownloadUpgradeFile("upgrade/"+config.GetServerAgentFile(), upgradeDir+"/"+config.GetClienAgentFile())
	if err != nil {
		logs.Error("download agent failed", err)
		return false, false, errors.New("download agent failed")
	}
	logs.Info("download agent done")

	logs.Info("download worker start")
	newWorkerMd5, err := api.DownloadUpgradeFile("jar/"+config.WorkAgentFile, upgradeDir+"/"+config.WorkAgentFile)
	if err != nil {
		logs.Error("download worker failed", err)
		return false, false, errors.New("download worker failed")
	}
	logs.Info("download worker done")

	agentMd5, err := fileutil.GetFileMd5(workDir + "/" + config.GetClienAgentFile())
	if err != nil {
		logs.Error("check agent md5 failed", err)
		return false, false, errors.New("check agent md5 failed")
	}
	workerMd5, err := fileutil.GetFileMd5(workDir + "/" + config.WorkAgentFile)
	if err != nil {
		logs.Error("check worker md5 failed", err)
		return false, false, errors.New("check agent md5 failed")
	}

	return agentMd5 != newAgentMd5, workerMd5 != newWorkerMd5, nil
}
