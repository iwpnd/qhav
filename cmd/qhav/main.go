package main

import (
	"fmt"
	"log"
	"os"

	"github.com/iwpnd/qhav"
	"github.com/urfave/cli/v2"
)

func calc(ctx *cli.Context) error {
	var unit = qhav.InKilometers()

	if ctx.String("unit") != "" {
		u := ctx.String("unit")
		switch u {
		case "km":
			unit = qhav.InKilometers()
		case "m":
			unit = qhav.InMeters()
		case "miles":
			unit = qhav.InMiles()
		}
	}

	fmt.Println(qhav.Haversine(ctx.Float64Slice("from"), ctx.Float64Slice("to"), unit))
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
		Usage: "distance unit (m, km, miles).",
	}
}

func main() {
	app := &cli.App{
		Name:   "qhav",
		Usage:  "calculate the haversine distance in a given unit between to two points",
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
