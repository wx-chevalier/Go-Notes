package api

type ThirdPartyAgentStartInfo struct {
	HostName      string `json:"hostname"`
	HostIp        string `json:"hostIp"`
	DetectOs      string `json:"detectOS"`
	MasterVersion string `json:"masterVersion"`
	SlaveVersion  string `json:"version"`
}

type ThirdPartyBuildInfo struct {
	ProjectId  string `json:"projectId"`
	BuildId    string `json:"buildId"`
	VmSeqId    string `json:"vmSeqId"`
	Workspace  string `json:"workspace"`
	PipelineId string `json:"pipelineId"`
}

type ThirdPartyBuildWithStatus struct {
	ThirdPartyBuildInfo
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type PipelineResponse struct {
	SeqId    string `json:"seqId"`
	Status   string `json:"status"`
	Response string `json:"response"`
}

type AgentHeartbeatInfo struct {
	MasterVersion     string                `json:"masterVersion"`
	SlaveVersion      string                `json:"slaveVersion"`
	HostName          string                `json:"hostName"`
	AgentIp           string                `json:"agentIp"`
	ParallelTaskCount int                   `json:"parallelTaskCount"`
	AgentInstallPath  string                `json:"agentInstallPath"`
	StartedUser       string                `json:"startedUser"`
	TaskList          []ThirdPartyBuildInfo `json:"taskList"`
}

type AgentHeartbeatResponse struct {
	MasterVersion     string            `json:"masterVersion"`
	SlaveVersion      string            `json:"slaveVersion"`
	AgentStatus       string            `json:"agentStatus"`
	ParallelTaskCount int               `json:"parallelTaskCount"`
	Envs              map[string]string `json:"envs"`
}

func NewPipelineResponse(seqId string, status string, response string) *PipelineResponse {
	return &PipelineResponse{
		SeqId:    seqId,
		Status:   status,
		Response: response,
	}
}
