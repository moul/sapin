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
	s.output = strings.Replace(s.output, "~", ansi.Color("~", fmt.Sprintf("yellow+%s", s.ColorOpts)), -1)
}

func (s *Sapin) Emojize() {
	s.compute()

	s.output = strings.Replace(s.output, "@", "🔴", -1)
	s.output = strings.Replace(s.output, "*", "🎄", -1)
	s.output = strings.Replace(s.output, "|", "🚪", -1)
	s.output = strings.Replace(s.output, "#", "💛", -1)
	s.output = strings.Replace(s.output, "~", "🔸", -1)
}

func (s *Sapin) AddStar() {
	s.compute()

	s.output = strings.Replace(s.output, "*", "#", 1)
}

func (s *Sapin) AddPresents() {
	s.compute()

	if s.Size > 3 {
		lines := strings.Split(s.output, "\n")
		lines[len(lines)-4] += "   _8_8_"
		lines[len(lines)-3] += "  |  |  |_8_"
		lines[len(lines)-2] += "  |__|__|___|"
		s.output = strings.Join(lines, "\n")
	}
}

func (s *Sapin) AddGarlands(quantity int) {
	s.compute()

	if s.Size < 2 {
		return
	}

	bodySize := s.GetBodySize()

	for garland := 0; garland < quantity; garland++ {
		lines := strings.Split(s.output, "\n")

		stop := false
		direction := 1
		if rand.Intn(2) > 0 {
			direction = direction * -1
		}
		y := rand.Intn(bodySize-6) + 3
		baseX := strings.Count(lines[y], " ")
		for ; !stop; y += direction {
			line := []byte(lines[y])
			for x := baseX; x < baseX+2; x++ {

				if x >= len(lines[y]) {
					stop = true
					break
				}
				switch lines[y][x] {
				case 32, 124: // ' ', '|'
					stop = true
				}
				if stop {
					break
				}

				line[x] = 126 // ~
			}
			lines[y] = string(line)
			baseX += 2
		}

		s.output = strings.Join(lines, "\n")
	}
}
