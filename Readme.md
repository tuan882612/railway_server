# Railway Reservation Server

This is a the backend service for the cse3330 project 1

## Quick Note

- Loaded the **RRS.zip** data into **CochroachDB**, which is an free cloud service provide for Postgres database instances.
- Project is structured in the **Hexagonal Design Pattern**
- Used **Go-lang** as the main language and frameworks used along with it was **Gin-Gonic** for the API and **pgx/v4** for raw sql queries and database connection.

## Requirements

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
make {your os}     # example: make windows
make run-{your os} # example: make run-windows
```

---

## Testing

To Unit Test the entire service

- or look in the **Makefile** to find the commands for specific test

```make
make testAll
```

---

## API Endpoints

Base Endpoints

```txt
GET    /api/v1
GET    /api/v1/health
```

Passenger Endpoints

```txt
GET    /api/v1/passenger/all
GET    /api/v1/passenger/waiting
GET    /api/v1/passenger/area - param: code 
GET    /api/v1/passenger/riding - param: first-name, last-name
GET    /api/v1/passenger/confirmed - param: train-name
GET    /api/v1/passenger/traveling - param: day
GET    /api/v1/passenger/traveling/confirmed - param: day
```

Train Endpoints

```txt
GET    /api/v1/train/all
GET    /api/v1/train/info - param: name
GET    /api/v1/train/info/age - param: start, end
```

___

## Important Locations and Project Structure

- SQL queries are in directory **internal/domain/{entity}/service.go**
- Schemas are in directory **database/schema**

```text
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
│   ├── create.go
│   ├── load.go
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
