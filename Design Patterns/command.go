package command


/* declares an interface for executing an operation */
type Command interface {
	Execute() string
}


type ToggleOnCommand struct {
	receiver *Receiver
}

func (self *ToggleOnCommand) Execute() string {
	return self.receiver.ToggleOn()
}


type ToggleOffCommand struct {
	receiver *Receiver
}

func (self *ToggleOffCommand) Execute() string {
	return self.receiver.ToggleOff()
}

type Receiver struct {
}


func (self *Receiver) ToggleOn() string {
	return "Toggle On"
}


func (self *Receiver) ToggleOff() string {
	return "Toggle Off"
}

/* ask command to carry out the request */

type Invoker struct {
	commands []Command
}

func (self *Invoker) StoreCommand(command Command) {
	self.commands = append(self.commands, command)
}

func (self *Invoker) UnStoreCommand() {
	if len(self.commands) != 0 {
		self.commands = self.commands[:len(self.commands)-1]
	}
}

func (self *Invoker) Execute() string {
	var result string
	for _, command := range self.commands {
		result += command.Execute() + "\n"
	}
	return result
}
