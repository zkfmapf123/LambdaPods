up:
	cd infrastructure && docker-compose up --build -d

down:
	cd infrastructure && docker-compose down

dev-run:
	cd cmd/examples && go run main.go

test:
	go test -v ./...
