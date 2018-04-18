package rover

import (
	"fmt"
)

type point struct{ X, Y int }

type targetDirection *int

// Rover represents a single rovable object on the surface of Mars
type Rover struct {
	facing   targetDirection
	flank    targetDirection
	currentX *int
	currentY *int
	id       string
	outputCh chan<- string
	inputCh  <-chan RovableCommand
}

// ProcessCommand will take the supplied command and
// either move or rotate the Rover, returning the
// resulting position of the rover
func (r *Rover) ProcessCommand(command *RovableCommand) (x, y int) {
	if command.IsMotion() {
		r.move(command)
	} else if command.IsRotation() {
		r.rotate(command)
	} else if *command == Status {
		r.reportStatus()
	} else {
		panic("Unknown Command! Ai yai yai!!!")
	}

	return *r.currentX, *r.currentY
}

func (r *Rover) move(command *RovableCommand) {
	if *command == RoveForward {
		*r.facing++
	} else {
		*r.facing--
	}
}

func (r *Rover) rotate(command *RovableCommand) {
	r.facing, r.flank = r.flank, r.facing
}

func (r *Rover) reportStatus() {
	r.outputCh <- r.String()
}

func (r *Rover) String() string {
	return fmt.Sprintf("Rover %s believes it is at: [x] %d, [y] %d", r.id, r.currentX, r.currentY)
}
