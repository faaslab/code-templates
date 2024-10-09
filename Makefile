# Define the binary name
BINARY_NAME=main

# Build command for Linux AMD64
build:
	cd golang-1x && GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME} main.go && zip -r ../golang-1x.zip .

# Clean command to remove the binary
clean:
	rm -f ${BINARY_NAME}

# Run command to build and execute the binary
run: build
	./${BINARY_NAME}
