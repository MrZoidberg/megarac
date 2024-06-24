package cmd

import (
	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/MrZoidberg/megarac/api"
	"github.com/MrZoidberg/megarac/lgr"
	"github.com/urfave/cli/v2"
)

// SensorList is a function that lists sensors using the BMC API
// It does the login to the BMC or gets the session if it's available,
// and then calls the sensor list API
func SensorList(c *cli.Context) error {
	profile, err := getProfile(c)
	if err != nil {
		return err
	}
	showAll := false
	if c.IsSet("all") {
		showAll = c.Bool("all")
	}
	find := false
	find_str := ""
	if c.IsSet("find") {
		find = c.Bool("find")
		find_str = c.Args().First()
	}

	srv := api.NewApi(func(ao *api.ApiOptions) {
		ao.InsecureSsl = *profile.InsureSsl
		ao.UseSsl = *profile.UseSSL
	})

	_, err = srv.Login(profile.Host, profile.User, profile.Password)
	if err != nil {
		return cli.Exit(fmt.Sprintf("FAIL: Failed to login to BMC host %s: %v", profile.Host, err), 1)
	}

	sensors, err := srv.GetSensorsList(profile.Host)
	if err != nil {
		return cli.Exit(fmt.Sprintf("FAIL: Failed to get sensor list from BMC host %s: %v", profile.Host, err), 1)
	}

	var output strings.Builder
	w := tabwriter.NewWriter(&output, 0, 0, 1, ' ', 0)
	fmt.Fprintln(w, "ID\tName\tType\tReading\tAlert\tState")
	for _, sensor := range sensors {
		if !showAll && sensor.State == "inactive" || sensor.Accessible == "inaccessible" {
			continue
		}
		if find && !strings.Contains(sensor.Name, find_str) {
			continue
		}

		fmt.Fprintf(w, "%d\t%s\t%s\t%s %s\t%s\t%s\n", sensor.ID, sensor.Name, sensor.Type,
			sensor.Reading, sensor.Unit, sensor.Alert, sensor.State)
	}
	w.Flush()

	lgr.Logger.Logf("%s", output.String())

	return nil
}