package simplehttp

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// SimpleHTTP class for creating HTTP requests
type SimpleHTTP struct {
	BaseURL string
	Headers []HTTPHeader
}

// HTTPHeader a HTTP header
type HTTPHeader struct {
	Key   string
	Value string
}

func (simpleHttp *SimpleHTTP) request(method string, url string, body io.Reader) *http.Request {

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatal(err)
	}

	for _, header := range simpleHttp.Headers {
		req.Header.Add(header.Key, header.Value)
	}

	return req
}

func (simpleHttp *SimpleHTTP) response(req *http.Request, v interface{}) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode >= 400 {
		log.Fatal(resp.Status)
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal(body, v)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
}

// SetHeaders sets SimpleHttps's headers
func (simpleHttp *SimpleHTTP) SetHeaders(headers []HTTPHeader) {
	simpleHttp.Headers = headers
}

// DoRequest fires a HTTP request with an optional body to the url using the specified method
// and returns unmarshalled json response body in respStruct
func (simpleHttp *SimpleHTTP) DoRequest(method, url string, body io.Reader, respStruct interface{}) {
	req := simpleHttp.request(method, simpleHttp.BaseURL+url, body)
	simpleHttp.response(req, respStruct)
}

// Get GET data from the specified url and unmarshall the json into respStruct
func (simpleHttp *SimpleHTTP) Get(url string, respStruct interface{}) {
	simpleHttp.DoRequest("GET", url, nil, respStruct)
}

// Post POST data in body to the specified url and unmarshall the json response into respStruct
func (simpleHttp *SimpleHTTP) Post(url string, body io.Reader, respStruct interface{}) {
	simpleHttp.DoRequest("POST", url, body, respStruct)
}

// ToReader creates an io.Reader from byteArr
func ToReader(byteArr []byte) io.Reader {
	return bytes.NewReader(byteArr)
}

func main() {}
