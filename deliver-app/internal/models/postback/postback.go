package postback

import (
	"net/url"
	"regexp"

	"github.com/sergiorra/postback-delivery/deliver-app/internal/models/enpoint"
)

// Postback representation of postback into struct
type Postback struct {
	Endpoint	endpoint.Endpoint   	`json:"endpoint"`
	Data   		[]map[string]string 	`json:"data"`
}

// MountURL replaces all query params by matching with each key, leaves empty string if it doesn't find any match
func (p *Postback) MountURL() {
	for k, v := range p.Data[0] {
		v = url.QueryEscape(v)
		re := regexp.MustCompile(regexp.QuoteMeta("{" + k + "}"))
		p.Endpoint.Url = re.ReplaceAllString(p.Endpoint.Url, v)
	}
	re := regexp.MustCompile("{.*?}")
	p.Endpoint.Url = re.ReplaceAllString(p.Endpoint.Url, "")
}