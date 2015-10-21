package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/jessevdk/go-flags"
	"github.com/moul/go-sapin"
)

var opts struct {
	Size  int  `short:"s" long:"size" description:"Size of the sapin" default:"5"`
	Balls int  `long:"balls" description:"Percent of balls" default:"10"`
	Color bool `short:"c" long:"color" description:"Colorize output"`
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	if _, err := flags.Parse(&opts); err == nil {
		sapin := sapin.NewSapin(opts.Size)
		sapin.AddBalls(opts.Balls)
		if opts.Color {
			sapin.Colorize()
		}
		fmt.Print(sapin.String())
		os.Exit(0)
	}
	os.Exit(1)
}
