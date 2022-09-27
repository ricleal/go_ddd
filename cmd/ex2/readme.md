# Clean Architecture Example 2

1. choose how the results are displayed:
   - Implements the `OutputPort` interface: `Present()`
   `presenter := NewPresenter(true)`
2. Choose the logic to used:
   - Implements the `InputPort` interface: `Execute()`
   `logic := BusinessLogic{presenter}`
3. Creates the client with the selected logic and calls the function `Call()` that triggers the execution of the logic
   `client := Client{logic}`
   `client.Call()`
