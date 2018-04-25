package satellite

import "github.com/sohjsolwin/go-rover-go/rover"

// ISpaceDialer connects to a rovable object and can send commands to that object.
type ISpaceDialer interface {
	ConnectToRovable(r *rover.Rovable) CommunicationPipe
	SendRovableCommand(command rover.RovableCommand)
}

// CommunicationPipe represents a direct communication connection with a single rovable object.
type CommunicationPipe struct {
	ReceiveCh <-chan string
	SendCh    chan<- rover.RovableCommand
}

// SendRovableCommand transmits a command to a connected rovable through a CommunicationPipe
func (c *CommunicationPipe) SendRovableCommand(command rover.RovableCommand) {
	c.SendCh <- command
}

// ReceiveRovableTransmission receives the results of the command submitted to the rover.
func (c *CommunicationPipe) ReceiveRovableTransmission(command rover.RovableCommand) string {
	return <-c.ReceiveCh
}
