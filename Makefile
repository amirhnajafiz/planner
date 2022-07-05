# make app
build:
	go build -o ./main

# run app
start:
	go run main.go

# migrate up
mig-up:
	go run setup.go migrate

# migrate down
mig-down:
	go run internal/cmd/setup/setup.go no-migrate

# setup table
db-set:
	docker exec planner-db createdb planner

# docker up
dk-up:
	docker compose up -d

# docker down
dk-down:
	docker compose down
