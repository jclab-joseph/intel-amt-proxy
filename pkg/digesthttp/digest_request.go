package digesthttp

import (
	"net/http"
)

type DigestRequest struct {
	Body     []byte
	Method   string
	Password string
	URI      string
	Username string
	Header   http.Header
	Auth     *authorization
	Wa       *wwwAuthenticate
	CertVal  bool
}

// UpdateRequest is called when you want to reuse an existing
//
//	DigestRequest connection with new request information
func (dr *DigestRequest) UpdateRequest(username, password, method, uri string, body []byte) *DigestRequest {
	dr.Body = body
	dr.Method = method
	dr.Password = password
	dr.URI = uri
	dr.Username = username
	dr.Header = make(map[string][]string)
	return dr
}
