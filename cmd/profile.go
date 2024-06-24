package cmd

import (
	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/MrZoidberg/megarac/config"
	"github.com/MrZoidberg/megarac/lgr"
	"github.com/urfave/cli/v2"
)

// ProfileAdd is a function that adds a new BMC profile to the configuration
func ProfileAdd(c *cli.Context) error {
	name := c.String("name")
	host := c.String("host")
	user := c.String("user")
	password := c.String("password")

	cfg := config.Cfg
	profile := config.Profile{
		Name:     name,
		Host:     host,
		User:     user,
		Password: password,
	}
	if c.IsSet("use-ssl") {
		val := c.Bool("use-ssl")
		profile.UseSSL = &val
	}
	if c.IsSet("insecure") {
		val := c.Bool("insecure")
		profile.InsureSsl = &val
	}
	cfg.AddProfile(&profile)

	err := cfg.Save()
	if err != nil {
		return cli.Exit(fmt.Sprintf("Failed to save profile: %v", err), 1)
	}

	lgr.Logger.Logf("[INFO] Profile %s was added", name)

	return nil
}

// ProfileRemove is a function that removes a BMC profile from the configuration
func ProfileRemove(c *cli.Context) error {
	name := c.String("name")

	cfg := config.Cfg
	cfg.RemoveProfile(name)

	err := cfg.Save()
	if err != nil {
		return cli.Exit(fmt.Sprintf("Failed to save profile: %v", err), 1)
	}

	lgr.Logger.Logf("[INFO] Profile %s was removed", name)

	return nil
}

// ProfileList is a function that lists all BMC profiles from the configuration
func ProfileList(c *cli.Context) error {
	cfg := config.Cfg

	var output strings.Builder
	w := tabwriter.NewWriter(&output, 0, 0, 2, ' ', 0)
	fmt.Fprintf(w, "Profile Name\tHost Name\n")
	fmt.Fprintf(w, "------------\t---------\n")
	for _, p := range cfg.Profiles {
		fmt.Fprintf(w, "%s\t%s\n", p.Name, p.Host)
	}
	w.Flush()

	lgr.Logger.Logf("%s", output.String())

	return nil
}

// ProfileShow is a function that shows a BMC profile from the configuration
func ProfileShow(c *cli.Context) error {
	name := c.String("name")

	cfg := config.Cfg
	profile := cfg.GetProfile(name)
	if profile == nil {
		return cli.Exit(fmt.Sprintf("Profile %s not found", name), 1)
	}

	lgr.Logger.Logf("Profile Name: %s", profile.Name)
	lgr.Logger.Logf("Host Name: %s", profile.Host)
	lgr.Logger.Logf("User: %s", profile.User)
	if profile.UseSSL != nil {
		lgr.Logger.Logf("Use SSL: %t", *profile.UseSSL)
	}
	if profile.InsureSsl != nil {
		lgr.Logger.Logf("Insecure SSL: %t", *profile.InsureSsl)
	}

	return nil
}
