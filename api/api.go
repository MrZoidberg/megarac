package api

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Session is a struct that holds the session information for the BMC
type Session struct {
	User       string
	Priveledge int
	SessionId  int
	RemoteAddr string
	ServerAddr string
	HTTPS      bool
	CSRFToken  string

	// cookies
	QSessionID string
}

// ApiOptions is a struct that holds the options for the API
type ApiOptions struct {
	InsecureSsl bool // skip ssl verification
	UseSsl      bool // use ssl
}

// Api is the main struct for interacting with the BMC
type Api struct {
	sessions   map[string]*Session
	httpClient *http.Client
	options    *ApiOptions
}

func NewApi(opts ...func(*ApiOptions)) *Api {
	var options = &ApiOptions{}
	for _, opt := range opts {
		opt(options)
	}

	tr := &http.Transport{}
	if options.InsecureSsl {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	return &Api{
		sessions:   make(map[string]*Session),
		httpClient: &http.Client{Transport: tr, Timeout: time.Second * 30},
		options:    options,
	}
}

func (a *Api) saveSession(host string, session *Session) {
	a.sessions[host] = session
}

func (a *Api) getSession(host string) *Session {
	session := a.sessions[host]
	return session
}

func (a *Api) deleteSession(host string) {
	delete(a.sessions, host)
}

func (a *Api) postRequest(host string, path string, data string) (*http.Response, error) {
	// get session
	session := a.getSession(host)
	if session == nil {
		return nil, &NoSessionError{"no session found"}
	}

	// call api
	var domain string
	if a.options.UseSsl {
		domain = "https://" + host
	} else {
		domain = "http://" + host
	}
	resource := "/api/" + path
	urlStr := domain + resource
	apiReq, err := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to construct request: %v", err)
	}
	apiReq.Header.Add("Content-Type", "application/json")
	apiReq.Header.Add("X-CSRFTOKEN", session.CSRFToken)
	apiReq.Header.Add("Cookie", "QSESSIONID="+session.QSessionID)

	resp, err := a.httpClient.Do(apiReq)
	if err != nil {
		return nil, fmt.Errorf("failed to power on: %v", err)
	}

	return resp, nil
}

func (a *Api) getRequest(host string, path string) (*http.Response, error) {
	// get session
	session := a.getSession(host)
	if session == nil {
		return nil, &NoSessionError{"no session found"}
	}

	// call api
	var domain string
	if a.options.UseSsl {
		domain = "https://" + host
	} else {
		domain = "http://" + host
	}
	resource := "/api/" + path
	urlStr := domain + resource
	apiReq, err := http.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to construct request: %v", err)
	}
	apiReq.Header.Add("Content-Type", "application/json")
	apiReq.Header.Add("X-CSRFTOKEN", session.CSRFToken)
	apiReq.Header.Add("Cookie", "QSESSIONID="+session.QSessionID)

	resp, err := a.httpClient.Do(apiReq)
	if err != nil {
		return nil, fmt.Errorf("failed to power on: %v", err)
	}

	return resp, nil
}
