# qhav - quick haversine

quick haversine distance between to points

## Installation

### cli

```
go install github.com/iwpnd/qhav/cmd/qhav@latest
```

### package

```
go get -u github.com/iwpnd/qhav
```

## Usage

```bash
➜ qhav --from 13.37,52.25 --to 13.37,52.26 --unit m
>> 1111.9508023352598 m

➜ qhav --help
NAME:
   qhav - calculate the haversine distance in a given unit between to two points

USAGE:
   qhav [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --from value [ --from value ]  from point. --from 0.5,0.5
   --help, -h                     show help (default: false)
   --to value [ --to value ]      to point. --to 0.0,0.0
   --unit value                   distance unit (m, km, miles). (default: "km")
```

```go
package main

import (
  "fmt"

  "github.com/iwpnd/qhav"
  )


func main() {
  p1 := []float64{0.0,0.0}
  p2 := []float64{0.5,0.5}

  distance := qhav.Haversine(p1, p2, qhav.InKilometers())

  fmt.Println(distance)
}

>> 78.63
```

## License

MIT

## Acknowledgement

## Maintainer

Benjamin Ramser - [@iwpnd](https://github.com/iwpnd)

Project Link: [https://github.com/iwpnd/qhav](https://github.com/iwpnd/qhav)
