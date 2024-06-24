package lgr

import (
	"os"

	"github.com/fatih/color"
	"github.com/go-pkgz/lgr"
)

var Logger lgr.L

func SetupLog() {
	logOpts := []lgr.Option{lgr.Format("{{.Message}}")} // default to discard
	dbg := os.Getenv("DEBUG") == "true"
	if dbg {
		logOpts = []lgr.Option{lgr.Debug, lgr.LevelBraces, lgr.StackTraceOnError}
	}
	colorizer := lgr.Mapper{
		ErrorFunc:  func(s string) string { return color.New(color.FgHiRed).Sprint(s) },
		WarnFunc:   func(s string) string { return color.New(color.FgRed).Sprint(s) },
		InfoFunc:   func(s string) string { return color.New(color.FgYellow).Sprint(s) },
		DebugFunc:  func(s string) string { return color.New(color.FgWhite).Sprint(s) },
		CallerFunc: func(s string) string { return color.New(color.FgBlue).Sprint(s) },
		TimeFunc:   func(s string) string { return color.New(color.FgCyan).Sprint(s) },
	}
	logOpts = append(logOpts, lgr.Map(colorizer))
	lgr.SetupStdLogger(logOpts...)
	lgr.Setup(logOpts...)

	Logger = lgr.Std
}
