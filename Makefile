start:
	docker-compose up

stop:
	docker-compose down --remove-orphans

rebuild:
	make stop
	docker-compose build
	docker-compose up --force-recreate --build

run-server-local:
	IMSERVER_PORT=1234 IMSERVER_CORS=* JWT_SECRET=secret IMMUDB_HOST=0.0.0.0 IMMUDB_USERNAME=immudb IMMUDB_PASSWORD=immudb IMMUDB_DATABASE=defaultdb go run ./server.go

unit-test:
	go test -v ./pkg/... -coverprofile=coverage.out

integration-test:
	make stop
	make start &
	while ! echo exit | nc 0.0.0.0 1323; do sleep 1; done
	go test -v ./tests/...
	make stop

run-linter:
	golangci-lint run

gen-swagger:
	swag init -g server.go
