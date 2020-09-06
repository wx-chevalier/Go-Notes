package config

import (
	"bytes"
	"errors"
	"io/ioutil"
	"strconv"
	"strings"

	bconfig "github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
	"go-agent.wx/pkg/util/command"
	"go-agent.wx/pkg/util/systemutil"
)

const (
	ConfigKeyProjectId     = "devops.project.id"
	ConfigKeyAgentId       = "devops.agent.id"
	ConfigKeySecretKey     = "devops.agent.secret.key"
	ConfigKeyDevopsGateway = "landun.gateway"
	ConfigKeyTaskCount     = "devops.parallel.task.count"
	ConfigKeyEnvType       = "landun.env"
	ConfigKeySlaveUser     = "devops.slave.user"
	ConfigKeyCollectorOn   = "devops.agent.collectorOn"
)

type AgentConfig struct {
	Gateway           string
	BuildType         string
	ProjectId         string
	AgentId           string
	SecretKey         string
	ParallelTaskCount int
	EnvType           string
	SlaveUser         string
	CollectorOn       bool
}

type AgentEnv struct {
	OsName           string
	AgentIp          string
	HostName         string
	SlaveVersion     string
	AgentVersion     string
	AgentInstallPath string
}

var GAgentEnv *AgentEnv
var GAgentConfig *AgentConfig
var GIsAgentUpgrading = false
var GWorkDir string
var GEnvVars map[string]string

func Init() {
	err := LoadAgentConfig()
	if err != nil {
		logs.Error("load agent config err: ", err)
		systemutil.ExitProcess(1)
	}
	LoadAgentEnv()
}

func LoadAgentEnv() {
	GAgentEnv = new(AgentEnv)
	GAgentEnv.AgentIp = systemutil.GetAgentIp()
	GAgentEnv.HostName = systemutil.GetHostName()
	GAgentEnv.OsName = systemutil.GetOsName()
	GAgentEnv.SlaveVersion = DetectWorkerVersion()
	GAgentEnv.AgentVersion = DetectAgentVersion()
}

func DetectAgentVersion() string {
	workDir := systemutil.GetWorkDir()
	output, err := command.RunCommand(workDir+"/"+GetClienAgentFile(), []string{"version"}, workDir, nil)
	if err != nil {
		logs.Warn("detect agent version failed: ", err.Error())
		GAgentEnv.AgentVersion = ""
		return ""
	}
	logs.Info("agent version: ", string(output))

	return strings.TrimSpace(string(output))
}

func DetectWorkerVersion() string {
	output, err := command.RunCommand(GetJava(),
		[]string{"-cp", "agent.jar", "com.tencent.devops.agent.AgentVersionKt"}, systemutil.GetWorkDir(), nil)

	if err != nil {
		logs.Warn("detect worker version failed: ", err.Error())
		GAgentEnv.SlaveVersion = ""
		return ""
	}
	logs.Info("worker version: ", string(output))

	return strings.TrimSpace(string(output))
}

func LoadAgentConfig() error {
	GAgentConfig = new(AgentConfig)

	conf, err := bconfig.NewConfig("ini", systemutil.GetWorkDir()+"/.agent.properties")
	if err != nil {
		logs.Error("load agent config failed, ", err)
		return errors.New("load agent config failed")
	}

	parallelTaskCount, err := conf.Int(ConfigKeyTaskCount)
	if err != nil || parallelTaskCount < 0 {
		return errors.New("invalid parallelTaskCount")
	}

	projectId := strings.TrimSpace(conf.String(ConfigKeyProjectId))
	if len(projectId) == 0 {
		return errors.New("invalid projectId")
	}

	agentId := conf.String(ConfigKeyAgentId)
	if len(agentId) == 0 {
		return errors.New("invalid agentId")
	}

	secretKey := strings.TrimSpace(conf.String(ConfigKeySecretKey))
	if len(secretKey) == 0 {
		return errors.New("invalid secretKey")
	}

	landunGateway := strings.TrimSpace(conf.String(ConfigKeyDevopsGateway))
	if len(landunGateway) == 0 {
		return errors.New("invalid landunGateway")
	}

	envType := strings.TrimSpace(conf.String(ConfigKeyEnvType))
	if len(envType) == 0 {
		return errors.New("invalid envType")
	}

	slaveUser := strings.TrimSpace(conf.String(ConfigKeySlaveUser))
	if len(slaveUser) == 0 {
		slaveUser = systemutil.GetCurrentUser().Username
	}

	collectorOn, err := conf.Bool(ConfigKeyCollectorOn)
	if err != nil {
		collectorOn = true
	}

	GAgentConfig.Gateway = landunGateway
	logs.Info("Gateway: ", GAgentConfig.Gateway)
	GAgentConfig.BuildType = BuildTypeAgent
	logs.Info("BuildType: ", GAgentConfig.BuildType)
	GAgentConfig.ProjectId = projectId
	logs.Info("ProjectId: ", GAgentConfig.ProjectId)
	GAgentConfig.AgentId = agentId
	logs.Info("AgentId: ", GAgentConfig.AgentId)
	GAgentConfig.SecretKey = secretKey
	logs.Info("SecretKey: ", GAgentConfig.SecretKey)
	GAgentConfig.EnvType = envType
	logs.Info("EnvType: ", GAgentConfig.EnvType)
	GAgentConfig.ParallelTaskCount = parallelTaskCount
	logs.Info("ParallelTaskCount: ", GAgentConfig.ParallelTaskCount)
	GAgentConfig.SlaveUser = slaveUser
	logs.Info("SlaveUser: ", GAgentConfig.SlaveUser)
	GAgentConfig.CollectorOn = collectorOn
	logs.Info("CollectorOn: ", GAgentConfig.CollectorOn)
	return nil
}

func (a *AgentConfig) SaveConfig() error {
	filePath := systemutil.GetWorkDir() + "/.agent.properties"

	systemutil.IsWindows()
	content := bytes.Buffer{}
	content.WriteString(ConfigKeyProjectId + "=" + GAgentConfig.ProjectId + "\n")
	content.WriteString(ConfigKeyAgentId + "=" + GAgentConfig.AgentId + "\n")
	content.WriteString(ConfigKeySecretKey + "=" + GAgentConfig.SecretKey + "\n")
	content.WriteString(ConfigKeyDevopsGateway + "=" + GAgentConfig.Gateway + "\n")
	content.WriteString(ConfigKeyTaskCount + "=" + strconv.Itoa(GAgentConfig.ParallelTaskCount) + "\n")
	content.WriteString(ConfigKeyEnvType + "=" + GAgentConfig.EnvType + "\n")
	content.WriteString(ConfigKeySlaveUser + "=" + GAgentConfig.SlaveUser + "\n")

	err := ioutil.WriteFile(filePath, []byte(content.String()), 0666)
	if err != nil {
		logs.Error("write config failed:", err.Error())
		return errors.New("write config failed")
	}
	return nil
}

func (a *AgentConfig) GetAuthHeaderMap() map[string]string {
	authHeaderMap := make(map[string]string)
	authHeaderMap[AuthHeaderBuildType] = a.BuildType
	authHeaderMap[AuthHeaderSodaProjectId] = a.ProjectId
	authHeaderMap[AuthHeaderProjectId] = a.ProjectId
	authHeaderMap[AuthHeaderAgentId] = a.AgentId
	authHeaderMap[AuthHeaderSecretKey] = a.SecretKey
	return authHeaderMap
}

func GetJava() string {
	workDir := systemutil.GetWorkDir()
	if systemutil.IsMacos() {
		return workDir + "/jre/Contents/Home/bin/java"
	} else {
		return workDir + "/jre/bin/java"
	}
}
