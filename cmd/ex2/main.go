package main

import "fmt"

// Reference: https://user-images.githubusercontent.com/20016700/47955442-f0a66c00-dfa0-11e8-83f9-2856ed6ba159.png

// outer

// starts the call
type Client struct {
	useCaseInputPort UseCaseInputPort
}

func (c Client) Call() {
	fmt.Println("Client.Call() - Get the request")
	c.useCaseInputPort.Execute()
	fmt.Println("Client.Call() - Send the response")
}

// show the results
type Presenter struct {
	fancy bool
}

func NewPresenter(fancy bool) Presenter {
	return Presenter{fancy}
}

func (p Presenter) Present() {
	if p.fancy {
		fmt.Println("    Presenter.Present() - Show the FANCY results")
	} else {
		fmt.Println("    Presenter.Present() - Show the RAW results")
	}
}

// inner interfaces

type UseCaseInputPort interface {
	Execute()
}

type UseCaseOutputPort interface {
	Present()
}

// inner concrete implementation

type BusinessLogic struct {
	useCaseOutputPort UseCaseOutputPort
}

func (u BusinessLogic) Execute() {
	fmt.Println("  BusinessLogic.Execute() - Business Logic")
	u.useCaseOutputPort.Present()
	fmt.Println("  BusinessLogic.Execute() - Business Logic cleanup")
}

func main() {
	presenter := NewPresenter(true)
	logic := BusinessLogic{presenter}
	client := Client{logic}
	client.Call()
}
