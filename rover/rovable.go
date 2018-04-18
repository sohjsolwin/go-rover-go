package rover

// Rovable is an object designed to move and rotate based on RovableCommands
type Rovable interface {
	ProcessCommand(command *RovableCommand)
}
