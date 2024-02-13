## TransactionRoutineAPI

API responsible for recording operations performed in a transaction

### 📋 Prerequisites

Tools: 

* [Golang](https://golang.org/doc/install)
* [Docker](https://www.docker.com/)


### 🛠️ Built with

* [echo](https://echo.labstack.com/) - Framework Web
* [go mod](https://blog.golang.org/using-go-modules) - Dependency
* [logrus](github.com/sirupsen/logrus) - Log
* [sqlx](https://github.com/jmoiron/sqlx) - Connection management of relational banks
* [validator](github.com/go-playground/validator/v10) - Struct Validator
* [mockgen](https://github.com/uber-go/mock) - Mock for testing
* [decimal](github.com/shopspring/decimal) - Work with monetary values 

### Local environment configuration

### Create the following structure in your root directory:

```
 root
 │  ├── go
 │      ├── src
 │          ├── github.com
 │              ├── pismo

```

### Project download

* Clone repository inside folder `pismo`


### ⚙️ Running the tests

* `make test`: runs the tests.


### 🚗 Running

* `make run`: default command to run the program.

- OBS.: The Project will run at the door `:8080`


### 🗂 Architecture

### Description of the most important directories and files:

- `./api/v1`: This directory has the configuration and registration of all sub-modules.
- `./api/v1/v1.go`: This file contains all the records of the sub-modules that exist in this directory with the path `/v1/**`.
- `./model`: This directory has all the project's global template files
- `./app`: Here you will find all the code used for orchestration and business rules of the service.
- `./store`: Here you will find all the code that is used to interact with the database.
- `./db`: Directory for creating databases and tables.
- `./utils`: Sub-modules required for project maintenance in general.


### Endpoints

* **Create**
`POST - http://localhost:8080/v1/accounts`
body request:

```
{
    "document_number": "12345678900",
}

```
Response
201 = Status Created

`POST - http://localhost:8080/v1/transactions`
body request:

```
{
    "account_id": 1,
    "operation_type_id": 4,
    "amount": 123.45
}

```
Response
201 = Status Created


* **ReadOne**
`GET - http://localhost:8080/v1/accounts/:accountId`
Response
200 = Status OK