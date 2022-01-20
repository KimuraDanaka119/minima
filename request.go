package minima

import (
	"encoding/json"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

/**
@info The request param structure
@property {string} [Path] Route path of the param
@property {string} [key] Key for the param
@property {string} [value] Value of the param
*/
type Param struct {
	Path  string
	key   string
	value string
}

/**
@info The request structure
@property {*http.Request} [ref] The net/http request instance
@property {string} [key] Key for the param
@property {string} [value] Value of the param
*/
type Request struct {
	ref        *http.Request
	fileReader *multipart.Reader
	body       map[string][]string
	method     string
	Params     []*Param
	query      url.Values
	header     *IncomingHeader
	json       *json.Decoder
	props      *map[string]interface{}
}

func request(httRequest *http.Request, props *map[string]interface{}) *Request {
	req := &Request{}
	req.ref = httRequest
	req.header = &IncomingHeader{}
	req.fileReader = nil
	req.method = httRequest.Proto
	req.props = props
	req.query = httRequest.URL.Query()
	for i, v := range httRequest.Header {
		req.header.Set(strings.ToLower(i), strings.Join(v, ","))
	}
	if req.header.Get("content-type") == "application/json" {
		req.json = json.NewDecoder(httRequest.Body)
	} else {
		httRequest.ParseForm()
	}
	if len(httRequest.PostForm) > 0 && len(req.body) == 0 {
		req.body = make(map[string][]string)
	}
	for key, value := range httRequest.PostForm {
		req.body[key] = value
	}
	return req

}

func (r *Request) GetParam(name string) string {
	var val string
	for _, v := range r.Params {
		if v.Path == r.GetPathURl() && v.key == name {
			val = v.value
		}
	}
	return val
}

func (r *Request) GetPathURl() string {
	return r.ref.URL.Path
}

func (r *Request) Body() map[string][]string {
	return r.body
}

func (r *Request) GetBodyValue(key string) []string {
	return r.body[key]
}

func (r *Request) Header() *IncomingHeader {
	return r.header
}

func (r *Request) Json() *json.Decoder {
	return r.json
}

func (r *Request) Method() string {
	return r.method
}
func (r *Request) Raw() *http.Request {
	return r.ref
}

func (r *Request) GetQuery(key string) string {
	if r.query[key][0] == "" {
		log.Panic("No query param found with given key")
	}
	return r.query[key][0]
}
