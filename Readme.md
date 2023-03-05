# Railway Reservation Server
This is a the backend service for the cse3330 project 1
## Quick Note
- Loaded the **RRS.zip** data into **CochroachDB**, which is an free cloud service provide for Postgres database instances.
- Project is structured in the **Hexagonal Design Pattern**
- Used **Go-lang** as the main language and frameworks used along with it was **Gin-Gonic** for the API and **pgx/v4** for raw sql queries and database connection.

## Requirements
#
You must have installed:
- [Golang 1.20.1](https://go.dev/doc/install)
---
## Project Setup

Change directory to railway_server.
```bash
cd railway_server
```

Install all dependencies.
```bash
go install .
```

Make an environment variable file and copy the contents **key.txt** file into it from the zip folder.
```bash
touch .env
```
---
## Running the Project
Running in Development
```make
make dev
```
Running in Production
- binary is located in directory **bin**
```make
make build
make run
```
---
## Testing
To Unit Test the entire service
- or look in the **Makefile** to find the commands for specific test
```make
make testAll
```
---
## Important Locations and Project Structure
- SQL queries are in directory **internal/domain/{entity}/service.go**
- Schemas are in directory **database/schema**
```
├── Makefile
├── Readme.md
├── bin
│   └── main-linux
├── cmd
│   └── api
│       └── main.go
├── database
│   ├── data
│   │   ├── booked.csv
│   │   ├── passenger.csv
│   │   ├── train.csv
│   │   └── train_status.csv
│   ├── pg.go
│   └── schema
│       ├── booked.sql
│       ├── passenger.sql
│       ├── train.sql
│       └── train_status.sql
├── go.mod
├── go.sum
├── internal
│   ├── api
│   │   ├── handler
│   │   │   ├── base.go
│   │   │   ├── passenger.go
│   │   │   └── train.go
│   │   ├── routes
│   │   │   ├── base.go
│   │   │   ├── passenger.go
│   │   │   └── train.go
│   │   └── server.go
│   ├── domain
│   │   ├── passenger
│   │   │   ├── passenger.go
│   │   │   ├── repository.go
│   │   │   └── service.go
│   │   └── train
│   │       ├── repository.go
│   │       ├── service.go
│   │       └── train.go
│   └── security
│       ├── middleware.go
│       ├── password.go
│       └── proxies.go
├── pkg
│   ├── config
│   │   └── dotenv.go
│   ├── response
│   │   └── base.go
│   └── validators
│       ├── dotenv.go
│       └── pg.go
└── test
    ├── config
    │   └── dotenv_test.go
    ├── constants.go
    ├── database
    │   └── pg_test.go
    └── handlers
        └── base_test.go