package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MrZoidberg/megarac/cmd"
	"github.com/MrZoidberg/megarac/config"
	"github.com/MrZoidberg/megarac/lgr"
	"github.com/urfave/cli/v2"
)

var revision = "latest"

func main() {

	lgr.SetupLog()
	err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	app := &cli.App{
		Name:  "megarac",
		Usage: "cli tool for managing MegaRAC BMCs",
		Action: func(cCtx *cli.Context) error {
			return cli.ShowAppHelp(cCtx)
		},
		Commands: []*cli.Command{
			{
				Name:  "configure",
				Usage: "configure BMC connection profiles",
				Subcommands: []*cli.Command{
					{
						Name:   "list",
						Usage:  "list configured BMC profiles",
						Action: cmd.Command(cmd.ProfileList),
					},
					{
						Name:   "add",
						Usage:  "add a new BMC profile",
						Action: cmd.Command(cmd.ProfileAdd),
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "name",
								Usage:    "profile name",
								Aliases:  []string{"n"},
								Required: true,
							},
							&cli.StringFlag{
								Name:     "host",
								Usage:    "BMC hostname or IP address",
								Aliases:  []string{"o"},
								Required: true,
							},
							&cli.StringFlag{
								Name:     "user",
								Usage:    "BMC username",
								Required: true,
								Aliases:  []string{"u"},
							},
							&cli.StringFlag{
								Name:     "password",
								Usage:    "BMC password",
								Required: true,
								Aliases:  []string{"p"},
							},
							&cli.BoolFlag{
								Name:     "use-ssl",
								Usage:    "use ssl for connection",
								Value:    true,
								Required: false,
							},
							&cli.BoolFlag{
								Name:     "insecure",
								Usage:    "skip ssl verification",
								Required: false,
							},
						},
					},
					{
						Name:   "remove",
						Usage:  "remove a BMC profile",
						Action: cmd.Command(cmd.ProfileRemove),
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "name",
								Usage:    "profile name",
								Aliases:  []string{"n"},
								Required: true,
							},
						},
					},
					{
						Name:   "show",
						Usage:  "show a BMC profile",
						Action: cmd.Command(cmd.ProfileShow),
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "name",
								Usage:    "profile name",
								Aliases:  []string{"n"},
								Required: true,
							},
						},
					},
				},
			},
			{
				Name:  "power",
				Usage: "power management",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "profile",
						Usage:    "BMC profile name",
						Aliases:  []string{"n"},
						Required: false,
					},
					&cli.StringFlag{
						Name:     "host",
						Usage:    "BMC hostname or IP address",
						Aliases:  []string{"o"},
						Required: false,
					},
					&cli.StringFlag{
						Name:     "user",
						Usage:    "BMC username",
						Required: false,
						Aliases:  []string{"u"},
					},
					&cli.StringFlag{
						Name:     "password",
						Usage:    "BMC password",
						Required: false,
						Aliases:  []string{"p"},
					},
					&cli.BoolFlag{
						Name:     "use-ssl",
						Usage:    "use ssl for connection",
						Value:    true,
						Required: false,
					},
					&cli.BoolFlag{
						Name:     "insecure",
						Usage:    "skip ssl verification",
						Required: false,
					},
				},
				Subcommands: []*cli.Command{
					{
						Name:   "on",
						Usage:  "power on",
						Action: cmd.Command(cmd.PowerOn),
						Before: func(c *cli.Context) error {
							if c.String("profile") == "" && c.String("host") == "" {
								return cli.Exit("please specify a profile or host", 1)
							}
							return nil
						},
					},
					{
						Name:   "off",
						Usage:  "power off",
						Action: cmd.Command(cmd.PowerOff),
						Before: func(c *cli.Context) error {
							if c.String("profile") == "" && c.String("host") == "" {
								return cli.Exit("please specify a profile or host", 1)
							}
							return nil
						},
					},
					{
						Name:   "status",
						Usage:  "power status",
						Action: cmd.Command(cmd.PowerStatus),
						Before: func(c *cli.Context) error {
							if c.String("profile") == "" && c.String("host") == "" {
								return cli.Exit("please specify a profile or host", 1)
							}
							return nil
						},
					},
				},
			},
			{
				Name:  "sensors",
				Usage: "sensor management",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "profile",
						Usage:    "BMC profile name",
						Aliases:  []string{"n"},
						Required: false,
					},
					&cli.StringFlag{
						Name:     "host",
						Usage:    "BMC hostname or IP address",
						Aliases:  []string{"o"},
						Required: false,
					},
					&cli.StringFlag{
						Name:     "user",
						Usage:    "BMC username",
						Required: false,
						Aliases:  []string{"u"},
					},
					&cli.StringFlag{
						Name:     "password",
						Usage:    "BMC password",
						Required: false,
						Aliases:  []string{"p"},
					},
					&cli.BoolFlag{
						Name:     "use-ssl",
						Usage:    "use ssl for connection",
						Value:    true,
						Required: false,
					},
					&cli.BoolFlag{
						Name:     "insecure",
						Usage:    "skip ssl verification",
						Required: false,
					},
				},
				Subcommands: []*cli.Command{
					{
						Name:   "list",
						Usage:  "list sensors",
						Action: cmd.Command(cmd.SensorList),
						Flags: []cli.Flag{
							&cli.BoolFlag{
								Name:     "all",
								Usage:    "show all sensors, including inactive and inaccessible",
								Required: false,
								Aliases:  []string{"a"},
							},
							&cli.BoolFlag{
								Name:     "find",
								Usage:    "find a sensor by name",
								Required: false,
								Aliases:  []string{"f"},
							},
						},
					},
				},
			},
			{
				Name:  "version",
				Usage: "show version",
				Action: func(c *cli.Context) error {
					fmt.Printf("megarac version %s\n", revision)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
