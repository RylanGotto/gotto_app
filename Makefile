BINARY_NAME=mvc

build:
	@go mod vendor
	@echo "Building ${BINARY_NAME}..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "${BINARY_NAME} built!"

run: build
	@echo "Starting ${BINARY_NAME}..."
	@./tmp/${BINARY_NAME} &
	@echo "${BINARY_NAME} started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping Celeritas..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped Celeritas!"

restart: stop start