# Clean Architecture Example 1

### Controller and Presenter are both in a outer layer

```mermaid
classDiagram
class UseCaseInputPort {
  <<interface>>
  Execute()
}
class UseCaseOutputPort {
  <<interface>> 
  PresentHTML()
  PresentTXT()
}
class Presenter {
  PresentHTML()
  PresentTXT()
}
class Controller {
  Call()
}
class UseCaseInteractor {
  Execute()
}
Controller --* UseCaseInputPort: useCaseInputPort
UseCaseInputPort <|-- UseCaseInteractor
UseCaseInteractor --* UseCaseOutputPort: useCaseOutputPort
UseCaseOutputPort <|-- Presenter
```

### Diagram

![](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)