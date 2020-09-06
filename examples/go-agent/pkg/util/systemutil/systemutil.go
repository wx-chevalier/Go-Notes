package systemutil

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gofrs/flock"
	"go-agent.wx/pkg/util/fileutil"
	"net"
	"os"
	"os/user"
	"runtime"
	"strings"
)

var GExecutableDir string

const (
	osWindows = "windows"
	osLinux   = "linux"
	osMacos   = "darwin"
	amd64     = "amd64"
	osOther   = "other"

	osNameWindows = "windows"
	osNameLinux   = "linux"
	osNameMacos   = "macos"
	osNameOther   = "other"

	TotalLock = "total-lock"
)

func IsWindows() bool {
	return runtime.GOOS == osWindows
}

func IsLinux() bool {
	return runtime.GOOS == osLinux
}

func IsMacos() bool {
	return runtime.GOOS == osMacos
}

func IsAmd64() bool {
	return runtime.GOARCH == amd64
}

func GetCurrentUser() *user.User {
	currentUser, _ := user.Current()
	return currentUser
}

func GetWorkDir() string {
	dir, _ := os.Getwd()
	return strings.Replace(dir, "\\", "/", -1)
}

func GetUpgradeDir() string {
	return GetWorkDir() + "/tmp"
}

func GetRuntimeDir() string {
	runDir := fmt.Sprintf("%s/runtime", GetWorkDir())
	if err := os.MkdirAll(runDir, 0755); err == nil {
		return runDir
	}
	return GetWorkDir()
}

func GetExecutableDir() string {
	if len(GExecutableDir) == 0 {
		executable := strings.Replace(os.Args[0], "\\", "/", -1)
		index := strings.LastIndex(executable, "/")
		GExecutableDir = executable[0:index]
	}
	return GExecutableDir
}

func GetOsName() string {
	switch runtime.GOOS {
	case osLinux:
		return osNameLinux
	case osWindows:
		return osNameWindows
	case osMacos:
		return osNameMacos
	default:
		return osNameOther
	}
}

func GetOs() string {
	switch runtime.GOOS {
	case osLinux:
		return osLinux
	case osWindows:
		return osWindows
	case osMacos:
		return osMacos
	default:
		return osOther
	}
}

func GetHostName() string {
	name, _ := os.Hostname()
	return name
}

func GetAgentIp() string {
	defaultIp := "127.0.0.1"
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return defaultIp
	}
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			if ipNet.IP.IsGlobalUnicast() {
				return ipNet.IP.String()
			}
		}
	}
	return defaultIp
}

func ExitProcess(exitCode int) {
	os.Exit(exitCode)
}

var processLock *flock.Flock

func KeepProcessAlive() {
	runtime.KeepAlive(*processLock)
}

func CheckProcess(name string) bool {
	processLockFile := fmt.Sprintf("%s/%s.lock", GetRuntimeDir(), name)
	pidFile := fmt.Sprintf("%s/%s.pid", GetRuntimeDir(), name)

	processLock = flock.New(processLockFile)
	ok, err := processLock.TryLock()
	if err != nil {
		logs.Error("failed to get process lock(%s), exit: %v", processLockFile, err)
		return false
	}

	if !ok {
		logs.Error("failed to get process lock(%s), exit: maybe already running.", processLockFile)
		return false
	}

	totalLock := flock.New(fmt.Sprintf("%s/%s.lock", GetRuntimeDir(), TotalLock))
	if err = totalLock.Lock(); err != nil {
		logs.Error("get total lock failed, exit", err.Error())
		return false
	}
	defer func() {
		_ = totalLock.Unlock()
	}()

	if err = fileutil.WriteString(pidFile, fmt.Sprintf("%d", os.Getpid())); err != nil {
		logs.Error("failed to save pid file(%s): %v", pidFile, err)
		return false
	}

	logs.Info("success to get process lock and save pid")
	return true
}
