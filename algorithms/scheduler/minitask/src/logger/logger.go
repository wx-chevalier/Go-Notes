package logger

import (
	"fmt"
	"os"
	"time"
)

var DEBUG_LEVEL = 2
var LOG_FILE = os.Args[4] + "log_" + GetTime() + ".log"

/**
 * Logs message in a file
 * 
 * @param string message Message content
 * @param int level Level of the message 1, 2 , 3 etc
 * @author Waqar Alamgir <bonjour@waqaralamgir.tk>
 */
func Log(message string, level int) {

	if level <= DEBUG_LEVEL {
		f, err := os.OpenFile(LOG_FILE , os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
		if err != nil {
    		panic(err)
		}
		defer f.Close()
    	f.WriteString(message + "\n")
    	f.Sync()
    	fmt.Println(message)
	}
}

/**
 * Returns current time
 *
 * @author Waqar Alamgir <bonjour@waqaralamgir.tk>
 * @return string Current Time
 */
func GetTime() string {
	t := time.Now()
	// return t.Format(time.RFC3339)
	return t.Format("20060102150405")
}