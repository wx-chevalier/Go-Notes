package request

/**
 * The following package contains a simple url caller function.
 *
 * @author Fahad Zia Syed <fzia@folio3.com>
 */
import (
	"fmt"
	"strconv"
	"io/ioutil"
	"net/http"
	"time"
)

/**
 * The generic dispatch function which calls URL
 * and returns the parsed response
 */
func Dispatch(url string, method string, retry int) (bool, []byte) {

	if (retry == 1) {
		fmt.Println("Dispatch URL [" + url + "]")
	}

	// Get the response
	resp, err := http.Get(url)
	
	// If there was error in response
	// try to retry
	if err != nil {
		// Retry 10 times
		if retry <= 25 {
			
			retry++

			var newUrl string

			// Adding sleep for [retry] times
			duration := time.Duration(retry)*time.Second
			time.Sleep(duration)

			// For the frist time add a retry
			if (retry == 2) {
				newUrl = url + "&retry=" + strconv.Itoa(retry)
			} else {
				// Just replace the number next time
				newUrl = url[:len(url)-1] + strconv.Itoa(retry)
			}

			fmt.Println("URL [" + newUrl + "] connection error, retying at: " + strconv.Itoa(retry))

			return Dispatch(newUrl, method, retry)
		}

		//if failed after retry. send the customer id to failed channel
		fmt.Println("No response for ", url)

		return false, nil
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	// if connection to database failed then retry
	if string(body[:]) == "{\"status\":\"503\",\"message\":\"MySQL database connection failed.\"}" {
  		
  		// Retry 10 times
		if retry <= 10 {
			
			retry++

			var newUrl string

			// Adding sleep for [retry] times
			duration := time.Duration(retry)*time.Second
			time.Sleep(duration)

			// For the frist time add a retry
			if (retry == 2) {
				newUrl = url + "&retry=" + strconv.Itoa(retry)
			} else {
				// Just replace the number next time
				newUrl = url[:len(url)-1] + strconv.Itoa(retry)
			}

			fmt.Println("URL [" + newUrl + "] MySQL database connection error, retying at: " + strconv.Itoa(retry))

			return Dispatch(newUrl, method, retry)
		}
	}

	return_bool := true
	if err != nil {

		fmt.Println("Error reading response")
		return_bool = false
	}

	return return_bool, body
}