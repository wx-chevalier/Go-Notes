package httputil

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"go-agent.wx/pkg/util/fileutil"

	"github.com/astaxie/beego/logs"
	"go-agent.wx/pkg/config"
)

type DevopsResult struct {
	Data    interface{} `json:"data"`
	Status  int64       `json:"status"`
	Message string      `json:"message"`
}

func (d *DevopsResult) IsOk() bool {
	return d.Status == 0
}

func (d *DevopsResult) IsNotOk() bool {
	return d.Status != 0
}

type AgentResult struct {
	DevopsResult
	AgentStatus string `json:"agentStatus"`
}

func (a *AgentResult) IsAgentDelete() bool {
	if a.AgentStatus == "" {
		return false
	}
	return a.AgentStatus == config.AgentStatusDelete
}

func (r *HttpResult) IntoDevopsResult() (*DevopsResult, error) {
	if nil != r.Error {
		return nil, r.Error
	}

	result := new(DevopsResult)
	err := json.Unmarshal(r.Body, result)
	if nil != err {
		logs.Error("parse result error: ", err.Error())
		return nil, errors.New("parse result error")
	} else {
		return result, nil
	}
}

func (r *HttpResult) IntoAgentResult() (*AgentResult, error) {
	if nil != r.Error {
		return nil, r.Error
	}

	result := new(AgentResult)
	err := json.Unmarshal(r.Body, result)
	if nil != err {
		logs.Error("parse result error: ", err.Error())
		return nil, errors.New("parse result error")
	} else {
		return result, nil
	}
}

func DownloadAgentInstallScript(url string, headers map[string]string, filepath string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()
	if err != nil {
		logs.Error("download agent install script failed", err)
		return errors.New("download agent install script failed")
	}

	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		body, _ := ioutil.ReadAll(resp.Body)
		logs.Error("download agent install script failed, status: " + resp.Status + ", responseBody: " + string(body))
		return errors.New("download agent install script failed")
	}

	err = writeToFile(filepath, resp.Body)
	if err != nil {
		logs.Error("write agent install script failed", err)
		return errors.New("write agent install script failed")
	}
	return nil
}

func DownloadUpgradeFile(url string, headers map[string]string, filepath string) (md5 string, err error) {
	oldFileMd5, err := fileutil.GetFileMd5(filepath)
	if err != nil {
		logs.Error("check file md5 failed", err)
		return "", errors.New("check file md5 failed")
	}
	if oldFileMd5 != "" {
		url = url + "&eTag=" + oldFileMd5
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// header
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()
	if err != nil {
		logs.Error("download upgrade file failed", err)
		return "", errors.New("download upgrade file failed")
	}

	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		if resp.StatusCode == http.StatusNotFound {
			return "", errors.New("file not found")
		}
		if resp.StatusCode == http.StatusNotModified {
			return oldFileMd5, nil
		}
		body, _ := ioutil.ReadAll(resp.Body)
		logs.Error("download upgrade file failed, status: " + resp.Status + ", responseBody: " + string(body))
		return "", errors.New("download upgrade file failed")
	}

	err = writeToFile(filepath, resp.Body)
	if err != nil {
		logs.Error("download upgrade file failed", err)
		return "", errors.New("download upgrade file failed")
	}

	fileMd5, err := fileutil.GetFileMd5(filepath)
	logs.Info("download file md5: ", fileMd5)
	if err != nil {
		logs.Error("check file md5 failed", err)
		return "", errors.New("check file md5 failed")
	}

	checksumMd5 := resp.Header.Get("X-Checksum-Md5")
	logs.Info("checksum md5: ", checksumMd5)
	if len(checksumMd5) > 0 && checksumMd5 != fileMd5 {
		return "", errors.New("file md5 not match")
	}

	return fileMd5, err
}

func writeToFile(file string, content io.Reader) error {
	out, err := os.Create(file)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, content)
	if err != nil {
		logs.Error("save file failed", err)
		return err
	}
	return nil
}
