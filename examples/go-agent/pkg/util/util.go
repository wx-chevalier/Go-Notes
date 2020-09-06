package util

import "encoding/json"

import "fmt"

func ParseJsonToData(jsonData interface{}, targetData interface{}) error {
	dataStr, err := json.Marshal(jsonData)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(dataStr), targetData)
	if err != nil {
		return err
	} else {
		return nil
	}
}

/** Test return single str **/
func Test() {
	fmt.Print("Test")
}
