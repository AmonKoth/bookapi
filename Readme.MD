Book API

A simple RESTful API in Go for managing books, built with only the standard library. Data is stored in memory and is lost when the server stops.

The server listens on http://localhost:8080.


| Method | Path          | Description       |
|--------|---------------|-------------------|
| GET    | `/books`      | List all books    |
| GET    | `/books/{id}` | Get a book        |
| POST   | `/books`      | Create a book     |
| PUT    | `/books/{id}` | Update a book     |
| DELETE | `/books/{id}` | Delete a book     |

`title` and `author` are required. `id` is assigned by the server.

## Examples

```
curl localhost:8080/books
curl -X POST localhost:8080/books -d '{"title":"Go","author":"Pike"}'
curl localhost:8080/books/1
curl -X PUT localhost:8080/books/1 -d '{"title":"Go 2","author":"Pike"}'
curl -X DELETE localhost:8080/books/1
```