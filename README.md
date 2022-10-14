# GO-URL-SHORTENER

Simple URL shortener using **Golang**, [PostgreSQL](https://www.postgresql.org/), [Fiber](https://gofiber.io/) and [GORM](https://gorm.io/)

## Setup

Rename `.env.exmaple` to `.env`

Run database:

```sh
docker run --name goshort-db \
-e POSTGRES_USER=$DB_USERNAME -e POSTGRES_PASSWORD=$DB_PASSWORD \
-p $DB_PORT:$DB_PORT -d postgres:14
```

Run application:

In the root folder run `go run main.go` or `make run`

## REST API

### Get all Go-Shorts

`GET /goshort`

```sh
curl -i -H 'Accept: application/json' http://localhost:3000/goshort
```

### Get Go-Short

`GET /goshort/<id>`

```sh
curl -i -H 'Accept: application/json' http://localhost:3000/goshort/1
```

### Create Go-short

`POST /goshort`

```sh
curl -i -H 'Content-Type: application/json' \
-d '{"goshort":"", "redirect":"http://github.com", "random":true}' \
-X POST http://localhost:3000/goshort
```

### Update Go-short

`PATCH /goshort`

```sh
curl -i -H 'Content-Type: application/json' \
-d '  {"id": 1, "redirect": "https://github.com", "goshort": "gh", "clicked": 0, "random": false}' \
-X PATCH http://localhost:3000/goshort
```

### Delete Go-short

`DELETE /goshort`

```sh
curl -i -H 'Accept: application/json' -X DELETE http://localhost:3000/goshort/1
```

### Redirect

`GET /r/<goshort>`
