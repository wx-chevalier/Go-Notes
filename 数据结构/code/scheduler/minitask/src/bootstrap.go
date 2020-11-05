/**
 * mini-go-cluster
 * Mini cluster to execute parallel jobs using go lang.
 *
 * @author Waqar Alamgir <bonjour@waqaralamgir.tk>
 */

package main

import (
	fileio "./common"
	logger "./logger"
	executer "./executer"
	"os"
	"strconv"
)

var MAX_CONCURRENT_CONNECTION = 15
var START_TIME_SCRIPT string
var END_TIME_SCRIPT string
var INPUT_FILE string
var OUTPUT_FILE string
var OUTPUT_FILE_FAILED string
var DEBUG_LEVEL_SHORT = 1
var DEBUG_LEVEL_LONG = 2

/**
 * Initize process here
 *
 * @author Waqar Alamgir <bonjour@waqaralamgir.tk>
 */
func init() {
	logger.Log("Enter init" , DEBUG_LEVEL_LONG)

	INPUT_FILE := os.Args[1] // get command line first parameter
	OUTPUT_FILE := os.Args[2] // get command line first parameter
	OUTPUT_FILE_FAILED := os.Args[3] // get command line first parameter

	logger.Log("Input file set: " + INPUT_FILE, DEBUG_LEVEL_LONG)
	logger.Log("Output file set: " + OUTPUT_FILE, DEBUG_LEVEL_LONG)
	logger.Log("Output error file set: " + OUTPUT_FILE_FAILED, DEBUG_LEVEL_LONG)

	logger.Log("Exit init" , DEBUG_LEVEL_LONG)
}

/**
 * Program main function
 *
 * @author Waqar Alamgir <bonjour@waqaralamgir.tk>
 */
func main() {
	logger.Log("Enter main" , DEBUG_LEVEL_LONG)
	START_TIME_SCRIPT = logger.GetTime()
	
	process()

	END_TIME_SCRIPT = logger.GetTime()
	logger.Log("Script started at: " + START_TIME_SCRIPT + " and ended at: " + END_TIME_SCRIPT , DEBUG_LEVEL_LONG)
	logger.Log("Exit main" , DEBUG_LEVEL_LONG)
}

/**
 * Core processing of the script
 *
 * @author Waqar Alamgir <bonjour@waqaralamgir.tk>
 */
func process() {
	logger.Log("Enter process" , DEBUG_LEVEL_LONG)
	
	INPUT_FILE := os.Args[1] // get command line first parameter
	OUTPUT_FILE := os.Args[2] // get command line first parameter
	OUTPUT_FILE_FAILED := os.Args[3] // get command line first parameter

	logger.Log("Creating channels" , DEBUG_LEVEL_LONG)
	// Create the read channel
	readChannel := make(chan string)
	// Create the write channel
	writeChannel := make(chan string)
	// Create the write channel
	failedChannel := make(chan string)
	// Create the thread quit channel
	threadChannel := make(chan bool)

	// Start reading the file and write each line to
	go fileio.ReadFromFile(INPUT_FILE, readChannel)
	go fileio.WriteToFile(OUTPUT_FILE, writeChannel, nil, false)
	go fileio.WriteToFile(OUTPUT_FILE_FAILED, failedChannel, nil, false)

	logger.Log("Staring threads" , DEBUG_LEVEL_SHORT)

	for threads := 0; threads < MAX_CONCURRENT_CONNECTION; threads++ {
		
		logger.Log("Thread " + strconv.Itoa(threads) , DEBUG_LEVEL_LONG)

		go func (dataList chan string, thread chan bool, success chan string, failed chan string, threadId int){

			for data := range dataList {
				// Run either the command or make an http request
				executer.Execute(data, success, failed)
			}
			
			logger.Log("At thread " + strconv.Itoa(threadId) , DEBUG_LEVEL_SHORT)
			thread <- true

		}(readChannel , threadChannel , writeChannel , failedChannel , threads)
	}
	
	logger.Log("Threads started " + strconv.Itoa(MAX_CONCURRENT_CONNECTION) , DEBUG_LEVEL_SHORT)

	// Ensure recommendations for all customer ids have been fetched
	for count := 0; count < MAX_CONCURRENT_CONNECTION; count++ {
		<-threadChannel
	}
}