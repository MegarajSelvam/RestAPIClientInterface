package Common

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func ReadContent(content io.ReadCloser) []byte {
	body, _ := ioutil.ReadAll(content)
	return body
}

var (
	RequestString  = `{"Id":141}`
	ResponseString = `{"firstName":"Megaraj","full_name":"Megaraj Selvam","Id":141}`
)

func GetResponse() (*http.Response, error) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(ResponseString)))
	return &http.Response{
		StatusCode: 200,
		Body:       r,
	}, nil
}

func GetRequest() *http.Request {
	r := ioutil.NopCloser(bytes.NewReader([]byte(RequestString)))
	url, _ := url.Parse("https://localhost:8080/api/v1/test")
	return &http.Request{
		Body:   r,
		Method: "Post",
		URL:    url,
	}
}
