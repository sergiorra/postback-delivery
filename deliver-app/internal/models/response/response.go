package response

import "time"

// Response representation of response into struct
type Response struct {
	Code   			int
	Body   			string
	Time   			time.Time
	DeliveryTime 	time.Time
}

// NewResponse initialize response
func NewResponse(code int, body string, time, deliveryTime time.Time) *Response {
	return &Response{
		Code: code,
		Body: body,
		Time: time,
		DeliveryTime: deliveryTime,
	}
}