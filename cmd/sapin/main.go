package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/moul/go-sapin"
)

var opts struct {
	Size int `short:"s" long:"size" description:"Size of the sapin" required:"true"`
}

func main() {
	if _, err := flags.Parse(&opts); err == nil {
		sapin := sapin.NewSapin(opts.Size)
		fmt.Print(sapin.String())
		os.Exit(0)
	}
	os.Exit(1)
}
