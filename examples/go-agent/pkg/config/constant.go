package config

import (
	"go-agent.wx/pkg/util/systemutil"
)

const ActionUpgrade = "upgrade"
const ActionUninstall = "uninstall"

const (
	ScriptFileInstallWindows   = "install.bat"
	ScriptFileInstallLinux     = "install.sh"
	ScriptFileUninstallWindows = "uninstall.bat"
	ScriptFileUnistallLinux    = "uninstall.sh"
	ScriptFileStartWindows     = "start.bat"
	ScriptFileStartLinux       = "start.sh"
	ScriptFileStopWindows      = "stop.bat"
	ScriptFileStopLinux        = "stop.sh"
)

const (
	AgentFileClientWindows = "devopsAgent.exe"
	AgentFileClientLinux   = "devopsAgent"
	AgentFileServerWindows = "devopsAgent.exe"
	AgentFileServerLinux   = "devopsAgent_linux"
	AgentFileServerMacos   = "devopsAgent_macos"

	UpgraderFileClientWindows = "upgrader.exe"
	UpgraderFileClientLinux   = "upgrader"
	UpgraderFileServerWindows = "upgrader.exe"
	UpgraderFileServerLinux   = "upgrader_linux"
	UpgraderFileServerMacOs   = "upgrader_macos"

	WorkAgentFile = "worker-agent.jar"
)

// Auth Header
const AuthHeaderSodaProjectId = "X-SODA-PID" //项目ID

const AuthHeaderBuildType = "X-DEVOPS-BUILD-TYPE"       // 构建类型
const AuthHeaderProjectId = "X-DEVOPS-PROJECT-ID"       // 项目ID
const AuthHeaderAgentId = "X-DEVOPS-AGENT-ID"           // Agent ID
const AuthHeaderSecretKey = "X-DEVOPS-AGENT-SECRET-KEY" // Agent密钥
const AuthHeaderPipelineId = "X-DEVOPS-PIPELINE-ID"     //流水线ID
const AuthHeaderBuildId = "X-DEVOPS-BUILD-ID"           //构建ID
const AuthHeaderVmSeqId = "X-DEVOPS-VM-SID"             //VM Seq Id
const AuthHeaderUserId = "X-DEVOPS-UID"                 //用户ID

// 环境变量
const EnvWorkspace = "WORKSPACE"

// 环境类型
const EnvTypeProd = "PROD"
const EnvTypeTest = "TEST"
const EnvTypeDev = "DEV"

// BuildType
const BuildTypeWorkder = "WORKER"
const BuildTypeAgent = "AGENT"
const BuildTypePluginAgent = "PLUGIN_AGENT"
const BuildTypeDocker = "DOCKER"
const BuildTypeDockerHost = "DOCKER_HOST"
const BuildTypeTstack = "TSTACK_AGENT"

// AgentStatus
const AgentStatusUnimport = "UN_IMPORT"
const AgentStatusUnimportOk = "UN_IMPORT_OK"
const AgentStatusImportOk = "IMPORT_OK"
const AgentStatusImportException = "IMPORT_EXCEPTION"
const AgentStatusDelete = "DELETE"

func GetServerAgentFile() string {
	if systemutil.IsWindows() {
		return AgentFileServerWindows
	} else if systemutil.IsMacos() {
		return AgentFileServerMacos
	} else {
		return AgentFileServerLinux
	}
}

func GetServerUpgraderFile() string {
	if systemutil.IsWindows() {
		return UpgraderFileServerWindows
	} else if systemutil.IsMacos() {
		return UpgraderFileServerMacOs
	} else {
		return UpgraderFileServerLinux
	}
}

func GetClienAgentFile() string {
	if systemutil.IsWindows() {
		return AgentFileClientWindows
	} else {
		return AgentFileClientLinux
	}
}

func GetClientUpgraderFile() string {
	if systemutil.IsWindows() {
		return UpgraderFileClientWindows
	} else {
		return UpgraderFileClientLinux
	}
}

func GetInstallScript() string {
	if systemutil.IsWindows() {
		return ScriptFileInstallWindows
	} else {
		return ScriptFileInstallLinux
	}
}

func GetUninstallScript() string {
	if systemutil.IsWindows() {
		return ScriptFileUninstallWindows
	} else {
		return ScriptFileUnistallLinux
	}
}

func GetStartScript() string {
	if systemutil.IsWindows() {
		return ScriptFileStartWindows
	} else {
		return ScriptFileStartLinux
	}
}

func GetStopScript() string {
	if systemutil.IsWindows() {
		return ScriptFileStopWindows
	} else {
		return ScriptFileStopLinux
	}
}
