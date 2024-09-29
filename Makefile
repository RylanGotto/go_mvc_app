BINARY_NAME=tonicApp

## build: builds all binaries
build:
	@go mod vendor
	@go build -o tmp/${BINARY_NAME} .
	@echo Celeritas built!

run: build
	@echo Staring Celeritas...
	@./tmp/${BINARY_NAME} &
	@echo Celeritas started!

clean:
	@echo Cleaning...
	@DEL ${BINARY_NAME}
	@go clean
	@echo Cleaned!

test:
	@echo Testing...
	@go test ./...
	@echo Done!

start: run
	
stop:
	@echo "Starting the front end..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo Stopped Celeritas

restart: stop start