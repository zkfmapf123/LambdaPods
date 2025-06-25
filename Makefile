_build:
	go build -o lambdapods main.go

run: _build
	./lambdapods