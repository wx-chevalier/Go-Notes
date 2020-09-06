

package job

import (
	"fmt"
	"testing"
	"time"
)

func Test_BuildManager_01(t *testing.T) {
	fmt.Println("start")
	GBuildManager.AddBuild(6124, &ThirdPartyBuildInfo{})
	for {
		time.Sleep(5 * time.Second)
		fmt.Println("instanceCount: ", GBuildManager.GetInstanceCount())
	}
}
