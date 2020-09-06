

package command

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"os"
	"os/exec"
)

func RunCommand(command string, args []string, workDir string, envMap map[string]string) (output []byte, err error) {
	cmd := exec.Command(command)

	if len(args) > 0 {
		cmd.Args = append(cmd.Args, args...)
	}

	if workDir != "" {
		cmd.Dir = workDir
	}

	cmd.Env = os.Environ()
	if envMap != nil {
		for k, v := range envMap {
			cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
		}
	}

	logs.Info("cmd.Path: ", cmd.Path)
	logs.Info("cmd.Args: ", cmd.Args)
	logs.Info("cmd.workDir: ", cmd.Dir)

	outPut, err := cmd.Output()
	logs.Info("output: ", string(outPut))
	if err != nil {
		return outPut, err
	}

	return outPut, nil
}

func StartProcess(command string, args []string, workDir string, envMap map[string]string, runUser string) (int, error) {
	cmd := exec.Command(command)

	if len(args) > 0 {
		cmd.Args = append(cmd.Args, args...)
	}

	if workDir != "" {
		cmd.Dir = workDir
	}

	cmd.Env = os.Environ()
	if envMap != nil {
		for k, v := range envMap {
			cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
		}
	}

	err := setUser(cmd, runUser)
	if err != nil {
		logs.Error("set user failed: ", err.Error())
		return -1, errors.New("set user failed")
	}

	logs.Info("cmd.Path: ", cmd.Path)
	logs.Info("cmd.Args: ", cmd.Args)
	logs.Info("cmd.workDir: ", cmd.Dir)
	logs.Info("runUser: ", runUser)

	err = cmd.Start()
	if err != nil {
		return -1, err
	}
	return cmd.Process.Pid, nil
}
