package agent

import (
	"github.com/astaxie/beego/logs"

	"go-agent.wx/pkg/collector"
	"go-agent.wx/pkg/config"
	"go-agent.wx/pkg/heartbeat"
	"go-agent.wx/pkg/job"
	"go-agent.wx/pkg/pipeline"
	"go-agent.wx/pkg/upgrade"
)

func Run() {
	config.Init()

	_, err := job.AgentStartup()
	if err != nil {
		logs.Warn("agent startup failed: ", err.Error())
	}

	// 数据采集
	go collector.DoAgentCollect()

	// 心跳
	go heartbeat.DoAgentHeartbeat()

	// 检查升级
	go upgrade.DoPollAndUpgradeAgent()

	// 启动pipeline
	go pipeline.Start()

	job.DoPollAndBuild()
}
