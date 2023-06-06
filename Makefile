build:
	go build -o bin/librarian cmd/main.go

run:
	go run cmd/main.go --config=$(config)

unit_tests:
	docker-compose -f ./test/mysql/docker-compose.yaml up -d
	go test -v ./test
	docker-compose -f ./test/mysql/docker-compose.yaml down

demo_run:
	docker-compose -f ./test/mysql/docker-compose.yaml up -d
	go run cmd/main.go
