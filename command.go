package main
import "fmt"
type Command interface {
	Execute()
}
type Receiver struct {
	Name string
}
func (r *Receiver) Action() {
	fmt.Printf("%s is performing an action.\n", r.Name)
}
type ConcreteCommand struct {
	receiver *Receiver
}

func NewConcreteCommand(receiver *Receiver) *ConcreteCommand {
	return &ConcreteCommand{
		receiver: receiver,
	}
}

func (cc *ConcreteCommand) Execute() {
	cc.receiver.Action()
}
type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}

func main() {
	receiver := &Receiver{Name: "Receiver 1"}

	command := NewConcreteCommand(receiver)

	invoker := &Invoker{}
	invoker.SetCommand(command)
	invoker.ExecuteCommand()
}
