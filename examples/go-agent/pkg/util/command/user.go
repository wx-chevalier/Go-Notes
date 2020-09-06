// +build linux darwin



package command

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"os/exec"
	"os/user"
	"go-agent.wx/pkg/util/systemutil"
	"strconv"
	"syscall"
)

func setUser(cmd *exec.Cmd, runUser string) error {
	logs.Info("set user(linux or darwin): ", runUser)
	if len(runUser) == 0 || runUser == systemutil.GetCurrentUser().Username {
		return nil
	}

	user, err := user.Lookup(runUser)
	if err != nil {
		logs.Error("user lookup failed, user: -", runUser, "-, error: ", err.Error())
		return errors.New("user lookup failed, user: " + runUser)
	}
	uid, _ := strconv.Atoi(user.Uid)
	gid, _ := strconv.Atoi(user.Gid)
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(uid), Gid: uint32(gid)}

	cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", "HOME", user.HomeDir))
	cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", "USER", runUser))
	cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", "USERNAME", runUser))
	cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", "LOGNAME", runUser))

	return nil
}
