package api

import (
	"encoding/json"
	"fmt"
)

// PowerOn is a function that powers on server using the BMC API
func (a *Api) PowerOn(host string) error {
	/*
		curl 'https://endor-bmc.lan/api/actions/power' -X POST -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:128.0) Gecko/20100101 Firefox/128.0' -H 'Accept: application/json, text/javascript, */ /*; q=0.01' -H 'Accept-Language: en-US,en;q=0.5' -H 'Accept-Encoding: gzip, deflate, br, zstd' -H 'Content-Type: application/json' -H 'X-CSRFTOKEN: Nwu17Ief' -H 'X-Requested-With: XMLHttpRequest' -H 'Origin: https://endor-bmc.lan' -H 'Connection: keep-alive' -H 'Referer: https://endor-bmc.lan/' -H 'Cookie: lang=en-us; i18next=en-us; QSESSIONID=298108048508e6d6eaeA1kdcvw8Jux; refresh_disable=1' -H 'Sec-Fetch-Dest: empty' -H 'Sec-Fetch-Mode: cors' -H 'Sec-Fetch-Site: same-origin' -H 'Priority: u=0' --data-raw '{"power_command":1}'
	 */

	data := `{"power_command":1}`

	resp, err := a.postRequest(host, "actions/power", data)
	if err != nil {
		return fmt.Errorf("failed to power on: %v", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return &BMCAPIError{StatusCode: resp.StatusCode, Message: "failed to power on"}
	}

	return nil
}

// PowerOn is a function that powers on server using the BMC API
func (a *Api) PowerOff(host string) error {
	/*
		curl 'https://endor-bmc.lan/api/actions/power' -X POST -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:128.0) Gecko/20100101 Firefox/128.0' -H 'Accept: application/json, text/javascript, */ /*; q=0.01' -H 'Accept-Language: en-US,en;q=0.5' -H 'Accept-Encoding: gzip, deflate, br, zstd' -H 'Content-Type: application/json' -H 'X-CSRFTOKEN: Nwu17Ief' -H 'X-Requested-With: XMLHttpRequest' -H 'Origin: https://endor-bmc.lan' -H 'Connection: keep-alive' -H 'Referer: https://endor-bmc.lan/' -H 'Cookie: lang=en-us; i18next=en-us; QSESSIONID=298108048508e6d6eaeA1kdcvw8Jux; refresh_disable=1' -H 'Sec-Fetch-Dest: empty' -H 'Sec-Fetch-Mode: cors' -H 'Sec-Fetch-Site: same-origin' -H 'Priority: u=0' --data-raw '{"power_command":0}'
	 */

	data := `{"power_command":0}`

	resp, err := a.postRequest(host, "actions/power", data)
	if err != nil {
		return fmt.Errorf("failed to power off: %v", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return &BMCAPIError{StatusCode: resp.StatusCode, Message: "failed to power off"}
	}

	return nil
}

type ChassisStatus struct {
	PowerStatus int `json:"power_status"`
	LEDStatus   int `json:"led_status"`
}

// ChassisStatus is a function that gets the power status of the server using the BMC API
func (a *Api) ChassisStatus(host string) (*ChassisStatus, error) {
	/*
		curl 'https://endor-bmc.lan/api/chassis-status' -X GET -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:128.0) Gecko/20100101 Firefox/128.0' -H 'Accept: application/json, text/javascript, */ /*; q=0.01' -H 'Accept-Language: en-US,en;q=0.5' -H 'Accept-Encoding: gzip, deflate, br, zstd' -H 'X-Requested-With: XMLHttpRequest' -H 'Connection: keep-alive' -H 'Referer: https://endor-bmc.lan/' -H 'Cookie: lang=en-us; i18next=en-us; QSESSIONID=298108048508e6d6eaeA1kdcvw8Jux; refresh_disable=1' -H 'Sec-Fetch-Dest: empty' -H 'Sec-Fetch-Mode: cors' -H 'Sec-Fetch-Site: same-origin' -H 'Priority: u=0'
	 */
	// response: { "power_status": 0, "led_status": 0 }

	resp, err := a.getRequest(host, "chassis-status")
	if err != nil {
		return nil, fmt.Errorf("failed to get power status: %v", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, &BMCAPIError{StatusCode: resp.StatusCode, Message: "failed to get power status"}
	}

	var status ChassisStatus
	err = json.NewDecoder(resp.Body).Decode(&status)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	return &status, nil
}
