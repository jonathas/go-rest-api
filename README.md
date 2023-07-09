# Go REST API

This is a basic project I implemented to practice how a REST API could be done in Go.
It uses gorilla/mux, gorm and Postgres (via Docker).

## Starting the project

Start and seed the DB:

```bash
docker compose -f ./resources/docker-compose.yml up
```

Start the server:

```bash
air
```

Ps: air provides live-reload for the server

Access the endpoint:

```bash
http://localhost:8080/api/personalities
```

A Postman collection can be found inside /docs
