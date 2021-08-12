package myHttp

import (
	fmt "fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type HTTPClient struct{}

var (
	HttpClient = HTTPClient{}
)
var backoffSchedule = []time.Duration{
	10 * time.Second,
	15 * time.Second,
	20 * time.Second,
	25 * time.Second,
	30 * time.Second,
}

func (c HTTPClient) GetRequest(pathURL string) (*http.Response, error) {
	//fmt.Print("\nVao get request")
	req, _ := http.NewRequest("GET", pathURL, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	c.info(fmt.Sprintf("GET %s -> %d", pathURL, resp.StatusCode))
	if resp.StatusCode != 200 {
		respErr := fmt.Errorf("err Unexpected Response %s", resp.Status)
		_ = fmt.Sprintf("request failed: %v", respErr)
		return nil, respErr
	}
	//fmt.Println("get request cuccess, resp= ",resp)
	return resp, nil
}

func (c HTTPClient) GetRequestWithRetries(url string) (*http.Response, error) {
	//fmt.Print("\nVao get request with retries")
	var body *http.Response
	var err error
	for _, backoff := range backoffSchedule {
		body, err = c.GetRequest(url)
		if err == nil {
			break
		}
		fmt.Println(fmt.Fprintf(os.Stderr, "Request error: %+v\n", err))
		fmt.Println(fmt.Fprintf(os.Stderr, "Retrying in %v\n", backoff))
		time.Sleep(backoff)
	}

	// All retries failed
	if err != nil {
		return nil, err
	}
	//fmt.Println("get request with retries cuccess, body:",body)
	return body, nil
}

func (c HTTPClient) info(msg string) {
	log.Printf("[client] %s\n", msg)
}

func CheckError(err error) {
	if err != nil {
		log.Println(err)
	}
}
