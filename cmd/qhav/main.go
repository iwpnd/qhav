package main

import (
	"fmt"
	"log"
	"os"

	"github.com/iwpnd/qhav"
	"github.com/urfave/cli/v2"
)

func calc(ctx *cli.Context) error {
	unit_option := qhav.InKilometers()
	unit_string := ctx.String("unit")

	switch unit_string {
	case "km":
		unit_option = qhav.InKilometers()
	case "m":
		unit_option = qhav.InMeters()
	case "miles":
		unit_option = qhav.InMiles()
	}

	fmt.Println(qhav.Haversine(ctx.Float64Slice("from"), ctx.Float64Slice("to"), unit_option), unit_string)
	return nil
}

var unitFlag cli.StringFlag
var fromFlag cli.Float64SliceFlag
var toFlag cli.Float64SliceFlag

func init() {
	fromFlag = cli.Float64SliceFlag{
		Name:     "from",
		Usage:    "from point. --from 0.5,0.5",
		Required: true,
	}
	toFlag = cli.Float64SliceFlag{
		Name:     "to",
		Usage:    "to point. --to 0.0,0.0",
		Required: true,
	}

	unitFlag = cli.StringFlag{
		Name:  "unit",
		Value: "km",
		Usage: "distance unit_option (m, km, miles).",
	}
}

func main() {
	app := &cli.App{
		Name:   "qhav",
		Usage:  "calculate the haversine distance in a given unit_option between to two points",
		Action: calc,
		Flags: []cli.Flag{
			&unitFlag,
			&fromFlag,
			&toFlag,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
