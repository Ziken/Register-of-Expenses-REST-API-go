<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Register of expenses (go)](#register-of-expenses-go)
- [Installation](#installation)
  - [Set Config FIle](#set-config-file)
  - [Run Server](#run-server)
- [Routes](#routes)
  - [POST /expenses](#post-expenses)
    - [Required:](#required)
    - [Response](#response)
  - [GET /expenses](#get-expenses)
    - [Required:](#required-1)
    - [Response](#response-1)
  - [GET /expenses/{id}](#get-expensesid)
    - [Required:](#required-2)
    - [Response](#response-2)
  - [PATH /expenses/{id}](#path-expensesid)
    - [Required:](#required-3)
    - [Response](#response-3)
  - [DELETE /expenses/{id}](#delete-expensesid)
    - [Required](#required)
    - [Response](#response-4)
  - [POST /users/](#post-users)
    - [Requried](#requried)
    - [Response](#response-5)
  - [GET /users/me](#get-usersme)
    - [Required](#required-1)
    - [Response](#response-6)
  - [POST /users/login](#post-userslogin)
    - [Required](#required-2)
    - [Response](#response-7)
  - [POST /users/logout](#post-userslogout)
    - [Required](#required-3)
    - [Response](#response-8)



## Register of expenses (go)
**Register of Expenses** is a simple REST API server for managing your expenses


## Installation
First, you need to download or copy this repository
Then, you need to download unnecessary files by typing in console `` go get ./... `` (in the directory when the server is installed)
### Set Config FIle
The next important step is to create your config file with environment variables.
The file should be created in `/config/`  folder, called `config.json`.
Inside the file you have to set:
- `MONGOHQ_SERVER` - url to your mongodb server
- `MONGOHQ_DATABASE` - name of your database
- `PORT` - in which port it should work
- `JWT_SECRET` - string of characters which is used to create tokens.
Example:
```JSON
{
	"development": {
		"MONGOHQ_SERVER": "mongodb://localhost:27017",
		"MONGOHQ_DATABASE": "ExpensesRegister",

		"PORT": 3000,
		"JWT_SECRET": "dbf9e3070734666156068246f3871cc2"
	}
}
```
You can also set configuration for `test` or `development`.
### Run Server
Type in console `go run server.go`.

## Routes
### POST /expenses
Create new expense document
#### Required:
- `x-auth` in header (token)
- in json format body:
	- `title` - title of the expense
	- `amount`- the amount of money
	- `spentAt`- date as time stamp

Example:
```JSON
{
	"title": "Food",
	"amount": 12.5,
	"spentAt": 1519760066378
}
```

#### Response
The created expense document
Example
```JSON
{
    "result": {
        "_id": "5aa2f11b319cfc4f0872006c",
        "title": "Food",
        "amount": 12.12,
        "spentAt": 123123123123,
        "_creator": "5aa2eb14319cfc49b7c529d4"
    }
}
```


###  GET /expenses
Returns all user's expenses

#### Required:
- `x-auth` in header (token)

#### Response
array of expenses
Example:
```JSON
{
    "result": [
        {
            "_id": "5a95b5215929d31a120f044b",
            "title": "Food",
            "amount": 12.5,
            "spentAt": 1519760066378,
            "_creator": "5a95b4da5929d31a120f0449"
        },
        {
            "_id": "5a95b53b5929d31a120f044c",
            "title": "Book",
            "amount": 50,
            "spentAt": 1519760066378,
            "_creator": "5a95b4da5929d31a120f0449"
        }
    ]
}
```

### GET /expenses/{id}
Get one document expense with the given id.

#### Required:
- `x-auth` in header (token)
- `{id}` of expense in the address

Example:
address: `localhost:3000/expenses/5a95b5215929d31a120f044b`

#### Response
The expense document
Example:
```JSON
{
    "result": {
        "_id": "5a95b5215929d31a120f044b",
        "title": "Food",
        "amount": 12.5,
        "spentAt": 1519760066378,
        "_creator": "5a95b4da5929d31a120f0449"
    }
}
```

### PATH /expenses/{id}

Edit an expense document with the given id.

#### Required:
- `x-auth` in header (token)
- `:id` of expense in the address

Example:
address: `localhost:3000/expenses/5a95b5215929d31a120f044b`
Body:
```JSON
{
	"title": "Pancakes"
}
```
#### Response
Status ``200`` if document was updated.

### DELETE /expenses/{id}
Delete an expense document with the given id
#### Required
- `x-auth` in header (token)
- `{id}` of expense in the address

address: `localhost:3000/expenses/5a95b5215929d31a120f044b`
#### Response
Status ``200`` if document was deleted.

### POST /users/
Create  an user
#### Requried
The body
- `email`
- `password`- at least 6 characters

Example:
```JSON
{
	"email": "email@something.com",
	"password": "123abc"
}
```

#### Response
Created user
Example:
```JSON
{
    "result": {
        "_id": "5a95bc265929d31a120f044d",
        "email": "email@something.com",
        "password: ""
    }
}
```
Also `x-auth` token in header to authenticate user:
`x-auth:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOi...`

### GET /users/me
Authenticate user by its token
#### Required
- `x-auth` in header

#### Response
The user
Example:
```JSON
{
    "result": {
        "_id": "5a95bc265929d31a120f044d",
        "email": "email@something.com",
        "password": ""
    }
}
```

### POST /users/login
Log in user by its credentials (`email` and `password`)

#### Required
- `email` and `password` in the body

Example:
```JSON
{
	"email": "email@something.com",
	"password": "123abc"
}
```
#### Response
The user
Example:
```JSON
{
    "result": {
        "_id": "5a95bc265929d31a120f044d",
        "email": "email@something.com",
        "password": ""
    }
}
```

Also `x-auth` token in header to authenticate user:
`x-auth:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOi...`

### POST /users/logout
Log out the user

#### Required
- `x-auth` token in header

#### Response
Nothing,
```JSON
```
