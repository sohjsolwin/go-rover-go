package rover

// RovableCommand represents a command to be executed by a rovable object.
type RovableCommand int

// Values defined to represent executable commands for a rovable object
const (
	RoveForward RovableCommand = iota + 1
	RoveReverse
	RoveLeft
	RoveRight
	Status
)

func (command RovableCommand) String() string {
	names := [...]string{
		"Forward",
		"Reverse",
		"Left",
		"Right",
		"Status",
	}

	if command < RoveForward || command > Status {
		return "Danger, Will Robinson! (We're confused)"
	}

	return names[command]
}

// IsMotion determines if the specified command is one that initiates motion (changing an X or Y value)
func (command RovableCommand) IsMotion() bool {
	return command == RoveForward || command == RoveReverse
}

// IsRotation determines if the specified command is one that changes orientation (turning left or right)
func (command RovableCommand) IsRotation() bool {
	return command == RoveLeft || command == RoveRight
}
