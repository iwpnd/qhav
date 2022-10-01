package main

import (
	"testing"

	"github.com/urfave/cli/v2"
)

func TestCliQhav(t *testing.T) {
	app := cli.NewApp()
	app.Name = "qhav"
	app.Action = calc
	app.Flags = []cli.Flag{
		&unitFlag,
		&fromFlag,
		&toFlag,
	}

	type tcase struct {
		cmd []string
	}

	fn := func(tc tcase) func(t *testing.T) {
		return func(t *testing.T) {
			err := app.Run(tc.cmd)
			if err != nil {
				t.Errorf("qhav should not return an error")
			}
		}
	}

	tests := map[string]tcase{
		"from to": {
			cmd: []string{"qhav", "--from", "0.0,0.0", "--to", "0.5,0.5"},
		},
		"from to with unit km": {
			cmd: []string{"qhav", "--from", "0.0,0.0", "--to", "0.5,0.5", "--unit", "km"},
		},
		"from to with unit m": {
			cmd: []string{"qhav", "--from", "0.0,0.0", "--to", "0.5,0.5", "--unit", "m"},
		},
		"from to with unit miles": {
			cmd: []string{"qhav", "--from", "0.0,0.0", "--to", "0.5,0.5", "--unit", "miles"},
		},
	}

	for name, tc := range tests {
		t.Run(name, fn(tc))
	}

}
