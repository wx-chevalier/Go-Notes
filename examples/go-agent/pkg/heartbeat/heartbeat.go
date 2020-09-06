package heartbeat

import (
	"errors"
	"time"

	"github.com/astaxie/beego/logs"
	"go-agent.wx/pkg/api"
	"go-agent.wx/pkg/config"
	"go-agent.wx/pkg/job"
	"go-agent.wx/pkg/upgrade"
	"go-agent.wx/pkg/util"
)

func DoAgentHeartbeat() {
	for {
		agentHeartbeat()

		time.Sleep(10 * time.Second)
	}
}

func agentHeartbeat() error {
	result, err := api.Heartbeat(job.GBuildManager.GetInstances())
	if err != nil {
		logs.Error("agent heartbeat failed: ", err.Error())
		return errors.New("agent heartbeat failed")
	}
	if result.IsNotOk() {
		logs.Error("agent heartbeat failed: ", result.Message)
		return errors.New("agent heartbeat failed")
	}

	heartbeatResponse := new(api.AgentHeartbeatResponse)
	err = util.ParseJsonToData(result.Data, &heartbeatResponse)
	if err != nil {
		logs.Error("agent heartbeat failed: ", err.Error())
		return errors.New("agent heartbeat failed")
	}

	if heartbeatResponse.AgentStatus == config.AgentStatusDelete {
		upgrade.UninstallAgent()
		return nil
	}

	// 修改agent配置
	if config.GAgentConfig.ParallelTaskCount != heartbeatResponse.ParallelTaskCount {
		config.GAgentConfig.ParallelTaskCount = heartbeatResponse.ParallelTaskCount
		config.GAgentConfig.SaveConfig()
	}

	// agent环境变量
	config.GEnvVars = heartbeatResponse.Envs
	logs.Info("agent heartbeat done")
	return nil
}
