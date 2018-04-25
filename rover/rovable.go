package rover

// IRovable is an object designed to move and rotate based on RovableCommands
type IRovable interface {
	ProcessCommand(command *RovableCommand)
}
