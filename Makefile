# make app
build:
	go build -o ./main

# migrate up
m-up:
	go run setup.go migrate

# migrate down
m-down:
	go run setup.go no-migrate
