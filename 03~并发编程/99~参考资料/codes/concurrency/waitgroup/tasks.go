package waitgroup

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func GetURL(url string) (*http.Response, error) {
	start := time.Now()
	log.Printf("getting %s", url)
	resp, err := http.Get(url)
	log.Printf("completed getting %s in %s", url, time.Since(start))
	return resp, err
}

type CrawlError struct {
	Errors []string
}

func (c *CrawlError) Add(err error) {
	c.Errors = append(c.Errors, err.Error())
}
func (c *CrawlError) Error() string {
	return fmt.Sprintf("All Errors: %s", strings.Join(c.Errors, ","))
}
func (c *CrawlError) Valid() bool {
	return len(c.Errors) != 0
}
