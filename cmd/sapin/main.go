package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/jessevdk/go-flags"
	"github.com/moul/sapin"
)

var opts struct {
	Size     int  `short:"s" long:"size" description:"Size of the sapin" default:"5"`
	Balls    int  `long:"balls" description:"Percent of balls" default:"4"`
	Garlands int  `long:"garlands" description:"Add some garlands"`
	Color    bool `short:"c" long:"color" description:"Colorize output"`
	Star     bool `long:"star" description:"Add top star"`
	Emoji    bool `short:"e" long:"emoji" description:"Use emojis"`
	Presents bool `long:"presents" description:"Add presents"`
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	if _, err := flags.Parse(&opts); err == nil {
		sapin := sapin.NewSapin(opts.Size)
		if opts.Star {
			sapin.AddStar()
		}
		sapin.AddBalls(opts.Balls)
		sapin.AddGarlands(opts.Garlands)
		if opts.Emoji {
			sapin.Emojize()
		}
		if opts.Color {
			sapin.ColorOpts = "bh"
			sapin.Colorize()
		}
		if opts.Presents {
			sapin.AddPresents()
		}
		fmt.Print(sapin.String())
		os.Exit(0)
	}
	os.Exit(1)
}
