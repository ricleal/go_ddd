go-cleanarchitecture
====================

An example Go application demonstrating The Clean Architecture.

[http://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/](http://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/)

- Domain:
  - Customer entity
  - Item entity
  - Order entity
- Use Cases:
  - User entity
  - Use case: Add Item to Order
  - Use case: Get Items in Order
  - Use case: Admin adds Item to Order
- Interfaces:
  - Web Services for Item/Order handling
  - Repositories for Use Cases and Domain entities persistence
- Infrastructure:
  - The Database
  - Code that handles DB connections
  - The HTTP server
  - Go Standard Library

```mermaid
graph TD;
    Infrastructure-->B;
    B-->Use Cases;
    Use Cases-->Domain;
```

Install
-------

Create the SQLite structure

    sqlite3 /tmp/db.sqlite < setup.sql

Run the server

    make build && make run

Access the web endpoint at [http://localhost:8080/orders?userId=40&orderId=60](http://localhost:8080/orders?userId=40&orderId=60)

To run the tests, for each module, run

    make test

Enjoy.


License
-------
The MIT License (MIT)

Copyright (c) 2012 Manuel Kiessling

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
