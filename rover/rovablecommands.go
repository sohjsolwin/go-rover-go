package rover

type RovableCommand int

const (
	RoveForward RovableCommand = iota + 1
	RoveReverse
	RoveLeft
	RoveRight
)

func (command RovableCommand) String() string {
	names := [...]string{
		"Forward",
		"Reverse",
		"Left",
		"Right",
	}

	if command < RoveForward || command > RoveRight {
		return "Danger, Will Robinson! (We're confused)"
	}

	return names[command]
}

func (command RovableCommand) IsMotion() bool {
	return command == RoveForward || command == RoveReverse
}

func (command RovableCommand) IsRotation() bool {
	return command == RoveLeft || command == RoveRight
}
