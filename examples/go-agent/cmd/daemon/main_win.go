// +build windows

package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/kardianos/service"
	"go-agent.wx/pkg/util/fileutil"
	"go-agent.wx/pkg/util/systemutil"
)

func main() {
	runtime.GOMAXPROCS(4)

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
			systemutil.ExitProcess(1)
		}
	}()

	logs.Info("devops daemon start")
	logs.Info("pid: ", os.Getpid())
	logs.Info("workDir: ", workDir)

	//服务定义
	serviceConfig := &service.Config{
		Name:             "name",
		DisplayName:      "displayName",
		Description:      "description",
		WorkingDirectory: "C:/data/landun",
	}

	daemonProgram := &program{}
	sys := service.ChosenSystem()
	daemonService, err := sys.New(daemonProgram, serviceConfig)
	if err != nil {
		logs.Error("Init service error: ", err.Error())
		systemutil.ExitProcess(1)
	}

	err = daemonService.Run()
	if err != nil {
		logs.Error("run agent program error: ", err.Error())
	}
}

var GAgentProcess *os.Process = nil

func initLog() {
	logConfig := make(map[string]string)
	logConfig["filename"] = systemutil.GetWorkDir() + "/logs/devopsDaemon.log"
	jsonConfig, _ := json.Marshal(logConfig)
	logs.SetLogger(logs.AdapterFile, string(jsonConfig))
}

func watch() {
	var agentPath = systemutil.GetWorkDir() + "/devopsAgent.exe"
	for {
		cmd := exec.Command(agentPath)
		cmd.Dir = workDir

		logs.Info("start devops agent")
		if !fileutil.Exists(agentPath) {
			logs.Error("agent file: ", agentPath, " not exists")
			logs.Info("restart after 30 seconds")
			time.Sleep(30 * time.Second)
		}

		err := fileutil.SetExecutable(agentPath)
		if err != nil {
			logs.Error("chmod failed, err: ", err.Error())
			logs.Info("restart after 30 seconds")
			time.Sleep(30 * time.Second)
			continue
		}

		err = cmd.Start()
		if err != nil {
			logs.Error("agent start failed, err: ", err.Error())
			logs.Info("restart after 30 seconds")
			time.Sleep(30 * time.Second)
			continue
		}

		GAgentProcess = cmd.Process
		logs.Info("devops agent started, pid: ", cmd.Process.Pid)
		_, err = cmd.Process.Wait()
		if err != nil {
			logs.Error("process wait error", err.Error())
		}
		logs.Info("agent process exited")

		logs.Info("restart after 30 seconds")
		time.Sleep(30 * time.Second)
	}
}

type program struct {
}

func (p *program) Start(s service.Service) error {
	go watch()
	return nil
}

func (p *program) Stop(s service.Service) error {
	p.tryStopAgent()
	return nil
}

func (p *program) tryStopAgent() {
	if GAgentProcess != nil {
		GAgentProcess.Kill()
	}
}
