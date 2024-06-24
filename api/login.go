package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Login makes a request to the BMC to authenticate the user
func (a *Api) Login(host string, user string, password string) (*Session, error) {
	/*
		curl 'https://endor-bmc.lan/api/session' -X POST -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:128.0) Gecko/20100101 Firefox/128.0' -H 'Accept: application/json, text/javascript, */ /*; q=0.01' -H 'Accept-Language: en-US,en;q=0.5' -H 'Accept-Encoding: gzip, deflate, br, zstd' -H 'Content-Type: application/x-www-form-urlencoded; charset=UTF-8' -H 'X-CSRFTOKEN: null' -H 'X-Requested-With: XMLHttpRequest' -H 'Origin: https://endor-bmc.lan' -H 'Connection: keep-alive' -H 'Referer: https://endor-bmc.lan/' -H 'Cookie: lang=en-us; i18next=en-us' -H 'Sec-Fetch-Dest: empty' -H 'Sec-Fetch-Mode: cors' -H 'Sec-Fetch-Site: same-origin' -H 'Priority: u=0' --data-raw 'username=admin&password=LG2N4700089'
	 */

	// call api
	var domain string
	if a.options.UseSsl {
		domain = "https://" + host
	} else {
		domain = "http://" + host
	}
	resource := "/api/session"
	data := url.Values{}
	data.Set("username", user)
	data.Set("password", password)

	u, _ := url.ParseRequestURI(domain)
	u.Path = resource
	urlStr := u.String()

	r, err := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode()))

	if err != nil {
		return nil, fmt.Errorf("failed to login: %v", err)
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := a.httpClient.Do(r)
	if err != nil {
		return nil, fmt.Errorf("failed to login: %v", err)
	}

	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		// success
		/*
			{
			"ok": 0,
			"privilege": 4,
			"extendedpriv": 259,
			"racsession_id": 3,
			"remote_addr": "10.0.3.4",
			"server_name": "10.0.125.151",
			"server_addr": "10.0.125.151",
			"HTTPSEnabled": 1,
			"CSRFToken": "dkJkp7ir"
			}

			Headers:
			Set-Cookie:QSESSIONID=77a1082bb008e6dd94Bm5koQd85ies; path=/; secure;HttpOnly
		*/

		// parse response into dict
		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			return nil, fmt.Errorf("failed to parse response: %v", err)
		}

		session := &Session{}
		session.User = user
		session.Priveledge = int(response["privilege"].(float64))
		session.SessionId = int(response["racsession_id"].(float64))
		session.RemoteAddr = response["remote_addr"].(string)
		session.ServerAddr = response["server_addr"].(string)
		session.HTTPS = response["HTTPSEnabled"].(float64) == 1
		session.CSRFToken = response["CSRFToken"].(string)
		// parse set-cookie header, extract QSESSIONID
		for _, cookie := range resp.Cookies() {
			if cookie.Name == "QSESSIONID" {
				session.QSESSIONID = cookie.Value
			}
		}

		a.saveSession(host, session)

		return session, nil
	case 401:
		// unauthorized
		return nil, &AuthorizationError{StatusCode: resp.StatusCode, Message: "unauthorized"}
	}

	return nil, &BMCAPIError{StatusCode: resp.StatusCode, Message: "failed to login"}
}
