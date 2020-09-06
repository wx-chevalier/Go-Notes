

package httputil

import (
	"testing"
	"time"
)

func Test_downloadFile_01(t *testing.T) {
	t.Log("start: ", time.Now())
	_, err := DownloadUpgradeFile("https://services.gradle.org/distributions/gradle-5.2.1-bin.zip", nil, "/Users/huangou/a.txt")
	if err != nil {
		t.Error("err: ", err.Error())
		return
	}
	t.Log("end: ", time.Now())
}
