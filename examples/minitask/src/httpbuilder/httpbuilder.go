package httpbuilder

import (
	urlreq "../url"
	logger "../logger"
	"strings"
)

var DEBUG_LEVEL_SHORT = 1
var DEBUG_LEVEL_LONG = 2

/**
 * Generates an http request
 *
 * @author Waqar Alamgir <bonjour@waqaralamgir.tk>
 */
func Request(data string, success chan string, failed chan string) {

	arr := strings.Split(data, "|")
    
    url := arr[1]
    method := arr[2]
    params := arr[3]

    // dispatch the request
    errb, resp := urlreq.Dispatch(url + params, method, 1)

    if errb == true {
        output := url + "|" + method + "|" + params + "|" + string(resp[:])
        logger.Log("Success url " + url + params , DEBUG_LEVEL_LONG)
        // logger.Log("Output " + output , DEBUG_LEVEL_LONG)
        success <- output
    } else {logger.Log("Invalid type: " + data , DEBUG_LEVEL_LONG)
        logger.Log("Failed url " + url + params, DEBUG_LEVEL_LONG)
        failed <- data
    }
}