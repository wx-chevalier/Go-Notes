// +build windows



package command

import (
	"github.com/astaxie/beego/logs"
	"os/exec"
)

func setUser(cmd *exec.Cmd, runUser string) error {
	logs.Info("set user(windows): ", runUser)
	return nil
}
