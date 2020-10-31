package request

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sergiorra/postback-delivery/deliver-app/internal/models/response"
)

// request representation of request into struct
type request struct {
	url 	string
	method 	string
	Body 	[]byte
}

const (
	ContentType = "application/json"
)

// NewRequest initialize request
func NewRequest(url, method string) *request {
	return &request{
		url: url,
		method: method,
	}
}

// Get sends a GET request to an specific endpoint and returns response info
func (req *request) Get() (*response.Response, error) {
	deliveryTime := time.Now()
	res, err := http.Get(req.url)
	if err != nil {
		return &response.Response{}, err
	}
	r, err := req.extractData(res, deliveryTime)
	if err != nil {
		return &response.Response{}, err
	}
	return r, nil
}

// Post sends a POST request to an specific endpoint and returns response info
func (req *request) Post() (*response.Response, error) {
	deliveryTime := time.Now()
	res, err := http.Post(req.url, ContentType, bytes.NewBuffer(req.Body))
	if err != nil {
		return &response.Response{}, err
	}
	r, err := req.extractData(res, deliveryTime)
	if err != nil {
		return &response.Response{}, err
	}
	return r, nil
}

// extractData extracts data from an http response
func (req *request) extractData(res *http.Response, deliveryTime time.Time) (*response.Response, error) {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &response.Response{}, err
	}
	r := response.NewResponse(res.StatusCode, string(body), time.Now(), deliveryTime)
	return r, nil
}
