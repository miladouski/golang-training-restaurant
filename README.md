### Description

Application for storing orders, allows to view, add, update and delete orders

|                Path                   |    Method     |                   Description               |           Body example       |
| --------------------------------------|---------------| --------------------------------------------| ---------------------------- |
| `/orders/{id}`                        |    `DELETE`   |   delete order                              |                              |
| `/orders`                             |    `POST`     |   create new order                          | {"Id":2,"Date":"2021-04-03T00:00:00Z","TableNumber":2,"WaiterId":2,"Price":226,"Payment":false}|
| `/orders/{id}`                        |    `PUT`      |   update order                              |{"Id":2,"Date":"2021-04-03T00:00:00Z","TableNumber":2,"WaiterId":2,"Price":226,"Payment":false}|
| `/orders/{id}`                        |    `GET`      |   get order by `id`                         |                               |
| `/orders`                             |    `GET`      |   get all orders                            |                               |

### Usage

1. Run server on port `8080`

```bash
go run ./cmd/main.go
```

2. Open URL  `http://localhost:8080`

## Usage unit tests
To run unit tests type:
`go test ./...`
