

package pipeline

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/systemutil"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
	"go-agent.wx/pkg/api"
	"go-agent.wx/pkg/util"
	"go-agent.wx/pkg/util/command"
)

func Start() {
	time.Sleep(10 * time.Second)
	for {
		runPipeline()
		time.Sleep(30 * time.Second)
	}
}

func parseAndRunCommandPipeline(pipelineData interface{}) (err error) {
	pipeline := new(CommandPipeline)
	err = util.ParseJsonToData(pipelineData, pipeline)
	if err != nil {
		return errors.New("parse command pipeline failed")
	}

	api.UpdatePipelineStatus(api.NewPipelineResponse(pipeline.SeqId, StatusExecuting, "Executing"))

	if strings.TrimSpace(pipeline.Command) == "" {
		api.UpdatePipelineStatus(api.NewPipelineResponse(pipeline.SeqId, StatusFailure, ""))
		return nil
	}

	lines := strings.Split(pipeline.Command, "\n")
	if len(lines) == 0 {
		api.UpdatePipelineStatus(api.NewPipelineResponse(pipeline.SeqId, StatusFailure, ""))
		return nil
	}

	buffer := bytes.Buffer{}
	if strings.HasPrefix(strings.TrimSpace(lines[0]), "#!") {
		buffer.WriteString(lines[0] + "\n")
		buffer.WriteString("cd " + systemutil.GetWorkDir() + "\n")
	} else {
		buffer.WriteString("#!/bin/bash\n\n")
		buffer.WriteString("cd " + systemutil.GetWorkDir() + "\n")
		buffer.WriteString(lines[0] + "\n")
	}

	for i := 1; i < len(lines); i++ {
		buffer.WriteString(lines[i] + "\n")
	}
	scriptContent := buffer.String()
	logs.Info("scriptContent:", scriptContent)

	scriptFile := fmt.Sprintf("%s/devops_pipeline_%s_%s.sh", systemutil.GetWorkDir(), pipeline.SeqId, pipeline.Type)
	err = ioutil.WriteFile(scriptFile, []byte(scriptContent), 0777)
	if err != nil {
		api.UpdatePipelineStatus(api.NewPipelineResponse(pipeline.SeqId, StatusFailure, "write pipeline script file failed: "+err.Error()))
		return errors.New("write pipeline script file failed: " + err.Error())
	}
	defer os.Remove(scriptFile)

	output, err := command.RunCommand(scriptFile, []string{} /*args*/, systemutil.GetWorkDir(), nil)
	if err != nil {
		api.UpdatePipelineStatus(api.NewPipelineResponse(pipeline.SeqId, StatusFailure, "run pipeline failed: "+err.Error()))
		return errors.New("run pipeline failed: " + err.Error())
	}
	api.UpdatePipelineStatus(api.NewPipelineResponse(pipeline.SeqId, StatusSuccess, string(output)))
	return nil
}

func runPipeline() {
	logs.Info("start run pipeline")
	result, err := api.GetAgentPipeline()
	if err != nil {
		logs.Error("get pipeline failed: ", err.Error())
		return
	}

	if result.IsNotOk() {
		logs.Error("get pipeline failed, message: ", result.Message)
		return
	}

	if result.Data == nil {
		logs.Info("no pipeline to run, skip")
		return
	}

	pipelineData, ok := result.Data.(map[string]interface{})
	if !ok {
		logs.Error("parse pipeline failed, invalid pipeline")
		return
	}

	if pipelineData["type"] == COMMAND {
		err = parseAndRunCommandPipeline(pipelineData)
	} else {
		logs.Warn("not support pipeline: type: ", pipelineData["type"])
		return
	}

	if err != nil {
		logs.Error("run pipeline failed: ", err.Error())
	} else {
		logs.Info("run pipeline done")
	}
}
