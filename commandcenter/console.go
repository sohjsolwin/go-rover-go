package commandcenter

import "github.com/sohjsolwin/go-rover-go/rover"

type CommunicationSatalite interface {
	ConnectToRovable(r *rover.Rovable) CommunicationPipe
	SendRovableCommand(command rover.RovableCommand)
}

type CommunicationPipe struct {
	ReceiveCh <-chan string
	SendCh    chan<- rover.RovableCommand
}

// SendCommand transmits a command to a connected rovable through a CommunicationPipe
func (c *CommunicationPipe) SendRovableCommand(command rover.RovableCommand) {
	c.SendCh <- command
}

func (c *CommunicationPipe) ReceiveRovableTransmission(command rover.RovableCommand) string {
	return <-c.ReceiveCh
}
