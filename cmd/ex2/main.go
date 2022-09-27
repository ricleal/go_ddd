package main

import "fmt"

// Reference: https://user-images.githubusercontent.com/20016700/47955442-f0a66c00-dfa0-11e8-83f9-2856ed6ba159.png

// outer

// starts the call
type Client struct {
	inputPort InputPort
	post      bool
}

func NewClient(inputPort InputPort, post bool) *Client {
	return &Client{inputPort, post}
}

func (c Client) Call() {
	if c.post {
		fmt.Println("Client.Call() - POST the request")
	} else {
		fmt.Println("Client.Call() - GET the request")
	}
	c.inputPort.Execute()
	fmt.Println("Client.Call() - Send the response")
}

// show the results
type Presenter struct {
	fancy bool
}

func NewPresenter(fancy bool) *Presenter {
	return &Presenter{fancy}
}

func (p Presenter) Present() {
	if p.fancy {
		fmt.Println("    Presenter.Present() - Show the FANCY results")
	} else {
		fmt.Println("    Presenter.Present() - Show the RAW results")
	}
}

// inner interfaces

type InputPort interface {
	Execute()
}

type OutputPort interface {
	Present()
}

// inner concrete implementation

type BusinessLogic1 struct {
	outputPort OutputPort
}

func (u BusinessLogic1) Execute() {
	fmt.Println("  BusinessLogic1.Execute() - Business Logic")
	u.outputPort.Present()
	fmt.Println("  BusinessLogic1.Execute() - Business Logic cleanup")
}

type BusinessLogic2 struct {
	outputPort OutputPort
}

func (u BusinessLogic2) Execute() {
	fmt.Println("  BusinessLogic2.Execute() - Business Logic")
	u.outputPort.Present()
	fmt.Println("  BusinessLogic2.Execute() - Business Logic cleanup")
}

// Main
func main() {
	presenter1 := NewPresenter(true)
	logic1 := BusinessLogic1{presenter1}
	client := NewClient(logic1, true)
	client.Call()

	presenter2 := NewPresenter(false)
	logic2 := BusinessLogic2{presenter2}
	client = NewClient(logic2, false)
	client.Call()
}
