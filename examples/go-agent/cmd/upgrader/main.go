

package main

import (
	"encoding/json"
	"flag"
	"github.com/astaxie/beego/logs"
	"os"
	"go-agent.wx/pkg/config"
	"go-agent.wx/pkg/upgrader"
	"go-agent.wx/pkg/util/systemutil"
	"runtime"
)

const (
	upgraderProcess = "upgrade"
)

func main() {
	runtime.GOMAXPROCS(4)
	initLog()
	defer func() {
		if err := recover(); err != nil {
			logs.Error("panic: ", err)
		}
	}()

	if ok := systemutil.CheckProcess(upgraderProcess); !ok {
		logs.Warn("get process lock failed, exit")
		return
	}

	action := flag.String("action", "upgrade", "action, upgrade or uninstall")
	flag.Parse()
	logs.Info("upgrader start, action: ", *action)
	logs.Info("pid: ", os.Getpid())
	logs.Info("current user userName: ", systemutil.GetCurrentUser().Username)
	logs.Info("work dir: ", systemutil.GetWorkDir())

	if config.ActionUpgrade == *action {
		err := upgrader.DoUpgradeAgent()
		if err != nil {
			logs.Error("upgrade agent failed")
			systemutil.ExitProcess(1)
		}
	} else if config.ActionUninstall == *action {
		err := upgrader.DoUninstallAgent()
		if err != nil {
			logs.Error("upgrade agent failed")
			systemutil.ExitProcess(1)
		}
	} else {
		logs.Error("unsupport action")
		systemutil.ExitProcess(1)
	}
	systemutil.ExitProcess(0)
}

func initLog() {
	logConfig := make(map[string]string)
	logConfig["filename"] = systemutil.GetWorkDir() + "/logs/devopsUpgrader.log"
	logConfig["perm"] = "0666"
	jsonConfig, _ := json.Marshal(logConfig)
	logs.SetLogger(logs.AdapterFile, string(jsonConfig))
}
