package commandrunner

import (
    logger "../logger"
    "os/exec"
    "strings"
	"runtime"
	"fmt"
)

var DEBUG_LEVEL_SHORT = 1
var DEBUG_LEVEL_LONG = 2

/**
 * Execute command on server
 *
 * @author Waqar Alamgir <bonjour@waqaralamgir.tk>
 */
func Run(data string, success chan string, failed chan string) {

    arr := strings.Split(data, "|")
    
    commandPath := arr[1]
    param := arr[2]
    params := arr[3]
	
	fmt.Println("OS: " + runtime.GOOS)
	
	// Excute the command
    if runtime.GOOS == "windows" {
		out, _ := exec.Command("cmd", "/c", commandPath + " " + params).Output()
		s := string(out)
		s = strings.Replace(s, "\r"," ",-1)
		s = strings.Replace(s, "\n"," ",-1)
		s = strings.Replace(s, "  "," ",-1)
		output := commandPath + "|" + param + "|" + params + "|" + s
		logger.Log("Success command " + commandPath + params , DEBUG_LEVEL_LONG)
		success <- output
	} else {
		out, _ := exec.Command("sh", "-c", commandPath + " " + params).Output()
		s := string(out)
		s = strings.Replace(s, "\r"," ",-1)
		s = strings.Replace(s, "\n"," ",-1)
		s = strings.Replace(s, "  "," ",-1)
		output := commandPath + "|" + param + "|" + params + "|" + s
		logger.Log("Success command " + commandPath + params , DEBUG_LEVEL_LONG)
		success <- output
	}
}