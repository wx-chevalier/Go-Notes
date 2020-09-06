

package job

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"os"
	"go-agent.wx/pkg/api"
)

type buildManager struct {
	instances map[int]*api.ThirdPartyBuildInfo
}

var GBuildManager *buildManager

func init() {
	GBuildManager = new(buildManager)
	GBuildManager.instances = make(map[int]*api.ThirdPartyBuildInfo)
}

func (b *buildManager) GetInstanceCount() int {
	return len(b.instances)
}

func (b *buildManager) GetInstances() []api.ThirdPartyBuildInfo {
	result := make([]api.ThirdPartyBuildInfo, 0)
	for _, value := range b.instances {
		result = append(result, *value)
	}
	return result
}

func (b *buildManager) AddBuild(processId int, buildInfo *api.ThirdPartyBuildInfo) {
	bytes, _ := json.Marshal(buildInfo)
	logs.Info("add build: processId: ", processId, ", buildInfo: ", string(bytes))
	b.instances[processId] = buildInfo
	go b.waitProcessDone(processId)
}

func (b *buildManager) waitProcessDone(processId int) {
	process, err := os.FindProcess(processId)
	if err != nil {
		errMsg := fmt.Sprintf("build process err, pid: %d, err: %s", processId, err.Error())
		logs.Warn(errMsg)
		delete(b.instances, processId)
		workerBuildFinish(&api.ThirdPartyBuildWithStatus{
			*b.instances[processId],
			false,
			"errMsg"})
		return
	}

	process.Wait()
	logs.Info("build process finish: pid: ", processId)
	buildInfo := b.instances[processId]
	delete(b.instances, processId)
	workerBuildFinish(&api.ThirdPartyBuildWithStatus{*buildInfo, true, "success"})
}
