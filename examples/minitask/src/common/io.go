package file

/**
 * The following package contains basic file writing functions.
 */

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/**
 * The function takes in the inputfile name
 * and the channel to where we want to send the customer_ids
 * The function is intended to run as a goroutine thus returns nothing
 *
 * @param file_name - string - File Name of the string that should be read
 * @param line_channel - channel string - The channel to write the lines to
 */
func ReadFromFile(file_name string, line_channel chan string) {
	//start reading file
	content, err := ioutil.ReadFile(file_name)
	//check if there was error in file reading
	if err != nil {
		//display error if file could not be read
		fmt.Println("File Read Error:", err)
	}

	//parse bytes[] to string
	s := string(content)

	//split lines for each \n
	lines := strings.SplitAfter(string(s), "\n")
	line_count := 0

	//write each line to the channel
	for _, line := range lines {
		//write to the channel
		line_channel <- strings.Trim(line, "\n\r ")
		//increment the counter
		line_count++
	}

	//close the channel
	//this signals the listeners to close
	//the loop
	close(line_channel)

	//log the total lines
	fmt.Println("Total Lines: %d", line_count)
}

/**
 * The function listens to the string channel 'result_channel' abd writes to the file present at
 * file_name. A bool value 'true' is sent after the result_channel is closed (Optional). If append
 * is true, the lines are appended to the file present at file_name
 * IMPORTANT : If append is not set or is set to false, file present at file_name will be deleted if
 * it exists
 *
 * @param file_name - string - The file to write
 * @param result_channel - chan string - The channel to listen
 * @param quit_channel - chan string - A bool will be sent when the result_channel is closed.
 * @param append - bool - Whether to append while writing to file_name or not
 */
func WriteToFile(file_name string, result_channel chan string, quit_channel chan bool, append bool) {

	//flags to open file with
	var flags int

	if append == false {
		//Since we are not supposed to append to the file.
		//We simply deleted it
		DeleteFile(file_name)
		//Open the file with create and write flags.
		flags = (os.O_CREATE | os.O_WRONLY)
	} else {
		//The file will open with append, will be created if it does not exist
		flags = (os.O_APPEND | os.O_CREATE | os.O_WRONLY)
	}

	//open the file
	f, err2 := os.OpenFile(file_name, flags, 0600)

	defer f.Close()

	if err2 != nil {
		//there was an error in opening file.
		fmt.Println("Error opening file")
	} else {
		//successfully opened file
		//Listen on the channel to
		for line := range result_channel {
			//write the string received from the channel
			_, err3 := f.WriteString(line)
			_, err4 := f.WriteString("\n")

			if err3 != nil {
				fmt.Println("Error writing file")
			}
			if err4 != nil {
				fmt.Println("Error writing file")
			}
		}
		//if the quit_channel is set.
		//send the true bool
		if quit_channel != nil {
			//send signal that write
			//operation has completed.
			quit_channel <- true
		}
	}

}

/**
 * Removes the file if it exists.
 *
 * @param filename - string - The file to remove
 */
func DeleteFile(filename string) {
	//Check if the file is present by opening it
	dataFile, _ := os.Open(filename)
	if dataFile != nil {
		//the file exists.
		//Delete the file
		dataFile.Close()
		os.Remove(filename)
	}
}