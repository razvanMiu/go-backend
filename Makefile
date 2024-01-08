BINARY_NAME=main.exe

build:
	go build -o ${BINARY_NAME} -v

run:
	go run main.go

clean:
	go clean
	rm -f ${BINARY_NAME}
	rm -rf ./frontend/out

dev:
ifeq ($(wildcard ./frontend/out), )
	cd frontend && npm run build && cd ..
endif
ifeq ($(shell which air),)
	echo "air is not installed, follow air installation guide: https://github.com/cosmtrek/air?tab=readme-ov-file#installation"
else
	air
endif

swag:
ifeq ($(shell which swag),)
	echo "swag is not installed, follow swag installation guide: https://github.com/swaggo/swag?tab=readme-ov-file#getting-started"
else
	swag init
endif

test:
	go test -coverprofile=coverage.out

coverage:
	go tool cover -html=coverage.out