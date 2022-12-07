# imserver


### Requirements:

- Docker
- Go 1.18+
- make
- nc (to run locally integration tests)
- [golangci-lint](https://github.com/golangci/golangci-lint)
- [swag](https://github.com/swaggo/swag)

*Note*: developed and tested on `macOS Ventura 13.0.1`

Commands:

- Start imserver `make start`
- Stop imserver `make stop`
- Rebuild imserver container `make rebuild`
- Run imserver on local machine `make run-server-local` (local immudb with defaults is required or from container)
- Run unit tests `make unit-test`
- Run integration tests `make integration-test`
- Run linter `make run-linter`
- Generate swagger documentation based on annotations `make gen-swagger`

### Predefined services

For testing and first usage there are created 3 services with different roles.

ServiceA:
- JWT: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzb3VyY2VJZCI6MX0.5VtXO9J1YF2sv8SwTfvsVseqHMjEwhFBHJLpSuj-i34`
- Roles: `ROLE_READ,ROLE_WRITE`

ServiceB:
- JWT: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzb3VyY2VJZCI6Mn0.KxzbtbC6E8TNt0NmmRBdNz5P9ixj6sSKE9JQVk3fkGg`
- Roles: `ROLE_WRITE`

ServiceC:
- JWT: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzb3VyY2VJZCI6M30.spUBA7KeBfZqevVxwfKGUAxXFHOJGvpOxV6-x339d-M`
- Roles: `` (deny all)

### Running locally
After starting imserver locally using `make start` available:

- Access to immudb http://127.0.0.1:8080/ (or 0.0.0.0:8080, depends on ENV and Docker configuration) with default login/password
- Access to swagger documentation of imserver http://127.0.0.1:1323/swagger/index.html
- Access to health endpoint of imserver http://127.0.0.1:1323/v1/health (other endpoints require JWT token)

### Generate logs using CLI

There is simple CLI to send dummy logs to local imserver using `ServiceA` JWT token.
example: `go run ./cmd/main.go send`

### Generate logs using Curl

*Note*: examples are using `ServiceA` JWT token.

#### Create single log

``
curl \
--header "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzb3VyY2VJZCI6MX0.5VtXO9J1YF2sv8SwTfvsVseqHMjEwhFBHJLpSuj-i34" \
--header "Content-Type: application/json" \
--request POST \
--data '{"metrics":"xyz"}' \
http://127.0.0.1:1323/v1/logs
``


#### Create logs in batch

``
curl \
--header "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzb3VyY2VJZCI6MX0.5VtXO9J1YF2sv8SwTfvsVseqHMjEwhFBHJLpSuj-i34" \
--header "Content-Type: application/json" \
--request POST \
--data '[{"metrics": "measure1"}, {"metrics": "measure2"}]' \
http://127.0.0.1:1323/v1/logs/batch
``

#### Get logs

``
curl http://127.0.0.1:1323/v1/logs -H 'Accept: application/json' -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzb3VyY2VJZCI6MX0.5VtXO9J1YF2sv8SwTfvsVseqHMjEwhFBHJLpSuj-i34"
``

#### Get logs limit 3:

``
curl http://127.0.0.1:1323/v1/logs?filter=1 -H 'Accept: application/json' -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzb3VyY2VJZCI6MX0.5VtXO9J1YF2sv8SwTfvsVseqHMjEwhFBHJLpSuj-i34"
``

#### Get logs count:

``
curl http://127.0.0.1:1323/v1/logs/count -H 'Accept: application/json' -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzb3VyY2VJZCI6MX0.5VtXO9J1YF2sv8SwTfvsVseqHMjEwhFBHJLpSuj-i34"
``
