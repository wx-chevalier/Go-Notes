package client

import (
	"log"

	model "dev.wx/model"

	iReq "github.com/imroc/req"
)

const (
	defaultBpServerHost string = "172.16.7.249:8080"
)

type BpClient struct {
	bpServerHost string
}

func NewBpClient(host string) (*BpClient, error) {

	bpClient := &BpClient{}

	bpClient.bpServerHost = host

	if host == "" {
		bpClient.bpServerHost = defaultBpServerHost
	}

	return bpClient, nil
}

/* 创建为新的任务 **/
func (bpClient *BpClient) CreateTask(deviceInfo, bppParaInfo *model.File, fileList []*model.File) (*iReq.Resp, error) {

	body := make(map[string]interface{})
	body["deviceInfo"] = *deviceInfo
	body["bppParaInfo"] = *bppParaInfo
	body["fileList"] = fileList

	bodyJson := iReq.BodyJSON(body)

	log.Println(deviceInfo)

	r, err := iReq.Post("http://"+bpClient.bpServerHost+"/createTask", bodyJson)

	if err != nil {
		log.Fatal(err)
	}

	var foo interface{}
	// response => struct/map
	r.ToJSON(&foo)

	// print info (try it, you may surprise)
	log.Printf("%+v", r)

	return r, err
}
