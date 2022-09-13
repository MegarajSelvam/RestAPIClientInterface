package RestAPIClient

import (
	"net/http"
	"time"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	Client = &http.Client{
		Timeout: time.Second * 50,
	}
}

// With Interface
// We can write unit test for this method. We can mock Client easily
func WebCallV1(request *http.Request) (*http.Response, error) {
	return Client.Do(request)
}

// Without Interface
// We cannot write unit test for this method. We don't have any option to mock httpClient
func WebCallV2(request *http.Request) (*http.Response, error) {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	return httpClient.Do(request)
}
