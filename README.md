# Order Management API

## Project Description
This service provides a REST API for managing orders, including authentication, caching, logging, and event handling. The project follows a **Clean Architecture** principles.

## Features
- **CRUD operations on orders** (`POST`, `PUT`, `GET`, `DELETE`)
- **Authentication** (Roles: `User`, `Admin`)
- **Order caching** (In-Memory)
- **Logging of operations** (to a file and/or database)
- **API metrics collection** (`GET /metrics`)

## How to Run Locally
To start the application locally, run:

```sh
docker-compose up
```

## Swagger
Once running, you can access Swagger API documentation at http://localhost:8080/docs/index.html