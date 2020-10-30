package response

import "time"

// Response representation of response into struct
type Response struct {
	code   			int
	body   			string
	time   			time.Time
	deliveryTime 	time.Time
}

// NewResponse initialize response
func NewResponse(code int, body string, time, deliveryTime time.Time) *Response {
	return &Response{
		code: code,
		body: body,
		time: time,
		deliveryTime: deliveryTime,
	}
}