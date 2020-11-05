package executer

import (
    httpbuilder "../httpbuilder"
    commandrunner "../commandrunner"
    logger "./../logger"
    "strings"
)

var DEBUG_LEVEL_SHORT = 1
var DEBUG_LEVEL_LONG = 2

/**
 * Decides and execute weather we need to run command or make an http request
 *
 * @author Waqar Alamgir <bonjour@waqaralamgir.tk>
 */
func Execute(data string, success chan string, failed chan string) {

    arr := strings.Split(data, "|")
    
    // If 4 items in array
    if len(arr) == 4 {
        jobType := arr[0]

        logger.Log("JobType: " + jobType, DEBUG_LEVEL_LONG)
        
        // Job type is URL, then send a http request
        if jobType == "URL" {
            // Dispatch the request
            httpbuilder.Request(data, success, failed)
        } else if jobType == "CMD" {
            // Run the command on the server
            commandrunner.Run(data, success, failed)
        } else {
            // Job type is invalid
            logger.Log("Invalid jobType: " + jobType, DEBUG_LEVEL_LONG)
        }

    } else {
        logger.Log("Invalid input: " + data, DEBUG_LEVEL_LONG)
    }
}