package api

import (
	"fmt"
	"strconv"
	"strings"

	"go-agent.wx/pkg/config"
	"go-agent.wx/pkg/util/httputil"
	"go-agent.wx/pkg/util/systemutil"
)

func buildUrl(url string) string {
	if strings.HasPrefix(config.GAgentConfig.Gateway, "http") {
		return config.GAgentConfig.Gateway + url
	} else {
		return "http://" + config.GAgentConfig.Gateway + url
	}
}

func Heartbeat(buildInfos []ThirdPartyBuildInfo) (*httputil.DevopsResult, error) {
	url := buildUrl("/ms/environment/api/buildAgent/agent/thirdPartyAgent/agents/newHeartbeat")

	agentHeartbeatInfo := &AgentHeartbeatInfo{
		MasterVersion:     config.AgentVersion,
		SlaveVersion:      config.GAgentEnv.SlaveVersion,
		HostName:          config.GAgentEnv.HostName,
		AgentIp:           config.GAgentEnv.AgentIp,
		ParallelTaskCount: config.GAgentConfig.ParallelTaskCount,
		AgentInstallPath:  systemutil.GetExecutableDir(),
		StartedUser:       systemutil.GetCurrentUser().Username,
		TaskList:          buildInfos,
	}

	return httputil.NewHttpClient().Post(url).Body(agentHeartbeatInfo).SetHeaders(config.GAgentConfig.GetAuthHeaderMap()).Execute().IntoDevopsResult()
}

func CheckUpgrade() (*httputil.AgentResult, error) {
	url := buildUrl("/ms/dispatch/api/buildAgent/agent/thirdPartyAgent/upgrade?version=" + config.GAgentEnv.SlaveVersion + "&masterVersion=" + config.AgentVersion)
	return httputil.NewHttpClient().Get(url).SetHeaders(config.GAgentConfig.GetAuthHeaderMap()).Execute().IntoAgentResult()
}

func FinishUpgrade(success bool) (*httputil.AgentResult, error) {
	url := buildUrl("/ms/dispatch/api/buildAgent/agent/thirdPartyAgent/upgrade?success=" + strconv.FormatBool(success))
	return httputil.NewHttpClient().Delete(url).SetHeaders(config.GAgentConfig.GetAuthHeaderMap()).Execute().IntoAgentResult()
}

func DownloadUpgradeFile(serverFile string, saveFile string) (fileMd5 string, err error) {
	url := buildUrl("/ms/environment/api/buildAgent/agent/thirdPartyAgent/upgrade/files/download?file=" + serverFile)
	return httputil.DownloadUpgradeFile(url, config.GAgentConfig.GetAuthHeaderMap(), saveFile)
}

func DownloadAgentInstallScript(saveFile string) error {
	url := buildUrl(fmt.Sprintf("/external/agents/%s/install", config.GAgentConfig.AgentId))
	return httputil.DownloadAgentInstallScript(url, config.GAgentConfig.GetAuthHeaderMap(), saveFile)
}

func AgentStartup() (*httputil.DevopsResult, error) {
	url := buildUrl("/ms/environment/api/buildAgent/agent/thirdPartyAgent/startup")

	startInfo := &ThirdPartyAgentStartInfo{
		HostName:      config.GAgentEnv.HostName,
		HostIp:        config.GAgentEnv.AgentIp,
		DetectOs:      config.GAgentEnv.OsName,
		MasterVersion: config.AgentVersion,
		SlaveVersion:  config.GAgentEnv.SlaveVersion,
	}

	return httputil.NewHttpClient().Post(url).Body(startInfo).SetHeaders(config.GAgentConfig.GetAuthHeaderMap()).Execute().IntoDevopsResult()
}

func GetAgentStatus() (*httputil.DevopsResult, error) {
	url := buildUrl("/ms/environment/api/buildAgent/agent/thirdPartyAgent/status")
	return httputil.NewHttpClient().Get(url).SetHeaders(config.GAgentConfig.GetAuthHeaderMap()).Execute().IntoDevopsResult()
}

func GetBuild() (*httputil.AgentResult, error) {
	url := buildUrl("/ms/dispatch/api/buildAgent/agent/thirdPartyAgent/startup")
	return httputil.NewHttpClient().Get(url).SetHeaders(config.GAgentConfig.GetAuthHeaderMap()).Execute().IntoAgentResult()
}

func WorkerBuildFinish(buildInfo *ThirdPartyBuildWithStatus) (*httputil.DevopsResult, error) {
	url := buildUrl("/ms/dispatch/api/buildAgent/agent/thirdPartyAgent/workerBuildFinish")
	return httputil.NewHttpClient().Post(url).Body(buildInfo).SetHeaders(config.GAgentConfig.GetAuthHeaderMap()).Execute().IntoDevopsResult()
}

func GetAgentPipeline() (*httputil.DevopsResult, error) {
	url := buildUrl("/ms/environment/api/buildAgent/agent/thirdPartyAgent/agents/pipelines")
	return httputil.NewHttpClient().Get(url).SetHeaders(config.GAgentConfig.GetAuthHeaderMap()).Execute().IntoDevopsResult()
}

func UpdatePipelineStatus(response *PipelineResponse) (*httputil.DevopsResult, error) {
	url := buildUrl("/ms/environment/api/buildAgent/agent/thirdPartyAgent/agents/pipelines")
	return httputil.NewHttpClient().Put(url).Body(response).SetHeaders(config.GAgentConfig.GetAuthHeaderMap()).Execute().IntoDevopsResult()
}
