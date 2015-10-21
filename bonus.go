package sapin

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/mgutz/ansi"
)

func (s *Sapin) AddBalls(percent int) {
	s.compute()

	slc := []byte(s.output)
	for idx, i := range slc {
		switch i {
		case 42: // *
			if rand.Intn(100) < percent {
				slc[idx] = 64 // @
			}
		}
	}
	s.output = string(slc)
}

func (s *Sapin) Colorize() {
	s.compute()

	s.output = strings.Replace(s.output, "@", ansi.Color("@", fmt.Sprintf("red+%s", s.ColorOpts)), -1)
	s.output = strings.Replace(s.output, "*", ansi.Color("*", fmt.Sprintf("green+%s", s.ColorOpts)), -1)
	s.output = strings.Replace(s.output, "|", ansi.Color("|", fmt.Sprintf("90+%s", s.ColorOpts)), -1)
	s.output = strings.Replace(s.output, "#", ansi.Color("#", fmt.Sprintf("yellow+%s", s.ColorOpts)), -1)
}

func (s *Sapin) Emojize() {
	s.compute()

	s.output = strings.Replace(s.output, "@", "🔴", -1)
	s.output = strings.Replace(s.output, "*", "🎄", -1)
	s.output = strings.Replace(s.output, "|", "🚪", -1)
	s.output = strings.Replace(s.output, "#", "💛", -1)
}

func (s *Sapin) AddStar() {
	s.compute()

	s.output = strings.Replace(s.output, "*", "#", 1)
}
