package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type InputData struct {
	Command string
	Args    []string
}

// starts the call
type Console struct {
	inputPort InputPort
	data      InputData
}

func NewConsole(inputPort InputPort, data InputData) *Console {
	return &Console{
		inputPort: inputPort,
		data:      data,
	}
}

func (c Console) Send() {
	fullCommand := append([]string{c.data.Command}, c.data.Args...)
	c.inputPort.Execute(fullCommand)
}

// show the results
type Presenter struct {
	isTable bool
}

func NewPresenter(isTable bool) *Presenter {
	return &Presenter{
		isTable: isTable,
	}
}

func (p Presenter) Present(commandOutput []string) {
	for _, line := range commandOutput {
		if p.isTable {
			fmt.Println("| " + line + " |")
		} else {
			fmt.Println(line)
		}
	}
}

// inner interfaces

type InputPort interface {
	Execute(fullCommand []string)
}

type OutputPort interface {
	Present(commandOutput []string)
}

// inner concrete implementation

type OSExecutor struct {
	outputPort OutputPort
}

func (u OSExecutor) Execute(fullCommand []string) {
	var cmd *exec.Cmd
	if len(fullCommand) == 1 {
		cmd = exec.Command(fullCommand[0])
	} else {
		cmd = exec.Command(fullCommand[0], fullCommand[1:]...)
	}

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	lines := append([]string{}, strings.Split(out.String(), "\n")...)

	u.outputPort.Present(lines)
}

// Main
func main() {
	presenter := NewPresenter(true)

	executor := OSExecutor{
		outputPort: presenter,
	}

	console := NewConsole(
		executor,
		InputData{
			Command: "ls",
			Args:    []string{"-lart"},
		},
	)
	console.Send()
}
