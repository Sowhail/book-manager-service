# book-manager-service
This is a mini project about book management, users can signUp, signIn, create books for themselves and actually do the "CRUD" operations on their own books.
# This project contains
- ### gin  (http web framework in golang)
- ### jwt  (user authentication)
- ### gorm (golang ORM)
- ### A simple layer architecture (DataAccess - Logic - Presentation)
- ### Docker
# How to try it
You need to have docker and docker compose plugin instlled in your system.
then just open a terminal in the bookManager folder of the project and type the commands below.
```go
// terminal
$ docker compose build
$ docker compose up
```
Try all the apis in localhost:3001
 ## Sign Up
 
 ```go
    path: /api/v1/auth/signup
    method: POST

// Request Body
{
	"userName": "yourUserName",
	"emailAddress": "yourEmail@gmail.com",
	"password": "yourPassword",
	"phoneNumber": "yourPhoneNumber", // should contain 11 numbers
	"gender": "yourGender",
	"lastName": "yourLastName",
	"firstName": "yourFirstName"
}
```
## Sign In
```go
    path: /api/v1/auth/login
    method: POST

// Request Body
{
	"userName": "Emily",
	"password": "Soheil@2"
}

// You can see the provided token in your response header
```
## Sign In By Token
```go
    path: /api/v1/auth/autoLogin
    method: POST

// Request Header
Authorization: <token string>
```
## Create book
```go
    path: /api/v1/books
    method: POST

// Request Header
Authorization: <token string>

// Request Body
{
    "name": "book_name",
    "author": {
        "first_name": "fname of author",
        "last_name": "lname of author",
        "birthday": "2000-01-12T00:00:00+03:30",
        "nationality": "french"
    },
    "category": "book_category",
    "volume": 1,
    "published_at": "2000-01-12T00:00:00+03:30",
    "summary": "this is a summary of the book.",
    "table_of_contents": [
        "fasle_1",
        "fasle_2"
    ],
    "publisher": "publisher_name"
}
```
## Get all books
```go
// You will get all the books created by all the users
    path: /api/v1/books
    method: GET

// Request Header
Authorization: <token string>
```
## Get Book By id
```go
    path: /api/v1/books/<bookId> // You can access the book id from database
    method: GET

// Request Header
Authorization: <token string>
```
## Update book info
```go
    path: /api/v1/books/<bookId> // You can access the book id from database
    method: PATCH

// Request Header
Authorization: <token string>

// Request Body
// You can ommit whatever the field that you don't want to change
{
    "name": "book_name",
    "author": {
        "first_name": "fname of author",
        "last_name": "lname of author",
        "birthday": "2000-01-12T00:00:00+03:30",
        "nationality": "french"
    },
    "category": "book_category",
    "volume": 1,
    "published_at": "2000-01-12T00:00:00+03:30",
    "summary": "this is a summary of the book.",
    "table_of_contents": [
        "fasle_1",
        "fasle_2"
    ],
    "publisher": "publisher_name"
}
```
## Delete Book
```go
    path: /api/v1/books/<bookId> // You can access the book id from database
    method: DELETE

// Request Header
Authorization: <token string>
```
