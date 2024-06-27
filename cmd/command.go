package cmd

import (
	"fmt"

	"github.com/MrZoidberg/megarac/config"
	"github.com/urfave/cli/v2"
)

func Command(fn func(*cli.Context) error) func(*cli.Context) error {

	wrapper := func(c *cli.Context) error {
		return fn(c)
	}

	return wrapper
}

// getProfile returns a profile from the configuration or from the command line flags
func getProfile(c *cli.Context) (*config.Profile, error) {
	profile := &config.Profile{}

	if c.IsSet("profile") {
		profile = config.Cfg.GetProfile(c.String("profile"))
		if profile == nil {
			return nil, cli.Exit(fmt.Sprintf("Profile %s not found", c.String("profile")), 1)
		}
	} else {
		profile.Name = "default"
		profile.Host = c.String("host")
		profile.User = c.String("user")
		profile.Password = c.String("password")
	}

	if c.IsSet("insecure") {
		val := c.Bool("insecure")
		profile.InsureSsl = &val
	} else if profile.InsureSsl == nil {
		val := false
		profile.InsureSsl = &val
	}

	if c.IsSet("use-ssl") {
		val := c.Bool("use-ssl")
		profile.UseSSL = &val
	} else if profile.UseSSL == nil {
		val := true
		profile.UseSSL = &val
	}

	return profile, nil
}

type OutputFormat string

const (
	OutputFormatText OutputFormat = "text"
	OutputFormatJSON OutputFormat = "json"
)

func getOutputFormat(c *cli.Context) OutputFormat {
	if c.IsSet("format") {
		return OutputFormat(c.String("format"))
	}
	return OutputFormatText
}
