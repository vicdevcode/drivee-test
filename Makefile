BINARY_PATH := ./bin/app
BINARY_PATH_WINDOWS := ./bin/app.exe
MAIN_PATH := ./cmd/app/main.go

PHONY: run-windows 

run-windows:
	@go build -o ${BINARY_PATH_WINDOWS} ${MAIN_PATH} 
	@${BINARY_PATH_WINDOWS}

build:
	@echo 'Building golang'
	@GOARCH=amd64 GOOS=linux go build -o ${BINARY_PATH} ${MAIN_PATH}	
	@echo 'Build completed'

run: build
	@echo 'Starting server'
	@${BINARY_PATH}
