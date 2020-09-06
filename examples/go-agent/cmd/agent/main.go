package main

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/astaxie/beego/logs"
	"go-agent.wx/pkg/agent"
	"go-agent.wx/pkg/config"
	"go-agent.wx/pkg/util/systemutil"
)

const (
	agentProcess = "agent"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "version" {
		fmt.Println(config.AgentVersion)
		systemutil.ExitProcess(0)
	}

	runtime.GOMAXPROCS(4)

	// 以 agent 安装目录为工作目录
	workDir := systemutil.GetExecutableDir()
	err := os.Chdir(workDir)
	if err != nil {
		logs.Info("change work dir failed, err: ", err.Error())
		systemutil.ExitProcess(1)
	}

	initLog()
	defer func() {
		if err := recover(); err != nil {
			logs.Error("panic: ", err)
		}
	}()

	if ok := systemutil.CheckProcess(agentProcess); !ok {
		logs.Warn("get process lock failed, exit")
		return
	}

	logs.Info("agent start")
	logs.Info("pid: ", os.Getpid())
	logs.Info("agent version: ", config.AgentVersion)
	logs.Info("current user userName: ", systemutil.GetCurrentUser().Username)
	logs.Info("work dir: ", systemutil.GetWorkDir())

	logEnv()

	agent.Run()
}

func initLog() {
	logConfig := make(map[string]string)
	logConfig["filename"] = systemutil.GetWorkDir() + "/logs/devopsAgent.log"
	logConfig["perm"] = "0666"
	jsonConfig, _ := json.Marshal(logConfig)
	logs.SetLogger(logs.AdapterFile, string(jsonConfig))
}

func logEnv() {
	logs.Info("agent envs: ")
	for _, v := range os.Environ() {
		index := strings.Index(v, "=")
		logs.Info("    " + v[0:index] + " = " + v[index+1:])
	}
}
