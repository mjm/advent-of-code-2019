package day17

// Cleaner provides movement programs to control the robot.
type Cleaner struct {
	Main      string
	MovementA string
	MovementB string
	MovementC string
}

// NewCleaner creates a new cleaner with the correct programs.
func NewCleaner() *Cleaner {
	return &Cleaner{
		Main:      "A,B,A,B,C,C,B,A,B,C",
		MovementA: "L,4,R,8,L,6,L,10",
		MovementB: "L,6,R,8,R,10,L,6,L,6",
		MovementC: "L,4,L,4,L,10",
	}
}

// Run sends the programs to the input channel, and reads the map and dust collected
// from the output. It returns the dust collected.
func (c *Cleaner) Run(input chan<- int, output <-chan int) int {
	go c.sendPrograms(input)

	var lastOutput int
	for out := range output {
		lastOutput = out
	}
	return lastOutput
}

func (c *Cleaner) sendPrograms(input chan<- int) {
	c.sendProgram(input, c.Main)
	c.sendProgram(input, c.MovementA)
	c.sendProgram(input, c.MovementB)
	c.sendProgram(input, c.MovementC)
	input <- int('n')
	input <- 10
}

func (c *Cleaner) sendProgram(input chan<- int, program string) {
	for _, r := range program {
		input <- int(r)
	}
	input <- 10 // newline
}
