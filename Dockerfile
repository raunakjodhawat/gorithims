FROM golang:1.14
# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/raunakjodhawat/gorithims

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download golangci-lint
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $GOPATH/bin v1.25.1
RUN go get -u golang.org/x/lint/golint
RUN go fmt ./...
RUN golint ./...

# Run tests
RUN go test -v -race ./... --cover -bench=.

# Run the executable
CMD ["go", "run", "src/dockerExecution.go"]