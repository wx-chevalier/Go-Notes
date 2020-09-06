// +build linux darwin



package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gofrs/flock"
	"os"
	"os/exec"
	"go-agent.wx/pkg/util/fileutil"
	"go-agent.wx/pkg/util/systemutil"
	"runtime"
	"time"

	"encoding/json"
	"go-agent.wx/pkg/config"
)

const (
	daemonProcess = "daemon"
	agentCheckGap = 5 * time.Second
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

	if ok := systemutil.CheckProcess(daemonProcess); !ok {
		logs.Warn("get process lock failed, exit")
		return
	}

	logs.Info("devops daemon start")
	logs.Info("pid: ", os.Getpid())

	watch()
	systemutil.KeepProcessAlive()
}

func initLog() {
	logConfig := make(map[string]string)
	logConfig["filename"] = systemutil.GetWorkDir() + "/logs/devopsDaemon.log"
	logConfig["perm"] = "0666"
	jsonConfig, _ := json.Marshal(logConfig)
	logs.SetLogger(logs.AdapterFile, string(jsonConfig))
}

func watch() {
	totalLock := flock.New(fmt.Sprintf("%s/%s.lock", systemutil.GetRuntimeDir(), systemutil.TotalLock))

	// first check immediately
	totalLock.Lock()
	doCheckAndLaunchAgent()
	totalLock.Unlock()

	checkTimeTicker := time.NewTicker(agentCheckGap)
	for ; ; totalLock.Unlock() {
		select {
		case <-checkTimeTicker.C:
			if err := totalLock.Lock(); err != nil {
				logs.Error("failed to get agent lock: %v", err)
				continue
			}

			doCheckAndLaunchAgent()
		}
	}
}

func doCheckAndLaunchAgent() {
	workDir := systemutil.GetWorkDir()
	agentLock := flock.New(fmt.Sprintf("%s/agent.lock", systemutil.GetRuntimeDir()))

	ok, err := agentLock.TryLock()
	if err != nil {
		logs.Error("try to get agent.lock failed: %v", err)
		return
	}
	if !ok {
		return
	}
	logs.Warn("agent is not available, will launch it")
	_ = agentLock.Unlock()

	process, err := launch(workDir + "/" + config.AgentFileClientLinux)
	if err != nil {
		logs.Error("launch agent failed: %v", err)
		return
	}
	if process == nil {
		logs.Error("launch agent failed: got a nil process")
		return
	}
	logs.Info("success to launch agent, pid: %d", process.Pid)
}

func launch(agentPath string) (*os.Process, error) {
	cmd := exec.Command(agentPath)
	cmd.Dir = systemutil.GetWorkDir()

	logs.Info("start devops agent: %s", agentPath)
	if !fileutil.Exists(agentPath) {
		return nil, fmt.Errorf("agent file %s not exists", agentPath)
	}

	err := fileutil.SetExecutable(agentPath)
	if err != nil {
		return nil, fmt.Errorf("chmod agent file failed: %v", err)
	}

	if err = cmd.Start(); err != nil {
		return nil, fmt.Errorf("start agent failed: %v", err)
	}

	go func() {
		_ = cmd.Wait()
	}()

	return cmd.Process, nil
}
