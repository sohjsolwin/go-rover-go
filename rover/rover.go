package rover

type point struct{ X, Y int }

type targetDirection *int

// Rover represents a single rovable object on the surface of Mars
type Rover struct {
	facing   targetDirection
	flank    targetDirection
	currentX *int
	currentY *int
}

// ProcessCommand will take the supplied command and either move or rotate the Rover
func (r *Rover) ProcessCommand(command *RovableCommand) {
	if command.IsMotion() {
		r.move(command)
	} else if command.IsRotation() {
		r.rotate(command)
	} else {
		panic("Unknown Command! Ai yai yai!!!")
	}
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
