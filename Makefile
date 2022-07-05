# make app
build:
	go build -o ./main

# migrate up
m-up:
	go run setup.go migrate

# migrate down
m-down:
	go run internal/cmd/setup/setup.go no-migrate

# setup table
d-start:
	docker exec planner-db createdb planner
