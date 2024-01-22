package digesthttp

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type Session struct {
	httpClient *http.Client
	mutex      sync.Mutex

	auth *authorization
	wa   *wwwAuthenticate

	username string
	password string
}

func New(httpClient *http.Client) *Session {
	return &Session{
		httpClient: httpClient,
	}
}

func (c *Session) SetAuth(username string, password string) {
	c.username = username
	c.password = password
}

func (c *Session) copyRequest(request *http.Request, body []byte) *http.Request {
	newRequest := &http.Request{}
	*newRequest = *request
	newRequest.RequestURI = ""

	newRequest.ContentLength = int64(len(body))
	newRequest.Body = io.NopCloser(bytes.NewReader(body))

	return newRequest
}

func (c *Session) Do(request *http.Request) (*http.Response, error) {
	var dr DigestRequest

	body, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	dr.UpdateRequest(c.username, c.password, request.Method, request.URL.RequestURI(), body)
	dr.Wa = c.wa

	newRequest := c.copyRequest(request, body)

	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.auth != nil {
		auth, err := c.auth.refreshAuthorization(&dr)
		if err != nil {
			return nil, err
		}
		auth.SetTo(request)
	}

	resp, err := c.httpClient.Do(newRequest)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 401 {
		return resp, nil
	}

	waString := resp.Header.Get("WWW-Authenticate")
	if waString == "" {
		return nil, fmt.Errorf("no WWW-Authenticate header")
	}
	wa := newWwwAuthenticate(waString)
	c.wa = wa

	newRequest = c.copyRequest(request, body)
	dr.Wa = c.wa

	auth, err := newAuthorization(&dr)
	if err != nil {
		return nil, err
	}
	c.auth = auth
	auth.SetTo(newRequest)

	return c.httpClient.Do(newRequest)
}
