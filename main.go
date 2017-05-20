package main

import (
	"fmt"
	"os"

	"github.com/gomatic/go-vbuild"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

//
type application struct {
	bytes, lines                         int64
	silent, verbose, null, follow, watch bool
}

//
func (config application) String() string {
	if buf, err := yaml.Marshal(config); err != nil {
		return fmt.Sprintf("%#v", config)
	} else {
		return string(buf)
	}
}

//
var settings application

//
func main() {
	app := cli.NewApp()
	app.Name = "head"
	app.Usage = "Print the first N lines of each FILE to standard output."
	app.Version = build.Version.String()
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		cli.Int64Flag{
			Name:        "c, bytes",
			Usage:       "Print the first NUM bytes of each file.",
			Destination: &settings.bytes,
		},
		cli.Int64Flag{
			Name:        "n, lines",
			Usage:       "Print the first NUM lines of each file.",
			Value:       10,
			Destination: &settings.lines,
		},
		cli.BoolFlag{
			Name:        "q, quiet, silent",
			Usage:       "Just the basics.",
			Destination: &settings.silent,
		},
		cli.BoolFlag{
			Name:        "verbose",
			Usage:       "Print extra stuff.",
			Destination: &settings.verbose,
		},
		cli.BoolFlag{
			Name:        "z, zero-terminated",
			Usage:       "NUL line delimited.",
			Destination: &settings.null,
		},
		cli.BoolFlag{
			Name:        "f, follow",
			Usage:       "Follow.",
			Destination: &settings.follow,
		},
		cli.BoolFlag{
			Name:        "F, watch",
			Usage:       "Follow across new files.",
			Destination: &settings.watch,
		},
	}

	app.Action = run
	app.Run(os.Args)
}

//
func run(ctx *cli.Context) error {
	return nil
}
