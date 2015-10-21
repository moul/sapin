package sapin

// Sapin is the structure defining a sapin.
type Sapin struct {
	Size      int
	output    string
	ColorOpts string
}

// NewSapin returns a Sapin instance of the specified size.
func NewSapin(size int) *Sapin {
	return &Sapin{
		Size: size,
	}
}

// String returns the output of a sapin drawing.
// If the sapin is too small the result will be empty.
func (s *Sapin) String() string {
	s.compute()
	return s.output
}

// putchar appends a char to the output of the sapin.
func (s *Sapin) putchar(char string) {
	s.output += char
}

// GetLineSize returns the size of the sapin for a specified floor and line.
func (s *Sapin) GetLineSize(floor, line int) int {
	return 1 + line*2 + floor*4 + int(floor/2*2)*int((floor+1)/2)
}

// GetMaxSize returns the size of the widest part of the sapin.
func (s *Sapin) GetMaxSize() int {
	return s.GetLineSize(s.Size-1, s.Size+3)
}

func triangularNumber(x int) int {
	if x == 0 {
		return 0
	}
	return x + triangularNumber(x-1)
}

func (s *Sapin) GetBodySize() int {
	return 4*s.Size + triangularNumber(s.Size-1)
}

// compute iterates over floors and lines to generate the output of the sapin.
func (s *Sapin) compute() {
	if s.output != "" {
		return
	}
	// size of the last line of the last floor
	maxSize := s.GetMaxSize()

	// each floor in the floors
	for floor := 0; floor < s.Size; floor++ {

		// each line in the lines of the floor
		for line := 0; line < floor+4; line++ {

			// size of the current line
			lineSize := s.GetLineSize(floor, line)

			// pad left with spaces
			for i := (maxSize-lineSize)/2 - 1; i > 0; i-- {
				s.putchar(" ")
			}

			// draw the body
			for i := 0; i < lineSize; i++ {
				s.putchar("*")
			}

			// new line
			s.putchar("\n")
		}
	}

	// the trunc
	for i := 0; i < s.Size; i++ {
		lineSize := s.Size + (s.Size+1)%2

		// pad left with spaces
		for i := (maxSize-lineSize)/2 - 1; i > 0; i-- {
			s.putchar(" ")
		}

		// draw the body
		for i := 0; i < lineSize; i++ {
			s.putchar("|")
		}

		// new line
		s.putchar("\n")
	}
}
