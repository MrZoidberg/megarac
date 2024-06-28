package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/MrZoidberg/megarac/api"
	"github.com/MrZoidberg/megarac/lgr"
	"github.com/urfave/cli/v2"
)

// PowerOn is a function that powers on server using the BMC API
// It does the login to the BMC or gets the session if it's available,
// and then calls the power on API
func PowerOn(c *cli.Context) error {
	profile, err := getProfile(c)
	if err != nil {
		return err
	}
	format := getOutputFormat(c)

	srv := api.NewApi(func(ao *api.ApiOptions) {
		ao.InsecureSsl = *profile.InsureSsl
		ao.UseSsl = *profile.UseSSL
	})

	_, err = srv.Login(profile.Host, profile.User, profile.Password)
	if err != nil {
		return cli.Exit(fmt.Sprintf("FAIL: Failed to login to BMC host %s: %v", profile.Host, err), 1)
	}

	defer func() {
		err = srv.Logout(profile.Host)
		if err != nil {
			lgr.Logger.Logf("[WARN] Failed to logout from BMC host %s: %v", profile.Host, err)
		}
	}()

	err = srv.PowerOn(profile.Host)
	if err != nil {
		return cli.Exit(fmt.Sprintf("FAIL: Failed to power on BMC host %s: %v", profile.Host, err), 1)
	}

	if format == OutputFormatText {
		lgr.Logger.Logf("[INFO] Server %s is powering on", profile.Host)
	} else {
		result := map[string]interface{}{
			"host": profile.Host,
			"power": map[string]interface{}{
				"status": "on",
				"msg":    "Server is powering on",
			},
		}
		output, err := json.Marshal(result)
		if err != nil {
			return cli.Exit(fmt.Sprintf("FAIL: Failed to marshal output: %v", err), 1)
		}
		lgr.Logger.Logf("%v", string(output))
	}

	return nil
}

// PowerOn is a function that powers off server using the BMC API
// It does the login to the BMC or gets the session if it's available,
// and then calls the power on API
func PowerOff(c *cli.Context) error {
	profile, err := getProfile(c)
	if err != nil {
		return err
	}
	format := getOutputFormat(c)

	srv := api.NewApi(func(ao *api.ApiOptions) {
		ao.InsecureSsl = *profile.InsureSsl
		ao.UseSsl = *profile.UseSSL
	})

	_, err = srv.Login(profile.Host, profile.User, profile.Password)
	if err != nil {
		return cli.Exit(fmt.Sprintf("FAIL: Failed to login to BMC host %s: %v", profile.Host, err), 1)
	}

	defer func() {
		err = srv.Logout(profile.Host)
		if err != nil {
			lgr.Logger.Logf("[WARN] Failed to logout from BMC host %s: %v", profile.Host, err)
		}
	}()

	err = srv.PowerOff(profile.Host)
	if err != nil {
		return cli.Exit(fmt.Sprintf("FAIL: Failed to power off BMC host %s: %v", profile.Host, err), 1)
	}

	if format == OutputFormatText {
		lgr.Logger.Logf("[INFO] Server %s is powering off", profile.Host)
	} else {
		result := map[string]interface{}{
			"host": profile.Host,
			"power": map[string]interface{}{
				"status": "off",
				"msg":    "Server is powering down",
			},
		}
		output, err := json.Marshal(result)
		if err != nil {
			return cli.Exit(fmt.Sprintf("FAIL: Failed to marshal output: %v", err), 1)
		}
		lgr.Logger.Logf("%v", string(output))
	}

	return nil
}

// PowerStatus is a function that gets the power status of the server using the BMC API
// It does the login to the BMC or gets the session if it's available,
// and then calls the power status API
func PowerStatus(c *cli.Context) error {
	profile, err := getProfile(c)
	if err != nil {
		return err
	}
	format := getOutputFormat(c)

	srv := api.NewApi(func(ao *api.ApiOptions) {
		ao.InsecureSsl = *profile.InsureSsl
		ao.UseSsl = *profile.UseSSL
	})

	_, err = srv.Login(profile.Host, profile.User, profile.Password)
	if err != nil {
		return cli.Exit(fmt.Sprintf("FAIL: Failed to login to BMC host %s: %v", profile.Host, err), 1)
	}

	defer func() {
		err = srv.Logout(profile.Host)
		if err != nil {
			lgr.Logger.Logf("[WARN] Failed to logout from BMC host %s: %v", profile.Host, err)
		}
	}()

	status, err := srv.ChassisStatus(profile.Host)
	if err != nil {
		return cli.Exit(fmt.Sprintf("FAIL: Failed to get power status for BMC host %s: %v", profile.Host, err), 1)
	}
	var powerStatus string
	if status.PowerStatus == 0 {
		powerStatus = "off"
	} else {
		powerStatus = "on"
	}

	if format == OutputFormatText {
		lgr.Logger.Logf("[INFO] Power status for %s: %v", profile.Host, powerStatus)
	} else {
		result := map[string]interface{}{
			"host": profile.Host,
			"power": map[string]interface{}{
				"status": powerStatus,
				"msg":    fmt.Sprintf("Server is %s", powerStatus),
			},
		}
		output, err := json.Marshal(result)
		if err != nil {
			return cli.Exit(fmt.Sprintf("FAIL: Failed to marshal output: %v", err), 1)
		}
		lgr.Logger.Logf("%v", string(output))
	}

	return nil
}
