# download dependencies of air
run-on-air:
	go get github.com/cosmtrek/air
	go run github.com/cosmtrek/air

# make app
build:
	go build -o ./main

# setup application
set:
	go run setup.go
