run:
	cd infrastructure && docker-compose up --build -d

example-run:
	cd cmd/examples && go run main.go

test:
	go test -v ./...
