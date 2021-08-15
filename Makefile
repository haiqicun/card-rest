POSTGRESQL_URL="postgres://cardrest:cardrest_pass@localhost:5432/cardrest_db?sslmode=disable"

start-server: # Build up the server and start the server and postgresql in docker
	@echo "starting the server and postgresql in docker..."
	docker-compose up --build

migrate-up: # Migrate up the database
	migrate -database ${POSTGRESQL_URL} -path db/migrations up

migrage-down: # Migrate down the database
	migrate -database ${POSTGRESQL_URL} -path db/migrations down

test: ## Run the tests
	@echo "Running the tests..."
	go test -cover ./pkg/utils

stop-server: # Stop the server in docker
	@echo "Stoping the server"
	docker-compose down