BINARY_NAME=main.exe

build:
	go build -o ${BINARY_NAME} -v

run:
	go run main.go

dev:
ifeq ($(shell which air),)
	echo "air is not installed, follow air installation guide: https://github.com/cosmtrek/air?tab=readme-ov-file#installation"
else
	air
endif

test:
	go test -coverprofile=coverage.out

coverage:
	go tool cover -html=coverage.out